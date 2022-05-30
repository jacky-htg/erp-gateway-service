package service

import (
	"context"

	"erp-gateway/pb/users"
)

// Auth struct
type Auth struct {
	Client users.AuthServiceClient
	users.UnimplementedAuthServiceServer
}

// Login service
func (u *Auth) Login(ctx context.Context, in *users.LoginRequest) (*users.LoginResponse, error) {
	return u.Client.Login(ctx, in)
}

// ForgotPassword service
func (u *Auth) ForgotPassword(ctx context.Context, in *users.ForgotPasswordRequest) (*users.Message, error) {
	return u.Client.ForgotPassword(ctx, in)
}

// ResetPassword service
func (u *Auth) ResetPassword(ctx context.Context, in *users.ResetPasswordRequest) (*users.Message, error) {
	return u.Client.ResetPassword(ctx, in)
}

// ChangePassword service
func (u *Auth) ChangePassword(ctx context.Context, in *users.ChangePasswordRequest) (*users.Message, error) {
	return u.Client.ChangePassword(ctx, in)
}

// IsAuth service
func (u *Auth) IsAuth(ctx context.Context, in *users.MyString) (*users.MyBoolean, error) {
	return u.Client.IsAuth(ctx, in)
}
