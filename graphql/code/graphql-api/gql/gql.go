package gql

import (
	"github.com/graphql-go/graphql"
	"graphql-api/logger"
)

// ExecuteQuery runs our graphql queries
func ExecuteQuery(query string, schema graphql.Schema) *graphql.Result {
	logger.Log().Infof("Graphql execute query %s", query)
	result := graphql.Do(graphql.Params{
		Schema:        schema,
		RequestString: query,
	})

	if len(result.Errors) > 0 {
		logger.Log().Errorf("Unexpected errors inside ExecuteQuery: %v", result.Errors)
	}
	logger.Log().Infof("Graphql execute query results %+v", result)
	return result
}
