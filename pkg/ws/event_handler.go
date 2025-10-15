package ws

import (
	"encoding/json"

	"github.com/sugito75/chat-app-server/internal/chat"
)

type MessageHandler struct {
	service chat.ChatService
}

func NewMessageHandler(service chat.ChatService) *MessageHandler {
	return &MessageHandler{
		service: service,
	}
}

func (h *MessageHandler) HandlePrivateMessage(e Event, c *Client) error {
	var data map[string]string
	json.Unmarshal(e.Payload, &data)

	body, _ := json.Marshal(data)
	c.manager.clients[data["id"]].conn.WriteMessage(1, body)
	return nil
}

func (h *MessageHandler) HandleGroupMessage(e Event, c *Client) error {
	return nil
}
