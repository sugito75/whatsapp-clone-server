package user

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sugito75/chat-app-server/config"
	"github.com/sugito75/chat-app-server/pkg/session"
)

func RegisterUserRoutes(app *fiber.App) {
	users := app.Group("/users")

	db := config.GetConn()
	repo := NewRepository(db)
	sessionService := session.NewSessionService(db)

	service := NewService(repo, sessionService)
	handler := NewHandler(service)

	users.Post("/new", handler.CreateUser)
	users.Post("/login", handler.GetUserCredentials)
}
