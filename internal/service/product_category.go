package service

import (
	"context"
	"erp-gateway/pb/inventories"
	"io"
	"log"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ProductCategory struct {
	Client inventories.ProductCategoryServiceClient
	inventories.UnimplementedProductCategoryServiceServer
}

func (u *ProductCategory) Create(ctx context.Context, in *inventories.ProductCategory) (*inventories.ProductCategory, error) {
	return u.Client.Create(ctx, in)
}

func (u *ProductCategory) Update(ctx context.Context, in *inventories.ProductCategory) (*inventories.ProductCategory, error) {
	return u.Client.Update(ctx, in)
}

func (u *ProductCategory) View(ctx context.Context, in *inventories.Id) (*inventories.ProductCategory, error) {
	return u.Client.View(ctx, in)
}

func (u *ProductCategory) Delete(ctx context.Context, in *inventories.Id) (*inventories.MyBoolean, error) {
	return u.Client.Delete(ctx, in)
}

func (u *ProductCategory) List(in *inventories.ListProductCategoryRequest, streamServer inventories.ProductCategoryService_ListServer) error {
	ctx := streamServer.Context()
	streamClient, err := u.Client.List(ctx, in)
	if err != nil {
		return status.Errorf(codes.Internal, "Error when calling ProductCategory.List service: %s", err)
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
