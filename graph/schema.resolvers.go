package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"time"

	"graphql-nosql/graph/generated"
	"graphql-nosql/graph/model"
)

// CreateEmployee is the resolver for the createEmployee field.
func (r *mutationResolver) CreateEmployee(ctx context.Context, input model.NewEmployee) (*model.Employee, error) {
	e := &model.Employee{
		Name:      input.Name,
		Age:       input.Age,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	return r.employeeRepository.Create(ctx, e)
}

// UpdateEmployee is the resolver for the updateEmployee field.
func (r *mutationResolver) UpdateEmployee(ctx context.Context, id string, input model.NewEmployee) (*model.Employee, error) {
	e := &model.Employee{
		ID:        id,
		Name:      input.Name,
		Age:       input.Age,
		UpdatedAt: time.Now(),
	}
	return r.employeeRepository.Update(ctx, e)
}

// DeleteEmployee is the resolver for the deleteEmployee field.
func (r *mutationResolver) DeleteEmployee(ctx context.Context, id string) (*bool, error) {
	deleted := true
	err := r.employeeRepository.Delete(ctx, id)
	if err != nil {
		deleted = false
		return &deleted, err
	}
	return &deleted, nil
}

// FindAllEmployee is the resolver for the findAllEmployee field.
func (r *queryResolver) FindAllEmployee(ctx context.Context) ([]*model.Employee, error) {
	return r.employeeRepository.FindAll(ctx)
}

// NewEmployee is the resolver for the newEmployee field.
func (r *subscriptionResolver) NewEmployee(ctx context.Context) (<-chan *model.Employee, error) {
	// token := uuid.New().String()
	// mc :=
	panic(fmt.Errorf("not implemented: NewEmployee - newEmployee"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// Subscription returns generated.SubscriptionResolver implementation.
func (r *Resolver) Subscription() generated.SubscriptionResolver { return &subscriptionResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type subscriptionResolver struct{ *Resolver }
