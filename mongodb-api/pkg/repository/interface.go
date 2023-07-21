package repository

import "context"

type RepositoryService interface {
	GetTodos(ctx context.Context) (*[]Todo, error)
	Create(ctx context.Context, data *NewTodo) (*Todo, error)
	Update(ctx context.Context, data *Todo) error
	Delete(ctx context.Context, id string) error
}
