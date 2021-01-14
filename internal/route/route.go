package route

import (
	"google.golang.org/grpc"

	"inventory-gateway-service/internal/service"
	users "inventory-gateway-service/pb"
)

// GrpcRoute func
func GrpcRoute(grpcServer *grpc.Server, grpcClient map[string]interface{}) {

	authServer := service.Auth{
		AuthClient: grpcClient["AuthClient"].(users.AuthServiceClient),
		UserClient: grpcClient["UserClient"].(users.UserServiceClient),
	}
	users.RegisterAuthServiceServer(grpcServer, &authServer)

	userServer := service.User{
		UserClient: grpcClient["UserClient"].(users.UserServiceClient),
	}
	users.RegisterUserServiceServer(grpcServer, &userServer)

	companyServer := service.Company{
		CompanyClient: grpcClient["CompanyClient"].(users.CompanyServiceClient),
		UserClient:    grpcClient["UserClient"].(users.UserServiceClient),
	}
	users.RegisterCompanyServiceServer(grpcServer, &companyServer)

	regionServer := service.Region{
		RegionClient: grpcClient["RegionClient"].(users.RegionServiceClient),
		UserClient:   grpcClient["UserClient"].(users.UserServiceClient),
	}
	users.RegisterRegionServiceServer(grpcServer, &regionServer)

	branchServer := service.Branch{
		BranchClient: grpcClient["BranchClient"].(users.BranchServiceClient),
		UserClient:   grpcClient["UserClient"].(users.UserServiceClient),
	}
	users.RegisterBranchServiceServer(grpcServer, &branchServer)

	employeeServer := service.Employee{
		EmployeeClient: grpcClient["EmployeeClient"].(users.EmployeeServiceClient),
		UserClient:     grpcClient["UserClient"].(users.UserServiceClient),
	}
	users.RegisterEmployeeServiceServer(grpcServer, &employeeServer)

	featureServer := service.Feature{
		FeatureClient: grpcClient["FeatureClient"].(users.FeatureServiceClient),
		UserClient:    grpcClient["UserClient"].(users.UserServiceClient),
	}
	users.RegisterFeatureServiceServer(grpcServer, &featureServer)

	packageFeatureServer := service.PackageFeature{
		PackageFeatureClient: grpcClient["PackageFeatureClient"].(users.PackageFeatureServiceClient),
		UserClient:           grpcClient["UserClient"].(users.UserServiceClient),
	}
	users.RegisterPackageFeatureServiceServer(grpcServer, &packageFeatureServer)

	accessServer := service.Access{
		AccessClient: grpcClient["AccessClient"].(users.AccessServiceClient),
		UserClient:   grpcClient["UserClient"].(users.UserServiceClient),
	}
	users.RegisterAccessServiceServer(grpcServer, &accessServer)

	groupServer := service.Group{
		GroupClient: grpcClient["GroupClient"].(users.GroupServiceClient),
		UserClient:  grpcClient["UserClient"].(users.UserServiceClient),
	}
	users.RegisterGroupServiceServer(grpcServer, &groupServer)
}
