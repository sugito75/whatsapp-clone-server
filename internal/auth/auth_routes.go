package auth

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sugito75/chat-app-server/config"
	"github.com/sugito75/chat-app-server/internal/user"
	"github.com/sugito75/chat-app-server/pkg/jwt"
	"github.com/sugito75/chat-app-server/pkg/middleware"
	"github.com/sugito75/chat-app-server/pkg/session"
)

func RegisterAuthRoutes(app fiber.Router) {
	auth := app.Group("/auth")

	db := config.GetConn()
	authRepo := NewRepo(db)
	userRepo := user.NewRepository(db)
	sessionService := session.NewSessionService(db)
	jwtService := jwt.NewService()
	service := NewService(authRepo, userRepo, sessionService, jwtService)
	handler := NewHandler(service)

	auth.Post("/tokens", handler.GenerateAccessToken)
	auth.Post("/register", handler.Register)
	auth.Post("/login", handler.Login)

	auth.Delete("/logout", middleware.Auth, handler.Logout)

}
