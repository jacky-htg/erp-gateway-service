package service

import (
	"context"
	"io"
	"log"

	"erp-gateway/pb/users"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Region struct
type Region struct {
	Client users.RegionServiceClient
}

// Create new region
func (u *Region) Create(ctx context.Context, in *users.Region) (*users.Region, error) {
	return u.Client.Create(ctx, in)
}

// Update region
func (u *Region) Update(ctx context.Context, in *users.Region) (*users.Region, error) {
	return u.Client.Update(ctx, in)
}

// View Region
func (u *Region) View(ctx context.Context, in *users.Id) (*users.Region, error) {
	return u.Client.View(ctx, in)
}

// Delete Region
func (u *Region) Delete(ctx context.Context, in *users.Id) (*users.MyBoolean, error) {
	return u.Client.Delete(ctx, in)
}

// List Region
func (u *Region) List(in *users.ListRegionRequest, streamServer users.RegionService_ListServer) error {
	ctx := streamServer.Context()
	streamClient, err := u.Client.List(ctx, in)
	if err != nil {
		return status.Errorf(codes.Internal, "Error when calling User.List service: %s", err)
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
