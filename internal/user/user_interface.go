package user

import "github.com/gofiber/fiber/v2"

type UserHandler interface {
	CreateUser(ctx *fiber.Ctx) error
	GetUserCredentials(ctx *fiber.Ctx) error
	CheckIsNumberRegistered(ctx *fiber.Ctx) error
}

type UserService interface {
	CreateUser(u CreateUserDTO) (uint, error)
	GetUserCredentials(u GetUserCredentialsDTO) (*UserCredentialsDTO, error)
	CheckIsNumberRegistered(p string) bool
}

type UserRepository interface {
	CreateUser(u User) (uint, error)
	GetUserByPhone(phone string) *User
}
