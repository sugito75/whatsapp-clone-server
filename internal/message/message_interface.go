package message

import "github.com/sugito75/chat-app-server/pkg/ws"

type MessageHandler interface {
	SendPrivateMessage(evt ws.Event) error
	SendGroupMessage(evt ws.Event) error
	EditMessage(evt ws.Event) error
	DeleteMessage(evt ws.Event) error
}

type MessageService interface {
	SendPrivateMessage() error
	SendGroupMessage() error
	EditMessage() error
	DeleteMessage() error
}

type MessageRepository interface {
	SaveMessage() error
	EditMessage() error
	DeleteMessage() error
}
