package repositories

import "context"

//go:generate mockgen -package=repositories -source=employee.go -destination=employee_mock.go

type EmployeeModel interface {
	ID() string
	Name() string
	DepartmentID() string
}

type EmployeeWithDepartmentModel interface {
	ID() string
	Name() string
	DepartmentID() string
	DepartmentName() string
}

type EmployeeRepository interface {
	GetEmployee(ctx context.Context, id string) (EmployeeModel, error)
	GetEmployees(ctx context.Context, ids []string) ([]EmployeeModel, error)
	GetEmployeesWithDepartment(ctx context.Context, ids []string) ([]EmployeeWithDepartmentModel, error)
}
