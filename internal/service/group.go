package service

import (
	"context"
	users "inventory-gateway-service/pb"
)

// Group struct
type Group struct {
	GroupClient users.GroupServiceClient
	UserClient  users.UserServiceClient
}

// Create Group
func (u *Group) Create(ctx context.Context, in *users.Group) (*users.Group, error) {
	return &users.Group{}, nil
}

// Update Group
func (u *Group) Update(ctx context.Context, in *users.Group) (*users.Group, error) {
	return &users.Group{}, nil
}

// View Group
func (u *Group) View(ctx context.Context, in *users.Id) (*users.Group, error) {
	return &users.Group{}, nil
}

// Delete Group
func (u *Group) Delete(ctx context.Context, in *users.Id) (*users.Boolean, error) {
	return &users.Boolean{}, nil
}

// List Group
func (u *Group) List(in *users.ListGroupRequest, stream users.GroupService_ListServer) error {
	return nil
}

// GrantAccess Group
func (u *Group) GrantAccess(ctx context.Context, in *users.GrantAccessRequest) (*users.Message, error) {
	return &users.Message{}, nil
}

// RevokeAccess Group
func (u *Group) RevokeAccess(ctx context.Context, in *users.GrantAccessRequest) (*users.Message, error) {
	return &users.Message{}, nil
}
