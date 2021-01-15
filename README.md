# Inventory Microservices
api gateway for inventory micro sevices using golang and grpc framework.
- As api gateway, the service will be accessed by public. 
- The service using grpc secure connection.
- The service will call other service in local network.

## Architecture
![Inventory Micro Services Arcgitecture](./inventory-micro-services.png)

## Features
- [X] [User Service](https://github.com/jacky-htg/user-service)
- [ ] Sales Service
- [ ] Purchase Service
- [ ] Warehouse Service
- [ ] General Ledger Service

## Get Started 
- git clone git@github.com:jacky-htg/api-gateway-service.git
- make init
- cp .env.example .env (and edit with your environment)
- make server
- You can test the service using `go run client/main.go` and select the test case on file client/main.go

## How To Contrubute
- Give star or clone and fork the repository
- Report the bug
- Submit issue for request of enhancement
- Pull Request for fixing bug or enhancement module 

## License
[The license of application is GPL-3.0](https://github.com/jacky-htg/api-gateway-service/blob/main/LICENSE), You can use this apllication for commercial use, distribution or modification. But there is no liability and warranty. Please read the license details carefully.

## Link Repository
- [Simple gRPC Skeleton](https://github.com/jacky-htg/grpc-skeleton)