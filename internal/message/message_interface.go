package message

type MessageService interface {
	SendPrivateMessage() error
	SendGroupMessage() error
	EditMessage() error
	DeleteMessage() error
}

type MessageRepository interface {
	SaveMessage(m Message) error
	EditMessage(m Message) error
	DeleteMessage(id uint64) error
}
