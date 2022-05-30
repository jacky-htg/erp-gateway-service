gen:
	protoc --proto_path=proto proto/users/*.proto --go_out=. --go-grpc_out=. 

init:
	go mod init erp-gateway

migrate:
	go run cmd/cli.go migrate
	
seed:
	go run cmd/cli.go seed

server:
	go run server.go

server-tls:
	go run server.go --tls

.PHONY: gen init migrate seed server