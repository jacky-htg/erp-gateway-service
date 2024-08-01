package service

import (
	"context"
	"io"
	"log"

	"github.com/jacky-htg/erp-proto/go/pb/inventories"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Receive struct
type Receive struct {
	Client inventories.ReceiveServiceClient
	inventories.UnimplementedReceiveServiceServer
}

// Create new Receive
func (u *Receive) Create(ctx context.Context, in *inventories.Receive) (*inventories.Receive, error) {
	return u.Client.Create(ctx, in)
}

// Update Receive
func (u *Receive) Update(ctx context.Context, in *inventories.Receive) (*inventories.Receive, error) {
	return u.Client.Update(ctx, in)
}

// View Receive
func (u *Receive) View(ctx context.Context, in *inventories.Id) (*inventories.Receive, error) {
	return u.Client.View(ctx, in)
}

func (u *Receive) OutstandingByPurchase(ctx context.Context, in *inventories.Id) (*inventories.OutstandingResponse, error) {
	return u.Client.OutstandingByPurchase(ctx, in)
}

// List Receives
func (u *Receive) List(in *inventories.ListReceiveRequest, streamServer inventories.ReceiveService_ListServer) error {
	ctx := streamServer.Context()
	streamClient, err := u.Client.List(ctx, in)
	if err != nil {
		return status.Errorf(codes.Internal, "Error when calling Receive.List service: %s", err)
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
