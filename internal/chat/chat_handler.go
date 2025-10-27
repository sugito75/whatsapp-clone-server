package chat

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/sugito75/chat-app-server/internal/user"
	"github.com/sugito75/chat-app-server/pkg/jwt"
	"github.com/sugito75/chat-app-server/pkg/validator"
)

type chatHandler struct {
	service ChatService
}

func NewHandler(service ChatService) ChatHandler {
	return &chatHandler{
		service: service,
	}
}

func (h *chatHandler) CreatePrivateChat(ctx *fiber.Ctx) error {
	start := time.Now()
	var body CreatePrivateChatDTO

	if err := ctx.BodyParser(&body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	if err := validator.ValidateStruct(body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "chat should contains minimum 2 members")
	}

	id, err := h.service.CreatePrivateChat(body)
	if err != nil {
		return err
	}

	ctx.Locals("duration", time.Since(start).Milliseconds())
	ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "success",
		"data":    fiber.Map{"id": id},
	})

	return ctx.Next()
}

func (h *chatHandler) CreateGroupChat(ctx *fiber.Ctx) error {
	start := time.Now()
	var body CreateGroupChatDTO

	if err := ctx.BodyParser(&body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	if err := validator.ValidateStruct(body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "group should contains minimum 2 members")
	}

	id, err := h.service.CreateGroupChat(body)
	if err != nil {
		return err
	}

	ctx.Locals("duration", time.Since(start).Milliseconds())
	ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "success",
		"data":    fiber.Map{"id": id},
	})

	return ctx.Next()

}

func (h *chatHandler) JoinGroupChat(ctx *fiber.Ctx) error {
	start := time.Now()
	user, ok := ctx.Locals("user").(user.GetUserInfoDTO)
	if !ok {
		return fiber.NewError(fiber.StatusBadRequest, "no user id provided")
	}

	chatId, err := ctx.ParamsInt("id")
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	if err := h.service.JoinGroupChat(user.Phone, uint64(chatId)); err != nil {
		return err
	}

	ctx.Locals("duration", time.Since(start).Milliseconds())
	ctx.Status(fiber.StatusOK).JSON(fiber.Map{"message": "success"})
	return ctx.Next()
}

func (h *chatHandler) GetChats(ctx *fiber.Ctx) error {
	user, ok := ctx.Locals("user").(*jwt.UserData)
	if !ok {
		return jwt.ErrClaimFormat
	}

	chats, _ := h.service.GetChats(user.Phone)

	ctx.Status(fiber.StatusOK).JSON(chats)
	return ctx.Next()
}

func (h *chatHandler) LeaveGroup(ctx *fiber.Ctx) error {
	start := time.Now()
	user, ok := ctx.Locals("user").(user.GetUserInfoDTO)
	if !ok {
		return fiber.NewError(fiber.StatusBadRequest, "no user id provided")
	}

	chatId, err := ctx.ParamsInt("id")
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	if err := h.service.LeaveGroup(user.Phone, uint64(chatId)); err != nil {
		return err
	}

	ctx.Locals("duration", time.Since(start).Milliseconds())
	ctx.Status(fiber.StatusOK).JSON(fiber.Map{"message": "success"})
	return ctx.Next()
}
