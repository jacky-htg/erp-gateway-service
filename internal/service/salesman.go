package service

import (
	"context"
	"io"
	"log"

	"github.com/jacky-htg/erp-proto/go/pb/sales"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Salesman struct
type Salesman struct {
	Client sales.SalesmanServiceClient
	sales.UnimplementedSalesmanServiceServer
}

// Create new Salesman
func (u *Salesman) SalesmanCreate(ctx context.Context, in *sales.Salesman) (*sales.Salesman, error) {
	return u.Client.SalesmanCreate(ctx, in)
}

// Update Salesman
func (u *Salesman) SalesmanUpdate(ctx context.Context, in *sales.Salesman) (*sales.Salesman, error) {
	return u.Client.SalesmanUpdate(ctx, in)
}

// View Salesman
func (u *Salesman) SalesmanView(ctx context.Context, in *sales.Id) (*sales.Salesman, error) {
	return u.Client.SalesmanView(ctx, in)
}

// List Salesmans
func (u *Salesman) SalesmanList(in *sales.ListSalesmanRequest, streamServer sales.SalesmanService_SalesmanListServer) error {
	ctx := streamServer.Context()
	streamClient, err := u.Client.SalesmanList(ctx, in)
	if err != nil {
		return status.Errorf(codes.Internal, "Error when calling Salesman.List service: %s", err)
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
