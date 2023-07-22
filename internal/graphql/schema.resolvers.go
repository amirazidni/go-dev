package graphql

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"go-generic/internal/graphql/generated"
	"go-generic/internal/graphql/model"

	"github.com/google/uuid"
)

// CreateTodo is the resolver for the createTodo field.
func (r *mutationResolver) CreateTodo(ctx context.Context, input model.TodoInput) (*model.Todo, error) {
	todo := &model.Todo{
		Text:   input.Text,
		ID:     uuid.New().String(),
		User:   &model.User{ID: input.UserID, Name: "user " + input.UserID},
		UserID: input.UserID,
	}

	data, err := r.todoRepository.Create(ctx, todo)
	r.mutex.Lock()
	for _, ch := range r.todoChan {
		ch <- data
	}
	r.mutex.Unlock()
	return data, err
}

// UpdateTodo is the resolver for the updateTodo field.
func (r *mutationResolver) UpdateTodo(ctx context.Context, id string, input model.TodoInput) (*model.Todo, error) {
	todo := &model.Todo{
		ID:     id,
		Text:   input.Text,
		User:   &model.User{ID: input.UserID},
		UserID: input.UserID,
	}
	return r.todoRepository.Update(ctx, todo)
}

// DeleteTodo is the resolver for the deleteTodo field.
func (r *mutationResolver) DeleteTodo(ctx context.Context, id string) (*model.Todo, error) {
	return r.todoRepository.Delete(ctx, id)
}

// Todos is the resolver for the todos field.
func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	return r.todoRepository.FindAll(ctx)
}

// AddedTodo is the resolver for the addedTodo field.
func (r *subscriptionResolver) AddedTodo(ctx context.Context) (<-chan *model.Todo, error) {
	token := uuid.New().String()
	mc := make(chan *model.Todo, 1)
	r.mutex.Lock()
	r.todoChan[token] = mc
	r.mutex.Unlock()
	go func() {
		<-ctx.Done()
		r.mutex.Lock()
		delete(r.todoChan, token)
		r.mutex.Unlock()
	}()
	return mc, nil
}

// User is the resolver for the user field.
func (r *todoResolver) User(ctx context.Context, obj *model.Todo) (*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// Subscription returns generated.SubscriptionResolver implementation.
func (r *Resolver) Subscription() generated.SubscriptionResolver { return &subscriptionResolver{r} }

// Todo returns generated.TodoResolver implementation.
func (r *Resolver) Todo() generated.TodoResolver { return &todoResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type subscriptionResolver struct{ *Resolver }
type todoResolver struct{ *Resolver }
