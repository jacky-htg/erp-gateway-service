package main

import (
	"crypto/tls"
	"crypto/x509"
	"flag"
	"fmt"
	"inventory-gateway-service/internal/middleware"
	"io/ioutil"
	"net"
	"os"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	"inventory-gateway-service/internal/config"
	"inventory-gateway-service/internal/pkg/log/logruslog"
	"inventory-gateway-service/internal/route"
	users "inventory-gateway-service/pb"
)

const defaultPort = "8080"

func main() {
	enableTLS := flag.Bool("tls", false, "enable SSL/TLS")
	flag.Parse()

	// lookup and setup env
	if _, ok := os.LookupEnv("PORT"); !ok {
		config.Setup(".env")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	// init log
	log := logruslog.Init()

	// listen tcp port
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	log.Info("listen port:", port)

	if err = runGrpcServer(*enableTLS, lis); err != nil {
		log.Fatalf("failed run grpc server: %v", err)
		return
	}
}

func getGrpcClient(userServiceConn *grpc.ClientConn) map[string]interface{} {
	grpcClient := make(map[string]interface{})
	grpcClient["AuthClient"] = users.NewAuthServiceClient(userServiceConn)
	grpcClient["UserClient"] = users.NewUserServiceClient(userServiceConn)
	grpcClient["CompanyClient"] = users.NewCompanyServiceClient(userServiceConn)
	grpcClient["RegionClient"] = users.NewRegionServiceClient(userServiceConn)
	grpcClient["BranchClient"] = users.NewBranchServiceClient(userServiceConn)
	grpcClient["EmployeeClient"] = users.NewEmployeeServiceClient(userServiceConn)
	grpcClient["FeatureClient"] = users.NewFeatureServiceClient(userServiceConn)
	grpcClient["PackageFeatureClient"] = users.NewPackageFeatureServiceClient(userServiceConn)
	grpcClient["AccessClient"] = users.NewAccessServiceClient(userServiceConn)
	grpcClient["GroupClient"] = users.NewGroupServiceClient(userServiceConn)

	return grpcClient
}

func runGrpcServer(enableTLS bool, listener net.Listener) error {
	userServiceConn, err := grpc.Dial(os.Getenv("USER_SERVICE"), grpc.WithInsecure())
	if err != nil {
		return fmt.Errorf("user service connection: %v", err)
	}
	defer userServiceConn.Close()
	grpcClient := getGrpcClient(userServiceConn)

	loginInterceptor := middleware.Login{Client: grpcClient["UserClient"].(users.UserServiceClient)}
	authInterceptor := middleware.Auth{Client: grpcClient["AuthClient"].(users.AuthServiceClient)}
	serverOptions := []grpc.ServerOption{
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			loginInterceptor.Unary(),
			authInterceptor.Unary(),
		)),
		grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(
			loginInterceptor.Stream(),
			authInterceptor.Stream(),
		)),
	}

	if enableTLS {
		tlsCredentials, err := loadTLSCredentials()
		if err != nil {
			return fmt.Errorf("cannot load TLS credentials: %w", err)
		}

		serverOptions = append(serverOptions, grpc.Creds(tlsCredentials))
	}

	grpcServer := grpc.NewServer(serverOptions...)

	// routing grpc services
	route.GrpcRoute(grpcServer, grpcClient)
	if err := grpcServer.Serve(listener); err != nil {
		return fmt.Errorf("failed to serve: %s", err)
	}

	return nil
}

func loadTLSCredentials() (credentials.TransportCredentials, error) {
	// Load certificate of the CA who signed client's certificate
	var certPool *x509.CertPool
	if len(os.Getenv("TSL_CLIENT_CERTIFICATE")) > 0 {
		pemClientCA, err := ioutil.ReadFile(os.Getenv("TSL_CLIENT_CERTIFICATE"))
		if err != nil {
			return nil, err
		}

		certPool := x509.NewCertPool()
		if !certPool.AppendCertsFromPEM(pemClientCA) {
			return nil, fmt.Errorf("failed to add client CA's certificate")
		}
	}

	// Load server's certificate and private key
	serverCert, err := tls.LoadX509KeyPair(os.Getenv("TLS_SERVER_CERTIFICATE"), os.Getenv("TLS_SERVER_KEY"))
	if err != nil {
		return nil, err
	}

	// Create the credentials and return it
	config := &tls.Config{
		Certificates: []tls.Certificate{serverCert},
		ClientAuth:   tls.RequireAndVerifyClientCert,
		ClientCAs:    certPool,
	}

	return credentials.NewTLS(config), nil
}
