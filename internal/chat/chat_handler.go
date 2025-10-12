package chat

import "github.com/gofiber/fiber/v2"

type chatHandler struct {
	service ChatService
}

func NewHandler(service ChatService) ChatHandler {
	return &chatHandler{
		service: service,
	}
}

func (h *chatHandler) CreatePrivateChat(ctx *fiber.Ctx) error {
	return nil
}

func (h *chatHandler) CreateGroupChat(ctx *fiber.Ctx) error {
	return nil
}

func (h *chatHandler) JoinGroupChat(ctx *fiber.Ctx) error {
	return nil
}

func (h *chatHandler) GetChats(ctx *fiber.Ctx) error {
	return nil
}

func (h *chatHandler) GetMessages(ctx *fiber.Ctx) error {
	return nil
}

func (h *chatHandler) SendMessage(ctx *fiber.Ctx) error {
	return nil
}

func (h *chatHandler) ReadMessage(ctx *fiber.Ctx) error {
	return nil
}

func (h *chatHandler) EditMessage(ctx *fiber.Ctx) error {
	return nil
}

func (h *chatHandler) DeleteMessage(ctx *fiber.Ctx) error {
	return nil
}
