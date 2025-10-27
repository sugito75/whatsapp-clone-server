package user

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sugito75/chat-app-server/config"
	"github.com/sugito75/chat-app-server/pkg/middleware"
	"github.com/sugito75/chat-app-server/pkg/session"
)

func RegisterUserRoutes(app fiber.Router) {
	users := app.Group("/users")

	db := config.GetConn()
	repo := NewRepository(db)
	sessionService := session.NewSessionService(db)

	service := NewService(repo, sessionService)
	handler := NewHandler(service)

	users.Use(middleware.Auth)

	users.Get("/check/:phone", handler.CheckIsNumberRegistered)
	users.Get("/info/:phone", handler.GetUserInfo)
}
