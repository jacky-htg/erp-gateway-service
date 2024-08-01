package service

import (
	"context"
	"io"
	"log"

	"github.com/jacky-htg/erp-proto/go/pb/inventories"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Warehouse struct {
	Client inventories.WarehouseServiceClient
	inventories.UnimplementedWarehouseServiceServer
}

// Create new Warehouse
func (u *Warehouse) Create(ctx context.Context, in *inventories.Warehouse) (*inventories.Warehouse, error) {
	return u.Client.Create(ctx, in)
}

// Update Warehouse
func (u *Warehouse) Update(ctx context.Context, in *inventories.Warehouse) (*inventories.Warehouse, error) {
	return u.Client.Update(ctx, in)
}

// View Warehouse
func (u *Warehouse) View(ctx context.Context, in *inventories.Id) (*inventories.Warehouse, error) {
	return u.Client.View(ctx, in)
}

// Delete Warehouse
func (u *Warehouse) Delete(ctx context.Context, in *inventories.Id) (*inventories.MyBoolean, error) {
	return u.Client.Delete(ctx, in)
}

// List Warehouses
func (u *Warehouse) List(in *inventories.ListWarehouseRequest, streamServer inventories.WarehouseService_ListServer) error {
	ctx := streamServer.Context()
	streamClient, err := u.Client.List(ctx, in)
	if err != nil {
		return status.Errorf(codes.Internal, "Error when calling Warehouse.List service: %s", err)
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
