package main

import (
	"log/slog"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
	"github.com/sugito75/chat-app-server/config"
	"github.com/sugito75/chat-app-server/internal/chat"
	"github.com/sugito75/chat-app-server/internal/user"
	"github.com/sugito75/chat-app-server/pkg/logger"
)

func main() {
	godotenv.Load()
	logger.InitLogger()

	app := fiber.New(config.NewFiberConfig())
	app.Use(cors.New(cors.ConfigDefault))

	app.Use(logger.LogRequestStart)

	user.RegisterUserRoutes(app)
	chat.RegisterChatRoutes(app)

	app.Use(logger.LogRequestEnd)
	slog.Info("Server is online", "port", os.Getenv("HTTP_PORT"))
	app.Listen(os.Getenv("HTTP_PORT"))
}
