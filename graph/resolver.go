package graph

import (
	"fmt"
	"sync"

	"graphql-nosql/graph/model"
	"graphql-nosql/repository"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.
//go:generate go run github.com/99designs/gqlgen generate

type Resolver struct {
	mutex              sync.Mutex
	notifChan          map[string]chan *model.Employee
	employeeRepository repository.EmployeeRepository
}

func NewResolver() *Resolver {
	return &Resolver{
		notifChan: map[string]chan *model.Employee{},
		mutex:     sync.Mutex{},
		// mongoDB:   database.GetDB(),
	}
}

func (r *Resolver) InjectEmployeeRepository(i repository.EmployeeRepository) {
	if i == nil {
		panic(fmt.Errorf("Employee repository can't be nil"))
	}
	r.employeeRepository = i
}
