package service

import (
	"context"
	"io"
	"log"

	"github.com/jacky-htg/erp-proto/go/pb/users"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Group struct
type Group struct {
	Client users.GroupServiceClient
	users.UnimplementedGroupServiceServer
}

// Create Group
func (u *Group) Create(ctx context.Context, in *users.Group) (*users.Group, error) {
	return u.Client.Create(ctx, in)
}

// Update Group
func (u *Group) Update(ctx context.Context, in *users.Group) (*users.Group, error) {
	return u.Client.Update(ctx, in)
}

// View Group
func (u *Group) View(ctx context.Context, in *users.Id) (*users.Group, error) {
	return u.Client.View(ctx, in)
}

// Delete Group
func (u *Group) Delete(ctx context.Context, in *users.Id) (*users.MyBoolean, error) {
	return u.Client.Delete(ctx, in)
}

// List Group
func (u *Group) List(in *users.ListGroupRequest, streamServer users.GroupService_ListServer) error {
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

// GrantAccess Group
func (u *Group) GrantAccess(ctx context.Context, in *users.GrantAccessRequest) (*users.Message, error) {
	return u.Client.GrantAccess(ctx, in)
}

// RevokeAccess Group
func (u *Group) RevokeAccess(ctx context.Context, in *users.GrantAccessRequest) (*users.Message, error) {
	return u.Client.RevokeAccess(ctx, in)
}
