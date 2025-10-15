package mq

type MessageType string

const (
	MessagePrivate MessageType = "private"
	MessageGroup   MessageType = "group"
)

type Message struct {
	Type    MessageType `json:"type"`
	ChatID  uint        `json:"chatId"`
	Content string      `json:"content"`
}
