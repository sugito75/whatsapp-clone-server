package config

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sugito75/chat-app-server/pkg/middleware"
)

func NewFiberConfig() fiber.Config {
	return fiber.Config{

		DisableStartupMessage: true,
		ErrorHandler:          middleware.ErrorHandler,
	}
}
