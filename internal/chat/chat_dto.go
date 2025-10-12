package chat

type CreatePrivateChatDTO struct {
	Phone string `json:"phone"`
}

type CreateGroupChatDTO struct {
	Title       string   `form:"title" validate:"required,max=100"`
	Members     []string `form:"members" validate:"required,min=1"`
	Icon        *string  `form:"icon,omitempty"`
	Description *string  `form:"description,omitempty"`
}

type JoinGroupDTO struct {
	GroupId uint `json:"groupId" validate:"required"`
}

type MessageDTO struct {
	Type    ChatType `json:"type" validate:"required"`
	ChatID  uint     `json:"chatId" validate:"required"`
	Message string   `json:"message" validate:"required"`
}

type EditMessageDTO struct {
	MessageID string `json:"messageId" validate:"required"`
	Message   string `json:"message" validate:"required"`
}
