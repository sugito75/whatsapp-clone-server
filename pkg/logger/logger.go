package logger

import (
	"fmt"
	"io"
	"log"
	"log/slog"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
)

func openLogFile() *os.File {
	filename := fmt.Sprintf("./logs/%s.log", time.Now().Format("2006-01-02"))
	flags := os.O_APPEND | os.O_CREATE | os.O_WRONLY

	logFile, err := os.OpenFile(filename, flags, os.FileMode(os.ModePerm))
	if err != nil {
		log.Fatalf("err: %+v", err)
	}

	return logFile
}

func InitLogger() {
	logFile := openLogFile()

	multiWriter := io.MultiWriter(os.Stdout, logFile)

	logger := slog.New(slog.NewJSONHandler(multiWriter, nil))
	slog.SetDefault(logger)
}

func LogRequestStart(ctx *fiber.Ctx) error {
	slog.Info(
		"Incoming Request",
		"method", ctx.Method(),
		"path", ctx.Path(),
		"remote_ip", ctx.IP(),
		"user_agent", ctx.Get("User-Agent"),
	)

	return ctx.Next()
}

func LogRequestEnd(ctx *fiber.Ctx) error {
	slog.Info(
		"Request completed",
		"method", ctx.Method(),
		"path", ctx.Path(),
		"remote_ip", ctx.IP(),
		"status_code", ctx.Response().StatusCode(),
		"duration_ms", ctx.Locals("duration"),
	)

	return nil
}
