package service

import (
	"context"
	"erp-gateway/pb/inventories"
	"io"
	"log"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Product struct {
	Client inventories.ProductServiceClient
	inventories.UnimplementedProductServiceServer
}

// Create new Product
func (u *Product) Create(ctx context.Context, in *inventories.Product) (*inventories.Product, error) {
	return u.Client.Create(ctx, in)
}

// Update Product
func (u *Product) Update(ctx context.Context, in *inventories.Product) (*inventories.Product, error) {
	return u.Client.Update(ctx, in)
}

// View Product
func (u *Product) View(ctx context.Context, in *inventories.Id) (*inventories.Product, error) {
	return u.Client.View(ctx, in)
}

// Delete Product
func (u *Product) Delete(ctx context.Context, in *inventories.Id) (*inventories.MyBoolean, error) {
	return u.Client.Delete(ctx, in)
}

// List Products
func (u *Product) List(in *inventories.ListProductRequest, streamServer inventories.ProductService_ListServer) error {
	ctx := streamServer.Context()
	streamClient, err := u.Client.List(ctx, in)
	if err != nil {
		return status.Errorf(codes.Internal, "Error when calling Product.List service: %s", err)
	}

	for {
		resp, err := streamClient.Recv()
		if err == io.EOF {
			log.Print("end stream")
			break
		}
		if err != nil {
			return status.Errorf(codes.Internal, "cannot receive %v", err)
		}

		err = streamServer.Send(resp)
		if err != nil {
			return status.Errorf(codes.Unknown, "cannot send stream response: %v", err)
		}
	}

	return nil
}
