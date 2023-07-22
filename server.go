package main

import (
	"log"
	"net/http"
	"os"

	"graphql-nosql/database"
	"graphql-nosql/graph"
	"graphql-nosql/graph/generated"
	"graphql-nosql/repository"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	r := graph.NewResolver()
	r.InjectEmployeeRepository(repository.NewEmployeeRepository(database.GetDB()))
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: r}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
