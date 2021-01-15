package service

import (
	"context"
	users "inventory-gateway-service/pb"
	"io"
	"log"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Employee struct
type Employee struct {
	Client users.EmployeeServiceClient
}

// Create Employee
func (u *Employee) Create(ctx context.Context, in *users.Employee) (*users.Employee, error) {
	return u.Client.Create(ctx, in)
}

// Update Employee
func (u *Employee) Update(ctx context.Context, in *users.Employee) (*users.Employee, error) {
	return u.Client.Update(ctx, in)
}

// View Employee
func (u *Employee) View(ctx context.Context, in *users.Id) (*users.Employee, error) {
	return u.Client.View(ctx, in)
}

// Delete Employee
func (u *Employee) Delete(ctx context.Context, in *users.Id) (*users.Boolean, error) {
	return u.Client.Delete(ctx, in)
}

// List Employee
func (u *Employee) List(in *users.ListEmployeeRequest, streamServer users.EmployeeService_ListServer) error {
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
