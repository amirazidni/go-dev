package repository

import (
	"context"
	"fmt"

	"graphql-nosql/graph/model"
	"graphql-nosql/helper"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type EmployeeRepositoryImp struct {
	MC *mongo.Collection
}

func NewEmployeeRepository(db *mongo.Database) EmployeeRepository {
	return &EmployeeRepositoryImp{
		MC: db.Collection("employee"),
	}
}

func (r *EmployeeRepositoryImp) Create(ctx context.Context, e *model.Employee) (*model.Employee, error) {
	result, err := r.MC.InsertOne(ctx, e)
	if helper.ErrorHandler(err, "failed to create") {
		return nil, err
	}
	e.ID = result.InsertedID.(primitive.ObjectID).Hex()
	return e, nil
}
func (r *EmployeeRepositoryImp) FindAll(ctx context.Context) ([]*model.Employee, error) {
	employees := []*model.Employee{}
	cur, err := r.MC.Find(ctx, bson.M{})
	if helper.ErrorHandler(err, "failed to find all") {
		return nil, err
	}
	if helper.ErrorHandler(cur.All(ctx, &employees), "failed to iterate result") {
		return nil, err
	}
	return employees, nil
}
func (r *EmployeeRepositoryImp) Update(ctx context.Context, e *model.Employee) (*model.Employee, error) {
	objID, _ := primitive.ObjectIDFromHex(e.ID)
	updated := bson.M{
		"$set": bson.M{
			"updatedat": e.UpdatedAt,
			"name":      e.Name,
			"age":       e.Age,
		},
	}
	result, err := r.MC.UpdateOne(ctx, bson.M{"_id": objID}, updated)
	if helper.ErrorHandler(err, "failed to update") {
		return nil, err
	}
	if result.ModifiedCount < result.MatchedCount || result.MatchedCount == 0 {
		return nil, fmt.Errorf("not found")
	}
	return e, nil
}

func (r *EmployeeRepositoryImp) Delete(ctx context.Context, id string) error {
	objID, _ := primitive.ObjectIDFromHex(id)
	result, err := r.MC.DeleteOne(ctx, bson.M{"_id": objID})
	if helper.ErrorHandler(err, "failed to delete") {
		return err
	}
	if result.DeletedCount < 1 {
		return fmt.Errorf("not found")
	}
	return nil
}
