package service

import (
	users "inventory-gateway-service/pb"
)

// Feature struct
type Feature struct {
	Client users.FeatureServiceClient
}

// List feature
func (u *Feature) List(in *users.Empty, stream users.FeatureService_ListServer) error {
	return nil
}
