package handler

import (
	db "go-mongo-fiber/internal/database"
	"go-mongo-fiber/pkg/repository"
	repo "go-mongo-fiber/pkg/repository"
	"go-mongo-fiber/pkg/util"

	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	Repository repo.RepositoryService
}

func NewHandler() HandlerService {
	return &Handler{
		Repository: repo.NewRepository(db.GetDB()),
	}
}

func (h *Handler) Get(c *fiber.Ctx) error {
	todos, err := h.Repository.GetTodos(c.Context())
	util.ErrorHandler(err, "failed to get todos")
	return c.JSON(fiber.Map{"data": todos})
}

func (h *Handler) Post(c *fiber.Ctx) error {
	newTodo := repository.NewTodo{}
	if err := c.BodyParser(&newTodo); err != nil {
		util.ErrorHandler(err, "failed to parse data")
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
	}
	todo, err := h.Repository.Create(c.Context(), &newTodo)
	util.ErrorHandler(err, "failed to create data")
	return c.JSON(fiber.Map{"data": todo})
}

func (h *Handler) Put(c *fiber.Ctx) error {
	id := c.Params("id")
	updateTodo := repository.NewTodo{}
	if err := c.BodyParser(&updateTodo); err != nil {
		util.ErrorHandler(err, "failed to parse data")
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
	}
	todo := repository.Todo{
		ID:   id,
		Name: updateTodo.Name,
	}
	err := h.Repository.Update(c.Context(), &todo)
	util.ErrorHandler(err, "failed to update data")
	return c.JSON(fiber.Map{"message": "success to update"})
}

func (h *Handler) Delete(c *fiber.Ctx) error {
	id := c.Params("id")
	err := h.Repository.Delete(c.Context(), id)
	util.ErrorHandler(err, "failed to delete data")
	return c.JSON(fiber.Map{"message": "success to delete"})
}
