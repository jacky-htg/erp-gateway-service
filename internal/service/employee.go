package service

import (
	"context"
	users "inventory-gateway-service/pb"
)

// Employee struct
type Employee struct {
	EmployeeClient users.EmployeeServiceClient
	UserClient     users.UserServiceClient
}

// Create Employee
func (u *Employee) Create(ctx context.Context, in *users.Employee) (*users.Employee, error) {
	return &users.Employee{}, nil
}

// Update Employee
func (u *Employee) Update(ctx context.Context, in *users.Employee) (*users.Employee, error) {
	return &users.Employee{}, nil
}

// View Employee
func (u *Employee) View(ctx context.Context, in *users.Id) (*users.Employee, error) {
	return &users.Employee{}, nil
}

// Delete Employee
func (u *Employee) Delete(ctx context.Context, in *users.Id) (*users.Boolean, error) {
	return &users.Boolean{}, nil
}

// List Employee
func (u *Employee) List(in *users.ListEmployeeRequest, stream users.EmployeeService_ListServer) error {
	return nil
}
