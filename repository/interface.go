package repository

import (
	"context"

	"graphql-nosql/graph/model"
)

type EmployeeRepository interface {
	Create(ctx context.Context, e *model.Employee) (*model.Employee, error)
	FindAll(ctx context.Context) ([]*model.Employee, error)
	Update(ctx context.Context, e *model.Employee) (*model.Employee, error)
	Delete(ctx context.Context, id string) error
}
