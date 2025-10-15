package ws

import (
	"encoding/json"

	"github.com/gorilla/websocket"
	"github.com/sugito75/chat-app-server/internal/chat"
)

type MessageHandler struct {
}

func NewMessageHandler(service chat.ChatService) *MessageHandler {
	return &MessageHandler{}
}

func (h *MessageHandler) HandlePrivateMessage(e Event, c *Client) error {
	var data map[string]string
	json.Unmarshal(e.Payload, &data)

	body, _ := json.Marshal(data)
	c.manager.clients[data["id"]].conn.WriteMessage(1, body)
	return nil
}

func (h *MessageHandler) HandleGroupMessage(e Event, c *Client) error {
	socketIds := []string{}
	for _, id := range socketIds {
		c.manager.clients[id].conn.WriteMessage(websocket.TextMessage, e.Payload)
	}

	return nil
}
