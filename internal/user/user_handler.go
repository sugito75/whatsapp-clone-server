package user

import (
	"time"

	"github.com/gofiber/fiber/v2"
)

type userHandler struct {
	service UserService
}

func NewHandler(service UserService) UserHandler {
	return &userHandler{
		service: service,
	}
}

func (h *userHandler) CheckIsNumberRegistered(ctx *fiber.Ctx) error {
	start := time.Now()
	phone := ctx.Params("phone")

	isRegistered := h.service.CheckIsNumberRegistered(phone)

	ctx.Locals("duration", time.Since(start).Milliseconds())
	ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "success",
		"data": fiber.Map{
			"isRegistered": isRegistered,
		},
	})

	return ctx.Next()
}

func (h *userHandler) GetUserInfo(ctx *fiber.Ctx) error {
	start := time.Now()
	phone := ctx.Params("phone")
	if phone == "" {
		return fiber.NewError(fiber.StatusBadRequest, "phone number is required!")
	}

	u, err := h.service.GetUserInfo(phone)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	ctx.Locals("duration", time.Since(start).Milliseconds())
	ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "success",
		"data":    u,
	})

	return ctx.Next()
}
