package service

import (
	"context"
	"erp-gateway/pb/users"
	"io"
	"log"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Branch struct
type Branch struct {
	Client users.BranchServiceClient
	users.UnimplementedBranchServiceServer
}

// Create new branch
func (u *Branch) Create(ctx context.Context, in *users.Branch) (*users.Branch, error) {
	return u.Client.Create(ctx, in)
}

// Update branch
func (u *Branch) Update(ctx context.Context, in *users.Branch) (*users.Branch, error) {
	return u.Client.Update(ctx, in)
}

// View branch
func (u *Branch) View(ctx context.Context, in *users.Id) (*users.Branch, error) {
	return u.Client.View(ctx, in)
}

// Delete branch
func (u *Branch) Delete(ctx context.Context, in *users.Id) (*users.MyBoolean, error) {
	return u.Client.Delete(ctx, in)
}

// List branches
func (u *Branch) List(in *users.ListBranchRequest, streamServer users.BranchService_ListServer) error {
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
