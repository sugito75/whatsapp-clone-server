package jwt

import (
	"strings"

	"github.com/gofiber/fiber/v2"
)

func ExtractTokenFromHeaders(ctx *fiber.Ctx) string {
	headers := ctx.Get("Authorization")
	if headers == "" {
		return ""
	}

	token := strings.Split(headers, " ")[1]
	if token == "" {
		return ""
	}

	return token
}
