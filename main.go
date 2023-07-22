package main

import (
	"go-generic/cmd/graphql"
)

func main() {
	// run graphql server here
	server := graphql.Server{}
	server.Start()
}

// go run github.com/99designs/gqlgen generate
// go generate ./...
