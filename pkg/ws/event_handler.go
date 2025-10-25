package ws

import (
	"github.com/sugito75/chat-app-server/internal/chat"
)

type MessageHandler struct {
	chatRepo chat.ChatRepository
}

func NewMessageHandler(repo chat.ChatRepository) *MessageHandler {
	return &MessageHandler{
		chatRepo: repo,
	}
}

// func (h *MessageHandler) HandlePrivateMessage(e Event, c *Client) error {
// 	var data chat.MessageDTO
// 	if err := json.Unmarshal(e.Payload, &data); err != nil {
// 		return err
// 	}

// 	if err := h.chatRepo.SaveMessage(&data.Message); err != nil {
// 		log.Printf("%+v", err)
// 		return err
// 	}

// 	client, ok := c.manager.clients[data.To]
// 	if !ok {
// 		return errors.New("client is offline")
// 	}

// 	data.Message.Status.Status = chat.StatusDelivered
// 	body, _ := json.Marshal(data.Message)

// 	if err := client.conn.WriteMessage(websocket.TextMessage, body); err != nil {
// 		slog.Error(err.Error())
// 	}

// 	if err := h.chatRepo.SetMessageStatus(data.Message.ID, chat.StatusDelivered); err != nil {
// 		return err
// 	}

// 	return nil
// }

// func (h *MessageHandler) HandleGroupMessage(e Event, c *Client) error {
// 	socketIds := []string{}
// 	for _, id := range socketIds {
// 		c.manager.clients[id].conn.WriteMessage(websocket.TextMessage, e.Payload)
// 	}

// 	return nil
// }
