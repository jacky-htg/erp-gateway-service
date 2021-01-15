package service

import (
	"context"
	users "inventory-gateway-service/pb"
)

// Access struct
type Access struct {
	Client users.AccessServiceClient
}

// List access
func (u *Access) List(ctx context.Context, in *users.Empty) (*users.Access, error) {
	return u.Client.List(ctx, in)
}
