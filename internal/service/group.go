package service

import (
	"context"
	users "inventory-gateway-service/pb"
)

// Group struct
type Group struct {
	Client users.GroupServiceClient
}

// Create Group
func (u *Group) Create(ctx context.Context, in *users.Group) (*users.Group, error) {
	return u.Client.Create(ctx, in)
}

// Update Group
func (u *Group) Update(ctx context.Context, in *users.Group) (*users.Group, error) {
	return u.Client.Update(ctx, in)
}

// View Group
func (u *Group) View(ctx context.Context, in *users.Id) (*users.Group, error) {
	return u.Client.View(ctx, in)
}

// Delete Group
func (u *Group) Delete(ctx context.Context, in *users.Id) (*users.Boolean, error) {
	return u.Client.Delete(ctx, in)
}

// List Group
func (u *Group) List(in *users.ListGroupRequest, stream users.GroupService_ListServer) error {
	return nil
}

// GrantAccess Group
func (u *Group) GrantAccess(ctx context.Context, in *users.GrantAccessRequest) (*users.Message, error) {
	return u.Client.GrantAccess(ctx, in)
}

// RevokeAccess Group
func (u *Group) RevokeAccess(ctx context.Context, in *users.GrantAccessRequest) (*users.Message, error) {
	return u.Client.RevokeAccess(ctx, in)
}
