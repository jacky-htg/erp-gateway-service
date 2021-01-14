package service

import (
	"context"
	users "inventory-gateway-service/pb"
)

// Branch struct
type Branch struct {
	BranchClient users.BranchServiceClient
	UserClient   users.UserServiceClient
}

// Create new branch
func (u *Branch) Create(ctx context.Context, in *users.Branch) (*users.Branch, error) {
	return &users.Branch{}, nil
}

// Update branch
func (u *Branch) Update(ctx context.Context, in *users.Branch) (*users.Branch, error) {
	return &users.Branch{}, nil
}

// View branch
func (u *Branch) View(ctx context.Context, in *users.Id) (*users.Branch, error) {
	return &users.Branch{}, nil
}

// Delete branch
func (u *Branch) Delete(ctx context.Context, in *users.Id) (*users.Boolean, error) {
	return &users.Boolean{}, nil
}

// List branches
func (u *Branch) List(in *users.ListBranchRequest, stream users.BranchService_ListServer) error {
	return nil
}
