package service

import (
	"context"
	users "inventory-gateway-service/pb"
)

// PackageFeature struct
type PackageFeature struct {
	Client users.PackageFeatureServiceClient
}

// View Package Feature
func (u *PackageFeature) View(ctx context.Context, in *users.Id) (*users.PackageOfFeature, error) {
	return u.Client.View(ctx, in)
}

// List PackageFeature
func (u *PackageFeature) List(in *users.Empty, stream users.PackageFeatureService_ListServer) error {
	return nil
}
