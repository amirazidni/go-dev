package database

import (
	"context"
	"go-dev/mongodb-api/pkg/util"
	"net/url"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetDB() *mongo.Database {
	username := "amira"
	password := "mongopassword"

	clientOptions := options.Client().ApplyURI(
		"mongodb+srv://" +
			url.QueryEscape(username) + ":" +
			url.QueryEscape(password) +
			"@cluster0.ykzdv8p.mongodb.net/?retryWrites=true&w=majority")
	client, err := mongo.NewClient(clientOptions)
	util.ErrorHandlerFatal(err, "failed to set up a new mongo client")
	util.ErrorHandlerFatal(client.Connect(context.Background()), "failed to connect mongo database")
	return client.Database("mongo-api")
}
