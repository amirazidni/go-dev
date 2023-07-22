package graphql

import (
	"go-generic/internal/database"
	"go-generic/internal/graphql"
	"go-generic/internal/graphql/generated"
	"go-generic/internal/graphql/model"
	"go-generic/internal/repository"
	"go-generic/pkg/utils"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

const defaultPort = "8080"

type Server struct{}

func newGraphQLHandler() *handler.Server {
	// define resolver content
	// r := &graphql.Resolver{}
	r := graphql.NewResolver()
	r.InjectTodoRepository(repository.NewTodoRepository(database.DbOrm()))
	config := generated.Config{Resolvers: r}
	newSchema := generated.NewExecutableSchema(config)
	return handler.NewDefaultServer(newSchema)
}

func (s *Server) Start() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	s.initDB()
	// logger.InitLogger()
	gqlServer := newGraphQLHandler()

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", gqlServer)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func (s *Server) initDB() {
	dbOrm := database.DbOrm()
	err := dbOrm.AutoMigrate(
		&model.Todo{},
		&model.User{},
	)
	utils.ErrorHandlerFatal(err, "failed to migrate")
}
