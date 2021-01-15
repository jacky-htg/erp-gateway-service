package service

import (
	"context"
	users "inventory-gateway-service/pb"
)

// Branch struct
type Branch struct {
	Client users.BranchServiceClient
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
func (u *Branch) Delete(ctx context.Context, in *users.Id) (*users.Boolean, error) {
	return u.Client.Delete(ctx, in)
}

// List branches
func (u *Branch) List(in *users.ListBranchRequest, stream users.BranchService_ListServer) error {
	return nil
}
