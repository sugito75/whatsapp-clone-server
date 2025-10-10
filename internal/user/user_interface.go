package user

import "github.com/gofiber/fiber/v2"

type UserHandler interface {
	CreateUser(ctx *fiber.Ctx) error
}

type UserService interface {
	CreateUser(u CreateUserDTO) (uint, error)
}

type UserRepository interface {
	CreateUser(u User) (uint, error)
}
