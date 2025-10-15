package chat

import "github.com/gofiber/fiber/v2"

type ChatHandler interface {
	CreatePrivateChat(ctx *fiber.Ctx) error
	CreateGroupChat(ctx *fiber.Ctx) error
	JoinGroupChat(ctx *fiber.Ctx) error
	GetChats(ctx *fiber.Ctx) error
	GetMessages(ctx *fiber.Ctx) error
	SendMessage(ctx *fiber.Ctx) error
	ReadMessage(ctx *fiber.Ctx) error
	EditMessage(ctx *fiber.Ctx) error
	DeleteMessage(ctx *fiber.Ctx) error
}

type ChatService interface {
	CreatePrivateChat(c CreatePrivateChatDTO) error
	CreateGroupChat(c CreateGroupChatDTO) error
	JoinGroupChat(g JoinGroupDTO) error
	GetChats(uid uint64) ([]GetChatsDTO, error)
	SendMessage(m MessageDTO) error
	ReadMessage(id uint) error
	EditMessage(m EditMessageDTO) error
	DeleteMessage(id uint) error
}

type ChatRepository interface {
	CreateChat(c Chat) (uint64, error)
	GetChats(uid uint64) ([]ChatMember, error)
	AddChatMember(m ChatMember) error
	EditMessage(m Message) error
	DeleteMessage(id uint64) error
}
