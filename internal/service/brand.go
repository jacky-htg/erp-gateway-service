package service

import (
	"context"
	"erp-gateway/pb/inventories"
	"io"
	"log"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Brand struct
type Brand struct {
	Client inventories.BrandServiceClient
	inventories.UnimplementedBrandServiceServer
}

// Create new Brand
func (u *Brand) Create(ctx context.Context, in *inventories.Brand) (*inventories.Brand, error) {
	return u.Client.Create(ctx, in)
}

// Update Brand
func (u *Brand) Update(ctx context.Context, in *inventories.Brand) (*inventories.Brand, error) {
	return u.Client.Update(ctx, in)
}

// View Brand
func (u *Brand) View(ctx context.Context, in *inventories.Id) (*inventories.Brand, error) {
	return u.Client.View(ctx, in)
}

// Delete Brand
func (u *Brand) Delete(ctx context.Context, in *inventories.Id) (*inventories.MyBoolean, error) {
	return u.Client.Delete(ctx, in)
}

// List Brands
func (u *Brand) List(in *inventories.Pagination, streamServer inventories.BrandService_ListServer) error {
	ctx := streamServer.Context()
	streamClient, err := u.Client.List(ctx, in)
	if err != nil {
		return status.Errorf(codes.Internal, "Error when calling Inventories.List service: %s", err)
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
