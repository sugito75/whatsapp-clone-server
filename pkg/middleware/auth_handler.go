package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sugito75/chat-app-server/pkg/jwt"
)

func Auth(ctx *fiber.Ctx) error {
	tokenString := jwt.ExtractTokenFromHeaders(ctx)
	if tokenString == "" {
		return fiber.NewError(fiber.StatusUnauthorized, "no token provided")
	}

	tokenService := jwt.NewService()

	user, err := tokenService.Verify(tokenString, true)
	if err != nil {
		return fiber.NewError(fiber.StatusUnauthorized, err.Error())
	}

	ctx.Locals("user", user)
	return ctx.Next()
}
