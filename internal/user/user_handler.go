package user

import (
	"fmt"
	"mime/multipart"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/sugito75/chat-app-server/pkg/validator"
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

	if err := ctx.BodyParser(&body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	if err := validator.ValidateStruct(body); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(err)
	}

	file, err := ctx.FormFile("profilePicture")
	if err != nil && !isNoKeyError(err) {
		return err
	}

	filepath := handleUploadedFile(ctx, file)
	body.ProfilePicture = filepath

	uid, err := h.service.CreateUser(body)
	if err != nil {
		return err
	}

	ctx.Locals("duration", time.Since(start).Milliseconds())
	ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
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

	if err := validator.ValidateStruct(body); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(err)
	}

	cred, err := h.service.GetUserCredentials(body)
	if err != nil {
		return err
	}

	ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "successfully get user's credentials!",
		"data":    cred,
	})

	ctx.Locals("duration", time.Since(start).Milliseconds())
	return ctx.Next()
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

func isNoKeyError(err error) bool {
	return err.Error() == "there is no uploaded file associated with the given key"
}

func handleUploadedFile(ctx *fiber.Ctx, file *multipart.FileHeader) string {
	if file == nil {
		return ""
	}

	filename := fmt.Sprintf("%d_%s", time.Now().UnixMicro(), file.Filename)
	filepath := fmt.Sprintf("./public/icons/%s", filename)
	err := ctx.SaveFile(file, filepath)
	if err != nil {
		return ""

	}

	return filename
}
