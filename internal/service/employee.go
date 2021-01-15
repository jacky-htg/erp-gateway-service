package service

import (
	"context"
	users "inventory-gateway-service/pb"
)

// Employee struct
type Employee struct {
	Client users.EmployeeServiceClient
}

// Create Employee
func (u *Employee) Create(ctx context.Context, in *users.Employee) (*users.Employee, error) {
	return u.Client.Create(ctx, in)
}

// Update Employee
func (u *Employee) Update(ctx context.Context, in *users.Employee) (*users.Employee, error) {
	return u.Client.Update(ctx, in)
}

// View Employee
func (u *Employee) View(ctx context.Context, in *users.Id) (*users.Employee, error) {
	return u.Client.View(ctx, in)
}

// Delete Employee
func (u *Employee) Delete(ctx context.Context, in *users.Id) (*users.Boolean, error) {
	return u.Client.Delete(ctx, in)
}

// List Employee
func (u *Employee) List(in *users.ListEmployeeRequest, stream users.EmployeeService_ListServer) error {
	return nil
}
