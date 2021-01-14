package service

import (
	"context"
	users "inventory-gateway-service/pb"
)

// PackageFeature struct
type PackageFeature struct {
	PackageFeatureClient users.PackageFeatureServiceClient
	UserClient           users.UserServiceClient
}

// View Package Feature
func (u *PackageFeature) View(ctx context.Context, in *users.Id) (*users.PackageOfFeature, error) {
	return &users.PackageOfFeature{}, nil
}

// List PackageFeature
func (u *PackageFeature) List(in *users.Empty, stream users.PackageFeatureService_ListServer) error {
	return nil
}
