package service

import (
	"context"

	"github.com/jacky-htg/erp-proto/go/pb/users"
)

// Access struct
type Access struct {
	Client users.AccessServiceClient
	users.UnimplementedAccessServiceServer
}

// List access
func (u *Access) List(ctx context.Context, in *users.MyEmpty) (*users.ListAccessResponse, error) {
	return u.Client.List(ctx, in)
}
