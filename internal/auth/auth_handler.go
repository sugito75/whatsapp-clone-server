package auth

import (
	"fmt"
	"mime/multipart"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/sugito75/chat-app-server/pkg/validator"
)

type authHandler struct {
	authService AuthService
}

func NewHandler(authService AuthService) AuthHandler {
	return &authHandler{
		authService: authService,
	}
}

func (h *authHandler) Register(ctx *fiber.Ctx) error {
	start := time.Now()
	var body RegisterDTO

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

	uid, err := h.authService.Register(body)
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

func (h *authHandler) Login(ctx *fiber.Ctx) error {
	start := time.Now()
	var body LoginDTO

	if err := ctx.BodyParser(&body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	if err := validator.ValidateStruct(body); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(err)
	}

	cred, err := h.authService.Login(body)
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

func (h *authHandler) Logout(ctx *fiber.Ctx) error {
	var body RefreshTokenDTO
	if err := ctx.BodyParser(&body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	if err := validator.ValidateStruct(body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "refreh token is required")
	}

	if err := h.authService.Logout(body.RefreshToken); err != nil {
		return err
	}
	ctx.Status(fiber.StatusOK).JSON(fiber.Map{"message": "success logout"})
	return ctx.Next()
}

func (h *authHandler) GenerateAccessToken(ctx *fiber.Ctx) error {
	var body RefreshTokenDTO
	if err := ctx.BodyParser(&body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	if err := validator.ValidateStruct(body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "refreh token is required")
	}

	accessToken, err := h.authService.GenerateAccessToken(body.RefreshToken)
	if err != nil {
		return err
	}
	ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "success",
		"data":    fiber.Map{"accessToken": accessToken},
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
