package server

import (
	"encoding/json"
	"graphql-api/logger"
	"net/http"

	"github.com/graphql-go/graphql"
	"graphql-api/gql"
)

// Server will hold connection to the db as well as handlers
type Server struct {
	GqlSchema *graphql.Schema
}

type reqBody struct {
	Query string `json:"query"`
}

// GraphQL returns an http.HandlerFunc for our /graphql endpoint
func (s *Server) GraphQL() http.HandlerFunc {
	logger.Log().Infof("Construct http handler function")
	return func(w http.ResponseWriter, r *http.Request) {
		logger.Log().Infof("========== Start Handling Request: %s - %s ==========", r.URL, r.Method)
		logger.Log().Info("Check to ensure query was provided in the request body")
		if r.Body == nil {
			logger.Log().Errorf("Received body is empty!")
			http.Error(w, "Must provide graphql query in request body", 400)
			return
		}

		var rBody reqBody

		// Decode the request body into rBody
		err := json.NewDecoder(r.Body).Decode(&rBody)
		if err != nil {
			logger.Log().Errorf("Error parsing JSON request body: %+v", err)
			http.Error(w, "Error parsing JSON request body", 400)
		}

		// Execute graphql query
		result := gql.ExecuteQuery(rBody.Query, *s.GqlSchema)

		logger.Log().Info("Writing body response")
		resp, err := json.Marshal(result)
		w.WriteHeader(http.StatusOK)
		w.Write(resp)
		logger.Log().Infof("========== End Handling Request: %s - %s ==========", r.URL, r.Method)
	}
}
