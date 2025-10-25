package chat

import "github.com/gofiber/fiber/v2"

type ChatHandler interface {
	CreatePrivateChat(ctx *fiber.Ctx) error
	CreateGroupChat(ctx *fiber.Ctx) error
	JoinGroupChat(ctx *fiber.Ctx) error
	GetChats(ctx *fiber.Ctx) error
	LeaveGroup(ctx *fiber.Ctx) error
}

type ChatService interface {
	CreatePrivateChat(c CreatePrivateChatDTO) (uint64, error)
	CreateGroupChat(c CreateGroupChatDTO) (uint64, error)
	JoinGroupChat(userPhone string, chatId uint64) error
	LeaveGroup(userPhone string, chatId uint64) error
	GetChats(uid uint64) ([]GetChatsDTO, error)
}

type ChatRepository interface {
	CreateChat(c Chat, phones []string) (uint64, error)
	GetChats(uid uint64) ([]ChatMember, error)
	GetChat(id uint64) *Chat
	AddChatMember(m ChatMember) error
	RemoveChatMember(userPhone string, chatId uint64) error
}
