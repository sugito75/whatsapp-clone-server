package chat

import "time"

type CreatePrivateChatDTO struct {
	Members []string `json:"members" validate:"required,min=2"`
}

type CreateGroupChatDTO struct {
	Title       string   `form:"title" validate:"required,max=100"`
	Members     []string `form:"members" validate:"required,min=2"`
	Icon        *string  `form:"icon,omitempty"`
	Description *string  `form:"description,omitempty"`
}

type GetChatsDTO struct {
	ID          uint64         `json:"id"`
	ChatType    ChatType       `json:"type"`
	Title       *string        `json:"title,omitempty"`
	Icon        *string        `json:"icon"`
	LastMessage LastMessageDTO `json:"lastMessage"`
}

type LastMessageDTO struct {
	Text     string     `json:"text"`
	Status   ChatStatus `json:"status"`
	SentAt   time.Time  `json:"sentAt"`
	SenderID *uint64    `json:"senderId"`
}

type JoinGroupDTO struct {
	GroupId uint `json:"groupId" validate:"required"`
}

type SendMessageDTO struct {
	Type    ChatType `json:"type" validate:"required"`
	ChatID  uint     `json:"chatId" validate:"required"`
	Message string   `json:"message" validate:"required"`
}

type MessageDTO struct {
	To      string  `json:"to"`
	Message Message `json:"message"`
}

type EditMessageDTO struct {
	MessageID string `json:"messageId" validate:"required"`
	Message   string `json:"message" validate:"required"`
}
