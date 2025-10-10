package middleware

import (
	"context"
	"log/slog"

	"github.com/gofiber/fiber/v2"
)

func ErrorHandler(ctx *fiber.Ctx, err error) error {
	code := fiber.StatusInternalServerError
	logLevel := slog.LevelWarn

	if e, ok := err.(*fiber.Error); ok {
		code = e.Code
	}

	if code > 400 {
		logLevel = slog.LevelError
	}

	slog.Log(
		context.Background(),
		logLevel,
		"Request Error",
		"method", ctx.Method(),
		"path", ctx.Path(),
		"remote_ip", ctx.IP(),
		"user_agent", ctx.Get("User-Agent"),
		"error", err.Error(),
		"status_code", code,
	)

	return ctx.Status(code).JSON(fiber.Map{
		"status":  code,
		"message": err.Error(),
	})
}
