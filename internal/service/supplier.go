package service

import (
	"context"
	"erp-gateway/pb/purchases"
	"io"
	"log"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Supplier struct
type Supplier struct {
	Client purchases.SupplierServiceClient
	purchases.UnimplementedSupplierServiceServer
}

// Create new Supplier
func (u *Supplier) SupplierCreate(ctx context.Context, in *purchases.Supplier) (*purchases.Supplier, error) {
	return u.Client.SupplierCreate(ctx, in)
}

// Update Supplier
func (u *Supplier) SupplierUpdate(ctx context.Context, in *purchases.Supplier) (*purchases.Supplier, error) {
	return u.Client.SupplierUpdate(ctx, in)
}

// View Supplier
func (u *Supplier) SupplierView(ctx context.Context, in *purchases.Id) (*purchases.Supplier, error) {
	return u.Client.SupplierView(ctx, in)
}

// List Suppliers
func (u *Supplier) SupplierList(in *purchases.ListSupplierRequest, streamServer purchases.SupplierService_SupplierListServer) error {
	ctx := streamServer.Context()
	streamClient, err := u.Client.SupplierList(ctx, in)
	if err != nil {
		return status.Errorf(codes.Internal, "Error when calling Supplier.List service: %s", err)
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
