package service

import (
	"context"
	"erp-gateway/pb/users"
)

// Access struct
type Access struct {
	Client users.AccessServiceClient
	users.UnimplementedAccessServiceServer
}

// List access
func (u *Access) List(ctx context.Context, in *users.MyEmpty) (*users.Access, error) {
	return u.Client.List(ctx, in)
}
