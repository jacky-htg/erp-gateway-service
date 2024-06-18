package service

import (
	"context"
	"erp-gateway/pb/sales"
	"io"
	"log"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Customer struct
type Customer struct {
	Client sales.CustomerServiceClient
	sales.UnimplementedCustomerServiceServer
}

// Create new Customer
func (u *Customer) CustomerCreate(ctx context.Context, in *sales.Customer) (*sales.Customer, error) {
	return u.Client.CustomerCreate(ctx, in)
}

// Update Customer
func (u *Customer) CustomerUpdate(ctx context.Context, in *sales.Customer) (*sales.Customer, error) {
	return u.Client.CustomerUpdate(ctx, in)
}

// View Customer
func (u *Customer) CustomerView(ctx context.Context, in *sales.Id) (*sales.Customer, error) {
	return u.Client.CustomerView(ctx, in)
}

// List Customers
func (u *Customer) CustomerList(in *sales.ListCustomerRequest, streamServer sales.CustomerService_CustomerListServer) error {
	ctx := streamServer.Context()
	streamClient, err := u.Client.CustomerList(ctx, in)
	if err != nil {
		return status.Errorf(codes.Internal, "Error when calling Customer.List service: %s", err)
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
