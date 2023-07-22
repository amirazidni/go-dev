package repository

import (
	"context"
	"go-generic/internal/graphql/model"
	"go-generic/pkg/utils"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type TodoRepositoryImp struct {
	DB *gorm.DB
}

func NewTodoRepository(db *gorm.DB) TodoRepository {
	return &TodoRepositoryImp{
		DB: db,
	}
}

func (t *TodoRepositoryImp) Create(ctx context.Context, todo *model.Todo) (*model.Todo, error) {
	err := t.DB.Create(&todo).Error
	if utils.ErrorHandler(err, "failed to create") {
		return nil, err
	}
	return todo, nil
}
func (t *TodoRepositoryImp) FindAll(ctx context.Context) ([]*model.Todo, error) {
	todos := []*model.Todo{}
	err := t.DB.Find(&todos).Error
	if utils.ErrorHandler(err, "failed to find all") {
		return nil, err
	}
	return todos, nil
}

func (t *TodoRepositoryImp) Update(ctx context.Context, todo *model.Todo) (*model.Todo, error) {
	err := t.DB.Model(&todo).Updates(&todo).Error
	if utils.ErrorHandler(err, "failed to update") {
		return nil, err
	}
	return todo, nil
}

func (t *TodoRepositoryImp) Delete(ctx context.Context, id string) (*model.Todo, error) {
	todo := model.Todo{}
	todo.ID = id
	err := t.DB.Clauses(clause.Returning{}).Delete(&todo).Error
	if utils.ErrorHandler(err, "failed to delete") {
		return nil, err
	}
	return &todo, nil
}
