package repository

import (
	"context"
	"go-generic/internal/graphql/model"
)

type TodoRepository interface {
	Create(ctx context.Context, todo *model.Todo) (*model.Todo, error)
	FindAll(ctx context.Context) ([]*model.Todo, error)
	Update(ctx context.Context, todo *model.Todo) (*model.Todo, error)
	Delete(ctx context.Context, id string) (*model.Todo, error)
}
