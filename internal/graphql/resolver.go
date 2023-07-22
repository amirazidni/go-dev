package graphql

import (
	"fmt"
	"go-generic/internal/graphql/model"
	"go-generic/internal/repository"
	"sync"
)

//go:generate go run github.com/99designs/gqlgen generate

// This file will not be regenerated automatically.

func (r *Resolver) InjectTodoRepository(i repository.TodoRepository) {
	if i == nil {
		panic(fmt.Errorf("todo repository can't be nil"))
	}
	r.todoRepository = i
}

// It serves as dependency injection for your app, add any dependencies you require here.
type Resolver struct {
	mutex          sync.Mutex
	todoChan       map[string]chan *model.Todo
	todoRepository repository.TodoRepository
}

func NewResolver() *Resolver {
	return &Resolver{
		todoChan: map[string]chan *model.Todo{},
		mutex:    sync.Mutex{},
	}
}
