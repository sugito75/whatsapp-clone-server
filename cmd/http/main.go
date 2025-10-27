package main

import (
	"log/slog"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
	"github.com/sugito75/chat-app-server/config"
	"github.com/sugito75/chat-app-server/internal/auth"
	"github.com/sugito75/chat-app-server/internal/chat"
	"github.com/sugito75/chat-app-server/internal/user"
	"github.com/sugito75/chat-app-server/pkg/logger"
)

func main() {
	godotenv.Load()
	logger.InitLogger()

	app := fiber.New(config.NewFiberConfig())
	app.Use(cors.New(cors.ConfigDefault))
	app.Static("/", "./public")

	app.Use(logger.LogRequestStart)
	api := app.Group("/api")

	auth.RegisterAuthRoutes(api)
	user.RegisterUserRoutes(api)
	chat.RegisterChatRoutes(api)

	app.Use(logger.LogRequestEnd)
	slog.Info("Server is online", "port", os.Getenv("HTTP_PORT"))
	app.Listen(os.Getenv("HTTP_PORT"))
}
