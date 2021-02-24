package service

import (
	"context"
	"erp-gateway/pb/users"
	"log"
)

// ViewAccessTree service client
func ViewAccessTree(ctx context.Context, access users.AccessServiceClient) {
	response, err := access.List(setMetadata(ctx), &users.Empty{})
	if err != nil {
		log.Fatalf("Error when calling grpc service: %s", err)
	}
	log.Printf("Resp received: %v", response)
}
