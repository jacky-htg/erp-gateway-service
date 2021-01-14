package service

import (
	"context"
	users "inventory-gateway-service/pb"
)

// User struct
type User struct {
	UserClient users.UserServiceClient
}

// Create func
func (u *User) Create(ctx context.Context, in *users.User) (*users.User, error) {
	return &users.User{}, nil
}

// Update func
func (u *User) Update(ctx context.Context, in *users.User) (*users.User, error) {
	return &users.User{}, nil
}

// View func
func (u *User) View(ctx context.Context, in *users.Id) (*users.User, error) {
	return &users.User{}, nil
}

// Delete func
func (u *User) Delete(ctx context.Context, in *users.Id) (*users.Boolean, error) {
	return &users.Boolean{}, nil
}

// List func
func (u *User) List(in *users.ListUserRequest, stream users.UserService_ListServer) error {
	return nil
}

// GetByToken func
func (u *User) GetByToken(ctx context.Context, in *users.Empty) (*users.User, error) {
	panic("not implemented")
}
