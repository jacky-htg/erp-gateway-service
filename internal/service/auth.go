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
	var output *users.Message
	output.Message = "Failed"

	ctx, err := getUserLoginCtx(ctx, u.UserClient)
	if err != nil {
		return output, err
	}

	return u.AuthClient.ChangePassword(setMetadata(ctx), in)
}

// IsAuth service
func (u *Auth) IsAuth(ctx context.Context, in *users.String) (*users.Boolean, error) {
	output := &users.Boolean{Boolean: false}

	ctx, err := getUserLoginCtx(ctx, u.UserClient)
	if err != nil {
		return output, err
	}

	return u.AuthClient.IsAuth(setMetadata(ctx), in)
}
