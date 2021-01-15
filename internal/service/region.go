package service

import (
	"context"

	users "inventory-gateway-service/pb"
)

// Region struct
type Region struct {
	Client users.RegionServiceClient
}

// Create new region
func (u *Region) Create(ctx context.Context, in *users.Region) (*users.Region, error) {
	return u.Client.Create(ctx, in)
}

// Update region
func (u *Region) Update(ctx context.Context, in *users.Region) (*users.Region, error) {
	return u.Client.Update(ctx, in)
}

// View Region
func (u *Region) View(ctx context.Context, in *users.Id) (*users.Region, error) {
	return u.Client.View(ctx, in)
}

// Delete Region
func (u *Region) Delete(ctx context.Context, in *users.Id) (*users.Boolean, error) {
	return u.Client.Delete(ctx, in)
}

// List Region
func (u *Region) List(in *users.ListRegionRequest, stream users.RegionService_ListServer) error {
	return nil
}
