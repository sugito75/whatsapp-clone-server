package user

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sugito75/chat-app-server/config"
)

func RegisterUserRoutes(app *fiber.App) {
	users := app.Group("/users")

	repo := NewRepository(config.GetConn())
	service := NewService(repo)
	handler := NewHandler(service)

	users.Post("/new", handler.CreateUser)
}
