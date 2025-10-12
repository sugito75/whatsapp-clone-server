package mq

type MessageType string

const (
	MessagePrivate MessageType = "private"
	MessageGroup   MessageType = "group"
)

type Message struct {
	Type    MessageType
	ChatID  uint
	Content string
}
