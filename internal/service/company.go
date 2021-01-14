package service

import (
	"context"
	users "inventory-gateway-service/pb"
)

// Company struct
type Company struct {
	CompanyClient users.CompanyServiceClient
	UserClient    users.UserServiceClient
}

// Registration new company
func (u *Company) Registration(ctx context.Context, in *users.CompanyRegistration) (*users.CompanyRegistration, error) {
	return &users.CompanyRegistration{}, nil
}

// Update Company
func (u *Company) Update(ctx context.Context, in *users.Company) (*users.Company, error) {
	return &users.Company{}, nil
}

// View Company
func (u *Company) View(ctx context.Context, in *users.Id) (*users.Company, error) {
	return &users.Company{}, nil
}
