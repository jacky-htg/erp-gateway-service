package service

import (
	"context"
	"erp-gateway/pb/users"
	"io"
	"log"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// User struct
type User struct {
	Client users.UserServiceClient
	users.UnimplementedUserServiceServer
}

// Create func
func (u *User) Create(ctx context.Context, in *users.User) (*users.User, error) {
	return u.Client.Create(ctx, in)
}

// Update func
func (u *User) Update(ctx context.Context, in *users.User) (*users.User, error) {
	return u.Client.Update(ctx, in)
}

// View func
func (u *User) View(ctx context.Context, in *users.Id) (*users.User, error) {
	return u.Client.View(ctx, in)
}

// Delete func
func (u *User) Delete(ctx context.Context, in *users.Id) (*users.MyBoolean, error) {
	return u.Client.Delete(ctx, in)
}

// List func
func (u *User) List(in *users.ListUserRequest, streamServer users.UserService_ListServer) error {
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

// GetByToken func
func (u *User) GetByToken(ctx context.Context, in *users.MyEmpty) (*users.User, error) {
	panic("not implemented")
}
