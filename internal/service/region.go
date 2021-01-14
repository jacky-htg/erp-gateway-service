package service

import (
	"context"

	users "inventory-gateway-service/pb"
)

// Region struct
type Region struct {
	RegionClient users.RegionServiceClient
	UserClient   users.UserServiceClient
}

// Create new region
func (u *Region) Create(ctx context.Context, in *users.Region) (*users.Region, error) {
	return &users.Region{}, nil
}

// Update region
func (u *Region) Update(ctx context.Context, in *users.Region) (*users.Region, error) {
	return &users.Region{}, nil
}

// View Region
func (u *Region) View(ctx context.Context, in *users.Id) (*users.Region, error) {
	return &users.Region{}, nil
}

// Delete Region
func (u *Region) Delete(ctx context.Context, in *users.Id) (*users.Boolean, error) {
	return &users.Boolean{}, nil
}

// List Region
func (u *Region) List(in *users.ListRegionRequest, stream users.RegionService_ListServer) error {
	return nil
}
