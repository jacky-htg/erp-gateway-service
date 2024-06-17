package main

import (
	"crypto/tls"
	"crypto/x509"
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/improbable-eng/grpc-web/go/grpcweb"
	_ "github.com/lib/pq"
	"github.com/rs/cors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"

	"erp-gateway/internal/config"
	"erp-gateway/internal/middleware"
	"erp-gateway/internal/route"
	"erp-gateway/pb/inventories"
	"erp-gateway/pb/users"
)

const defaultPort = "8080"

type RpcServer struct {
	Grpc        *grpc.Server
	WrappedGrpc *grpcweb.WrappedGrpcServer
}

func NewServer(serverOptions []grpc.ServerOption) *RpcServer {
	gs := grpc.NewServer(serverOptions...)
	return &RpcServer{
		Grpc:        gs,
		WrappedGrpc: grpcweb.WrapServer(gs),
	}
}

func main() {
	enableTLS := flag.Bool("tls", false, "enable SSL/TLS")
	flag.Parse()

	if _, ok := os.LookupEnv("PORT"); !ok {
		config.Setup(".env")
	}

	log := log.New(os.Stdout, "Essentials : ", log.LstdFlags|log.Lmicroseconds|log.Lshortfile)
	if err := run(enableTLS); err != nil {
		log.Printf("error: shutting down: %s", err)
		os.Exit(1)
	}
}

func run(enableTLS *bool) error {

	if len(os.Getenv("PORT")) == 0 {
		os.Setenv("PORT", defaultPort)
	}

	port := map[string]string{"grpc": os.Getenv("PORT"), "web": "9090"}
	clientCreds := insecure.NewCredentials()
	userServiceConn, err := grpc.NewClient(os.Getenv("USER_SERVICE"), grpc.WithTransportCredentials(clientCreds))
	if err != nil {
		return fmt.Errorf("user service connection: %v", err)
	}
	defer userServiceConn.Close()

	inventoryServiceConn, err := grpc.NewClient(os.Getenv("INVENTORY_SERVICE"), grpc.WithTransportCredentials(clientCreds))
	if err != nil {
		return fmt.Errorf("inventory service connection: %v", err)
	}
	defer userServiceConn.Close()

	grpcClient := getGrpcClient(userServiceConn, inventoryServiceConn)

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

	if *enableTLS {
		tlsCredentials, err := loadTLSCredentials()
		if err != nil {
			return fmt.Errorf("cannot load TLS credentials: %w", err)
		}

		serverOptions = append(serverOptions, grpc.Creds(tlsCredentials))
	}

	rpcServer := NewServer(serverOptions)
	route.GrpcRoute(rpcServer.Grpc, grpcClient)

	errChan := make(chan error)
	go func() {
		errChan <- runGrpcServer(port["grpc"], rpcServer)
	}()

	go func() {
		errChan <- runWebServer(port["web"], rpcServer)
	}()

	if e := <-errChan; e != nil {
		return e
	}

	return nil
}

func getGrpcClient(userServiceConn, inventoryServiceConn *grpc.ClientConn) map[string]interface{} {
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

	grpcClient["BrandClient"] = inventories.NewBrandServiceClient(inventoryServiceConn)

	return grpcClient
}

func runGrpcServer(port string, rpcServer *RpcServer) error {
	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		return err
	}

	if err := rpcServer.Grpc.Serve(listener); err != nil {
		return fmt.Errorf("failed to serve: %s", err)
	}

	return nil
}

func loadTLSCredentials() (credentials.TransportCredentials, error) {
	// Load certificate of the CA who signed client's certificate
	var certPool *x509.CertPool
	if len(os.Getenv("TSL_CLIENT_CERTIFICATE")) > 0 {
		pemClientCA, err := os.ReadFile(os.Getenv("TSL_CLIENT_CERTIFICATE"))
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

func runWebServer(httpPort string, rpcServer *RpcServer) error {
	grpc := rpcServer.WrappedGrpc

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(resp http.ResponseWriter, req *http.Request) {
		if grpc.IsGrpcWebRequest(req) || grpc.IsAcceptableGrpcCorsRequest(req) {
			grpc.ServeHTTP(resp, req)
		}
	})

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"*"},
		AllowCredentials: true,
		// Enable Debugging for testing, consider disabling in production
		Debug: false,
	})

	// Insert the middleware
	handler := c.Handler(mux)

	//handler := cors.Default().Handler(mux)

	err := http.ListenAndServe(":"+httpPort, handler)
	if err != nil {
		return err
	}

	return nil
}
