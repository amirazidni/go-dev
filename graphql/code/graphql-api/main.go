package main

import (
	"fmt"
	"graphql-api/config"
	"graphql-api/logger"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/graphql-go/graphql"
	"graphql-api/gql"
	"graphql-api/postgres"
	"graphql-api/server"
)

var (
	c = config.GetConfig()
)

func main() {
	// Initialize our api and return a pointer to our router for http.ListenAndServe
	// and a pointer to our db to defer its closing when main() is finished
	router, db := initializeAPI()
	defer db.Close()

	// Start app based on configuration port and if there's an error, log it and exit
	appPort := fmt.Sprintf(":%+v", c.App.Port)
	logger.Log().Fatal(http.ListenAndServe(appPort, router))
}

func initializeAPI() (*mux.Router, *postgres.Db) {

	config.ConfigureLogger()
	logger.Log().Infof("Starting graphql http service in port %+v", c.App.Port)

	// Create a new gorilla mux router
	router := mux.NewRouter()

	// Create a new connection to our pg database
	db, err := postgres.New(
		postgres.ConnString(c.Db.Host, c.Db.Port, c.Db.User, c.Db.DbName, c.Db.Password),
	)
	if err != nil {
		logger.Log().Fatal(err)
	}

	// Create our root query for graphql
	rootQuery := gql.NewRoot(db)

	// Create a new graphql schema, passing in the the root query
	sc, err := graphql.NewSchema(
		graphql.SchemaConfig{Query: rootQuery.Query},
	)
	if err != nil {
		logger.Log().Println("Error creating schema: ", err)
	}

	// Create a server struct that holds a pointer to our database as well
	// as the address of our graphql schema
	s := server.Server{GqlSchema: &sc}

	// Create the graphql route with a Server method to handle it
	router.HandleFunc("/graphql", s.GraphQL()).Methods(http.MethodPost)

	return router, db
}
