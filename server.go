package main

import (
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

	grpcServer := grpc.NewServer()

	grpcClient := make(map[string]interface{})

	var userSeviceConn *grpc.ClientConn
	userSeviceConn, err = grpc.Dial(":8000", grpc.WithInsecure())
	if err != nil {
		log.Errorf("user service connection: %v", err)
	}
	defer userSeviceConn.Close()

	grpcClient["AuthClient"] = users.NewAuthServiceClient(userSeviceConn)
	grpcClient["UserClient"] = users.NewUserServiceClient(userSeviceConn)
	grpcClient["CompanyClient"] = users.NewCompanyServiceClient(userSeviceConn)
	grpcClient["RegionClient"] = users.NewRegionServiceClient(userSeviceConn)
	grpcClient["BranchClient"] = users.NewBranchServiceClient(userSeviceConn)
	grpcClient["EmployeeClient"] = users.NewEmployeeServiceClient(userSeviceConn)
	grpcClient["FeatureClient"] = users.NewFeatureServiceClient(userSeviceConn)
	grpcClient["PackageFeatureClient"] = users.NewPackageFeatureServiceClient(userSeviceConn)
	grpcClient["AccessClient"] = users.NewAccessServiceClient(userSeviceConn)
	grpcClient["GroupClient"] = users.NewGroupServiceClient(userSeviceConn)

	// routing grpc services
	route.GrpcRoute(grpcServer, grpcClient)
	log.Info("serve grpc on port: " + port)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
		return
	}
}
