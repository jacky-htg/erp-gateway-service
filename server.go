package main

import (
	"inventory-gateway-service/internal/middleware"
	"net"
	"os"

	_ "github.com/lib/pq"
	"google.golang.org/grpc"

	"inventory-gateway-service/internal/config"
	"inventory-gateway-service/internal/pkg/log/logruslog"
	"inventory-gateway-service/internal/route"
	users "inventory-gateway-service/pb"
)

const defaultPort = "8080"

func main() {
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

	var userServiceConn *grpc.ClientConn
	userServiceConn, err = grpc.Dial(os.Getenv("USER_SERVICE"), grpc.WithInsecure())
	if err != nil {
		log.Errorf("user service connection: %v", err)
	}
	defer userServiceConn.Close()
	grpcClient := getGrpcClient(userServiceConn)

	loginInterceptor := middleware.Login{Client: grpcClient["UserClient"].(users.UserServiceClient)}
	serverOptions := []grpc.ServerOption{
		grpc.UnaryInterceptor(loginInterceptor.Unary()),
		// grpc.StreamInterceptor(interceptor.Stream()),
	}

	grpcServer := grpc.NewServer(serverOptions...)

	// routing grpc services
	route.GrpcRoute(grpcServer, grpcClient)
	log.Info("serve grpc on port: " + port)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
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
