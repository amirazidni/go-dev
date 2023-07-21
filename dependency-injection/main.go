package main

import (
	"go-dev/dependency-injection/helper"
	"go-dev/dependency-injection/middleware"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func NewServer(authMiddleware *middleware.AuthMiddleware) *http.Server {
	return &http.Server{
		Addr:    "localhost:3000",
		Handler: authMiddleware,
	}
}

func main() {
	server := InitializedServer()
	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
