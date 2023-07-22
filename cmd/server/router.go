package server

import (
	"go-mongo-fiber/pkg/handler"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

type Server struct {
	App     *fiber.App
	Handler handler.HandlerService
}

func (s *Server) route() {
	s.App.Get("/", s.Handler.Get)
	s.App.Post("/", s.Handler.Post)
	s.App.Put("/:id", s.Handler.Put)
	s.App.Delete("/:id", s.Handler.Delete)
}

func (s *Server) StartServer() {
	log.Default().Println("Starting API server")
	s.App = fiber.New()
	s.Handler = handler.NewHandler()
	s.App.Use(logger.New())
	s.route()
	s.App.Listen("0.0.0.0:3000")
}
