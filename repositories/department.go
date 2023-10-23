package repositories

import "context"

//go:generate mockgen -package=repositories -source=department.go -destination=department_mock.go

type DepartmentModel interface {
	ID() string
	Name() string
}

type DepartmentRepository interface {
	GetDepartment(ctx context.Context, id string) (DepartmentModel, error)
	GetDepartments(ctx context.Context, ids []string) ([]DepartmentModel, error)
}
