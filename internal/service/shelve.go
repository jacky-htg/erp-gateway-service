package service

import (
	"context"
	"erp-gateway/pb/inventories"
	"io"
	"log"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Shelve struct {
	Client inventories.ShelveServiceClient
	inventories.UnimplementedShelveServiceServer
}

// Create new shelve
func (u *Shelve) Create(ctx context.Context, in *inventories.Shelve) (*inventories.Shelve, error) {
	return u.Client.Create(ctx, in)
}

// Update shelve
func (u *Shelve) Update(ctx context.Context, in *inventories.Shelve) (*inventories.Shelve, error) {
	return u.Client.Update(ctx, in)
}

// View shelve
func (u *Shelve) View(ctx context.Context, in *inventories.Id) (*inventories.Shelve, error) {
	return u.Client.View(ctx, in)
}

// Delete shelve
func (u *Shelve) Delete(ctx context.Context, in *inventories.Id) (*inventories.MyBoolean, error) {
	return u.Client.Delete(ctx, in)
}

// List Shelves
func (u *Shelve) List(in *inventories.ListShelveRequest, streamServer inventories.ShelveService_ListServer) error {
	ctx := streamServer.Context()
	streamClient, err := u.Client.List(ctx, in)
	if err != nil {
		return status.Errorf(codes.Internal, "Error when calling Shelve.List service: %s", err)
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
