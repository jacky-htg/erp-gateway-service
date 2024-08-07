package route

import (
	"github.com/jacky-htg/erp-gateway-service/internal/service"
	"github.com/jacky-htg/erp-proto/go/pb/inventories"
	"github.com/jacky-htg/erp-proto/go/pb/purchases"
	"github.com/jacky-htg/erp-proto/go/pb/sales"
	"github.com/jacky-htg/erp-proto/go/pb/users"
	"google.golang.org/grpc"
)

// GrpcRoute func
func GrpcRoute(grpcServer *grpc.Server, grpcClient map[string]interface{}) {

	authServer := service.Auth{Client: grpcClient["AuthClient"].(users.AuthServiceClient)}
	users.RegisterAuthServiceServer(grpcServer, &authServer)

	userServer := service.User{Client: grpcClient["UserClient"].(users.UserServiceClient)}
	users.RegisterUserServiceServer(grpcServer, &userServer)

	companyServer := service.Company{Client: grpcClient["CompanyClient"].(users.CompanyServiceClient)}
	users.RegisterCompanyServiceServer(grpcServer, &companyServer)

	regionServer := service.Region{Client: grpcClient["RegionClient"].(users.RegionServiceClient)}
	users.RegisterRegionServiceServer(grpcServer, &regionServer)

	branchServer := service.Branch{Client: grpcClient["BranchClient"].(users.BranchServiceClient)}
	users.RegisterBranchServiceServer(grpcServer, &branchServer)

	employeeServer := service.Employee{Client: grpcClient["EmployeeClient"].(users.EmployeeServiceClient)}
	users.RegisterEmployeeServiceServer(grpcServer, &employeeServer)

	featureServer := service.Feature{Client: grpcClient["FeatureClient"].(users.FeatureServiceClient)}
	users.RegisterFeatureServiceServer(grpcServer, &featureServer)

	packageFeatureServer := service.PackageFeature{Client: grpcClient["PackageFeatureClient"].(users.PackageFeatureServiceClient)}
	users.RegisterPackageFeatureServiceServer(grpcServer, &packageFeatureServer)

	accessServer := service.Access{Client: grpcClient["AccessClient"].(users.AccessServiceClient)}
	users.RegisterAccessServiceServer(grpcServer, &accessServer)

	groupServer := service.Group{Client: grpcClient["GroupClient"].(users.GroupServiceClient)}
	users.RegisterGroupServiceServer(grpcServer, &groupServer)

	brandServer := service.Brand{Client: grpcClient["BrandClient"].(inventories.BrandServiceClient)}
	inventories.RegisterBrandServiceServer(grpcServer, &brandServer)

	categoryServer := service.Category{Client: grpcClient["CategoryClient"].(inventories.CategoryServiceClient)}
	inventories.RegisterCategoryServiceServer(grpcServer, &categoryServer)

	productCategoryServer := service.ProductCategory{Client: grpcClient["ProductCategoryClient"].(inventories.ProductCategoryServiceClient)}
	inventories.RegisterProductCategoryServiceServer(grpcServer, &productCategoryServer)

	productServer := service.Product{Client: grpcClient["ProductClient"].(inventories.ProductServiceClient)}
	inventories.RegisterProductServiceServer(grpcServer, &productServer)

	warehouseServer := service.Warehouse{Client: grpcClient["WarehouseClient"].(inventories.WarehouseServiceClient)}
	inventories.RegisterWarehouseServiceServer(grpcServer, &warehouseServer)

	shelveServer := service.Shelve{Client: grpcClient["ShelveClient"].(inventories.ShelveServiceClient)}
	inventories.RegisterShelveServiceServer(grpcServer, &shelveServer)

	receiveServer := service.Receive{Client: grpcClient["ReceiveClient"].(inventories.ReceiveServiceClient)}
	inventories.RegisterReceiveServiceServer(grpcServer, &receiveServer)

	supplierServer := service.Supplier{Client: grpcClient["SupplierClient"].(purchases.SupplierServiceClient)}
	purchases.RegisterSupplierServiceServer(grpcServer, &supplierServer)

	purchaseServer := service.Purchase{Client: grpcClient["PurchaseClient"].(purchases.PurchaseServiceClient)}
	purchases.RegisterPurchaseServiceServer(grpcServer, &purchaseServer)

	purchaseReturnServer := service.PurchaseReturn{Client: grpcClient["PurchaseReturnClient"].(purchases.PurchaseReturnServiceClient)}
	purchases.RegisterPurchaseReturnServiceServer(grpcServer, &purchaseReturnServer)

	customerServer := service.Customer{Client: grpcClient["CustomerClient"].(sales.CustomerServiceClient)}
	sales.RegisterCustomerServiceServer(grpcServer, &customerServer)

	salemanServer := service.Salesman{Client: grpcClient["SalesmanClient"].(sales.SalesmanServiceClient)}
	sales.RegisterSalesmanServiceServer(grpcServer, &salemanServer)
}
