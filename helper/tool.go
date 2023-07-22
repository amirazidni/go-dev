//go:build tools
// +build tools

package helper

import (
	_ "github.com/99designs/gqlgen"
	_ "github.com/99designs/gqlgen/graphql/introspection"
)

//  go run github.com/99designs/gqlgen init
