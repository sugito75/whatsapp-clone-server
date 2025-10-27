package user

import "github.com/gofiber/fiber/v2"

type UserHandler interface {
	CheckIsNumberRegistered(ctx *fiber.Ctx) error
	GetUserInfo(ctx *fiber.Ctx) error
}

type UserService interface {
	CheckIsNumberRegistered(p string) bool
	GetUserInfo(p string) (*GetUserInfoDTO, error)
}

type UserRepository interface {
	CreateUser(u User) (uint, error)
	GetUserByPhone(phone string) *User
}
