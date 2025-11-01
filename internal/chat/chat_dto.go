package chat

import (
	"time"

	"github.com/sugito75/chat-app-server/internal/message"
)

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
	Text     string             `json:"text"`
	Status   message.ChatStatus `json:"status"`
	SentAt   time.Time          `json:"sentAt"`
	SenderID *uint64            `json:"senderId"`
}

type JoinGroupDTO struct {
	GroupId uint `json:"groupId" validate:"required"`
}
