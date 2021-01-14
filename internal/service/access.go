package service

import (
	"context"
	users "inventory-gateway-service/pb"
)

// Access struct
type Access struct {
	AccessClient users.AccessServiceClient
	UserClient   users.UserServiceClient
}

// List access
func (u *Access) List(ctx context.Context, in *users.Empty) (*users.Access, error) {
	return &users.Access{}, nil
}
