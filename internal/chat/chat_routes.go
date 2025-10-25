package chat

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sugito75/chat-app-server/config"
)

func RegisterChatRoutes(app *fiber.App) {
	chat := app.Group("/chats")

	db := config.GetConn()
	repo := NewRepo(db)
	service := NewService(repo)
	handler := NewHandler(service)

	chat.Get("/", handler.GetChats)
	chat.Post("/privates", handler.CreatePrivateChat)

	chat.Post("/groups", handler.CreateGroupChat)
	chat.Put("/groups/joins/:id", handler.JoinGroupChat)
	chat.Delete("/groups/leaves/:id", handler.LeaveGroup)

}
