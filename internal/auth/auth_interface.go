package auth

import "github.com/gofiber/fiber/v2"

type AuthHandler interface {
	Register(ctx *fiber.Ctx) error
	Login(ctx *fiber.Ctx) error
	Logout(ctx *fiber.Ctx) error
	GenerateAccessToken(ctx *fiber.Ctx) error
}

type AuthService interface {
	Register(dto RegisterDTO) (uint, error)
	Login(dto LoginDTO) (*UserCredentialsDTO, error)
	GenerateAccessToken(frToken string) (string, error)
	Logout(token string) error
}

type AuthRepository interface {
	SaveToken(userId uint64, token string) error
	RemoveToken(token string) error
	GetToken(userId uint64, token string) string
}
