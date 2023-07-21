package repository

import (
	"context"
	"fmt"
	"go-dev/mongodb-api/pkg/util"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repository struct {
	MC *mongo.Collection
}

func NewRepository(db *mongo.Database) RepositoryService {
	return &Repository{
		MC: db.Collection("todo"),
	}
}

func (r *Repository) GetTodos(ctx context.Context) (*[]Todo, error) {
	todo := []Todo{}
	cur, err := r.MC.Find(ctx, bson.M{})
	if util.ErrorHandler(err, "failed to find all") {
		return nil, err
	}
	if util.ErrorHandler(cur.All(ctx, &todo), "failed to iterate result") {
		return nil, err
	}
	return &todo, nil
}

func (r *Repository) Create(ctx context.Context, data *NewTodo) (*Todo, error) {
	result, err := r.MC.InsertOne(ctx, data)
	if util.ErrorHandler(err, "failed to create") {
		return nil, err
	}
	todo := Todo{
		ID:   result.InsertedID.(primitive.ObjectID).Hex(),
		Name: data.Name,
	}
	return &todo, nil
}
func (r *Repository) Update(ctx context.Context, data *Todo) error {
	objID, _ := primitive.ObjectIDFromHex(data.ID)
	updated := bson.M{
		"$set": bson.M{"name": data.Name},
	}
	result, err := r.MC.UpdateOne(ctx, bson.M{"_id": objID}, updated)
	if util.ErrorHandler(err, "failed to update") {
		return err
	}
	if result.ModifiedCount < result.MatchedCount || result.MatchedCount == 0 {
		return fmt.Errorf("not found")
	}
	return nil
}
func (r *Repository) Delete(ctx context.Context, id string) error {
	objID, _ := primitive.ObjectIDFromHex(id)
	result, err := r.MC.DeleteOne(ctx, bson.M{"_id": objID})
	if util.ErrorHandler(err, "failed to delete") {
		return err
	}
	if result.DeletedCount < 1 {
		return fmt.Errorf("not found")
	}
	return nil
}
