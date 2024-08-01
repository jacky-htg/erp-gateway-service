package service

import (
	"context"

	"github.com/jacky-htg/erp-proto/go/pb/users"
)

// Company struct
type Company struct {
	Client users.CompanyServiceClient
	users.UnimplementedCompanyServiceServer
}

// Registration new company
func (u *Company) Registration(ctx context.Context, in *users.CompanyRegistration) (*users.CompanyRegistration, error) {
	return u.Client.Registration(ctx, in)
}

// Update Company
func (u *Company) Update(ctx context.Context, in *users.Company) (*users.Company, error) {
	return u.Client.Update(ctx, in)
}

// View Company
func (u *Company) View(ctx context.Context, in *users.Id) (*users.Company, error) {
	return u.Client.View(ctx, in)
}
