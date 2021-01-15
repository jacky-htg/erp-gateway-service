package service

import (
	"context"

	users "inventory-gateway-service/pb"
)

// Auth struct
type Auth struct {
	AuthClient users.AuthServiceClient
	UserClient users.UserServiceClient
}

// Login service
func (u *Auth) Login(ctx context.Context, in *users.LoginRequest) (*users.LoginResponse, error) {
	return u.AuthClient.Login(ctx, in)
}

// ForgotPassword service
func (u *Auth) ForgotPassword(ctx context.Context, in *users.ForgotPasswordRequest) (*users.Message, error) {
	return u.AuthClient.ForgotPassword(ctx, in)
}

// ResetPassword service
func (u *Auth) ResetPassword(ctx context.Context, in *users.ResetPasswordRequest) (*users.Message, error) {
	return u.AuthClient.ResetPassword(ctx, in)
}

// ChangePassword service
func (u *Auth) ChangePassword(ctx context.Context, in *users.ChangePasswordRequest) (*users.Message, error) {
	return u.AuthClient.ChangePassword(ctx, in)
}

// IsAuth service
func (u *Auth) IsAuth(ctx context.Context, in *users.String) (*users.Boolean, error) {
	return u.AuthClient.IsAuth(ctx, in)
}
