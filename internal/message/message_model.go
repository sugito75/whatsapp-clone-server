package message

import (
	"time"

	"github.com/sugito75/chat-app-server/internal/user"
)

type ChatStatus string

const (
	StatusSent      ChatStatus = "sent"
	StatusDelivered ChatStatus = "delivered"
	StatusReaded    ChatStatus = "readed"
)

type Message struct {
	ID        uint64     `gorm:"primaryKey;autoIncrement" json:"id,omitempty"`
	ChatID    uint64     `gorm:"index" json:"chatId"`
	SenderID  *uint64    `gorm:"index" json:"senderId"`
	Content   string     `gorm:"type:text" json:"content"`
	ReplyTo   *uint64    `gorm:"index" json:"reply_to,omitempty"`
	SentAt    time.Time  `gorm:"autoCreateTime" json:"sentAt"`
	EditedAt  *time.Time `json:"edited_at,omitempty"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`

	// Relations

	Sender *user.User     `gorm:"foreignKey:SenderID;constraint:OnDelete:SET NULL" json:"sender,omitempty"`
	Reply  *Message       `gorm:"foreignKey:ReplyTo;constraint:OnDelete:SET NULL" json:"reply,omitempty"`
	Status *MessageStatus `gorm:"foreignKey:MessageID;constraint:OnDelete:SET NULL" json:"status,omitempty"`
}

type MessageStatus struct {
	ID        uint64     `gorm:"primaryKey;autoIncrement" json:"id"`
	MessageID uint64     `gorm:"not null;index;uniqueIndex:idx_message_user" json:"message_id"`
	Status    ChatStatus `gorm:"type:varchar(10);default:'sent';check:status IN ('sent','delivered','readed')" json:"status"`
	UpdatedAt time.Time  `gorm:"autoUpdateTime" json:"updated_at"`

	// Relations
	Message *Message `gorm:"foreignKey:MessageID;constraint:OnDelete:CASCADE" json:"message,omitempty"`
}
