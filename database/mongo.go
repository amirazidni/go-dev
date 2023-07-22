package database

import (
	"context"

	"graphql-nosql/helper"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetDB() *mongo.Database {
	clientOptions := options.Client().ApplyURI("mongodb://root:root@localhost:27017")
	client, err := mongo.NewClient(clientOptions)
	helper.ErrorHandlerFatal(err, "failed to set up a new mongo client")
	helper.ErrorHandlerFatal(client.Connect(context.Background()), "failed to connect mongo database")
	return client.Database("graphql")
}
