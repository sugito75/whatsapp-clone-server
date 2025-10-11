package user

import (
	"time"

	"github.com/go-playground/validator/v10"
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

func (h *userHandler) CreateUser(ctx *fiber.Ctx) error {
	start := time.Now()
	var body CreateUserDTO

	validate := validator.New()

	if err := ctx.BodyParser(&body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	if err := validate.Struct(body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	file, _ := ctx.FormFile("profilePicture")
	body.ProfilePicture = file

	uid, err := h.service.CreateUser(body)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	ctx.Locals("duration", time.Since(start).Milliseconds())
	ctx.Status(201).JSON(fiber.Map{
		"message": "successfully create new user!",
		"data": fiber.Map{
			"id": uid,
		},
	})
	return ctx.Next()
}

func (h *userHandler) GetUserCredentials(ctx *fiber.Ctx) error {
	start := time.Now()
	var body GetUserCredentialsDTO

	if err := ctx.BodyParser(&body); err != nil {
		return fiber.NewError(400, err.Error())
	}

	cred, err := h.service.GetUserCredentials(body)
	if err != nil {
		return err
	}

	ctx.Status(201).JSON(fiber.Map{
		"message": "successfully get user's credentials!",
		"data":    cred,
	})

	ctx.Locals("duration", time.Since(start).Milliseconds())
	return ctx.Next()
}
