package chat

import (
	"time"

	"github.com/sugito75/chat-app-server/internal/message"
	"github.com/sugito75/chat-app-server/internal/user"
)

type ChatType string
type ChatRole string

const (
	ChatTypePrivate ChatType = "private"
	ChatTypeGroup   ChatType = "group"
)

const (
	RoleMember ChatRole = "member"
	RoleAdmin  ChatRole = "admin"
)

type Chat struct {
	ID          uint64    `gorm:"primaryKey;autoIncrement" json:"id"`
	ChatType    ChatType  `gorm:"type:varchar(10);not null;check:chat_type IN ('private','group')" json:"chat_type"`
	Title       *string   `gorm:"type:varchar(100)" json:"title,omitempty"`
	Description *string   `gorm:"type:text" json:"description,omitempty"`
	Icon        *string   `gorm:"type:text" json:"icon,omitempty"`
	CreatedAt   time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime" json:"updated_at"`

	// Relations

	Members  []ChatMember      `gorm:"foreignKey:ChatID;constraint:OnDelete:CASCADE" json:"members"`
	Messages []message.Message `gorm:"foreignKey:ChatID;constraint:OnDelete:CASCADE" json:"messages"`
}

type ChatMember struct {
	ID            uint64    `gorm:"primaryKey;autoIncrement" json:"id"`
	ChatID        uint64    `gorm:"not null;index;uniqueIndex:idx_chat_user" json:"chat_id"`
	UserPhone     string    `gorm:"not null;index;uniqueIndex:idx_chat_user" json:"user_phone"`
	Role          ChatRole  `gorm:"type:varchar(10);default:'member';check:role IN ('member','admin')" json:"role"`
	JoinedAt      time.Time `gorm:"autoCreateTime" json:"joined_at"`
	LastMessageID *uint64   `gorm:"default:null" json:"last_read_message_id,omitempty"`

	// Relations
	Chat        *Chat           `gorm:"foreignKey:ChatID;constraint:OnDelete:CASCADE" json:"chat,omitempty"`
	User        user.User       `gorm:"foreignKey:UserPhone;references:Phone;constraint:OnDelete:CASCADE" json:"user"`
	LastMessage message.Message `gorm:"foreignKey:LastMessageID"`
}
