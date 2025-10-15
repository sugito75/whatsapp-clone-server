package chat

import (
	"time"

	"github.com/sugito75/chat-app-server/internal/user"
)

type ChatType string
type ChatRole string
type ChatStatus string

const (
	ChatTypePrivate ChatType = "private"
	ChatTypeGroup   ChatType = "group"
)

const (
	RoleMember ChatRole = "member"
	RoleAdmin  ChatRole = "admin"
)

const (
	StatusSent      ChatStatus = "sent"
	StatusDelivered ChatStatus = "delivered"
	StatusReaded    ChatStatus = "readed"
)

type Chat struct {
	ID          uint64    `gorm:"primaryKey;autoIncrement" json:"id"`
	ChatType    ChatType  `gorm:"type:varchar(10);not null;check:chat_type IN ('private','group')" json:"chat_type"`
	Title       *string   `gorm:"type:varchar(100)" json:"title,omitempty"`
	Description *string   `gorm:"type:text" json:"description,omitempty"`
	Icon        *string   `gorm:"type:text" json:"icon,omitempty"`
	CreatedBy   uint64    `gorm:"not null;index" json:"created_by"`
	CreatedAt   time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime" json:"updated_at"`

	// Relations
	Creator  user.User    `gorm:"foreignKey:CreatedBy;constraint:OnDelete:CASCADE" json:"creator"`
	Members  []ChatMember `gorm:"foreignKey:ChatID;constraint:OnDelete:CASCADE" json:"members"`
	Messages []Message    `gorm:"foreignKey:ChatID;constraint:OnDelete:CASCADE" json:"messages"`
}

type ChatMember struct {
	ID                uint64    `gorm:"primaryKey;autoIncrement" json:"id"`
	ChatID            uint64    `gorm:"not null;index;uniqueIndex:idx_chat_user" json:"chat_id"`
	UserID            uint64    `gorm:"not null;index;uniqueIndex:idx_chat_user" json:"user_id"`
	Role              ChatRole  `gorm:"type:varchar(10);default:'member';check:role IN ('member','admin')" json:"role"`
	JoinedAt          time.Time `gorm:"autoCreateTime" json:"joined_at"`
	LastReadMessageID *uint64   `gorm:"default:null" json:"last_read_message_id,omitempty"`

	// Relations
	Chat        *Chat      `gorm:"foreignKey:ChatID;constraint:OnDelete:CASCADE" json:"chat,omitempty"`
	User        *user.User `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE" json:"user,omitempty"`
	LastMessage Message    `gorm:"foreignKey:LastReadMessageID"`
}

type Message struct {
	ID        uint64     `gorm:"primaryKey;autoIncrement" json:"id"`
	ChatID    uint64     `gorm:"index" json:"chat_id"`
	SenderID  *uint64    `gorm:"index" json:"sender_id,omitempty"`
	Content   string     `gorm:"type:text" json:"content,omitempty"`
	ReplyTo   *uint64    `gorm:"index" json:"reply_to,omitempty"`
	SentAt    time.Time  `gorm:"autoCreateTime" json:"sent_at"`
	EditedAt  *time.Time `json:"edited_at,omitempty"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`

	// Relations
	Chat   *Chat          `gorm:"foreignKey:ChatID;constraint:OnDelete:CASCADE" json:"chat,omitempty"`
	Sender *user.User     `gorm:"foreignKey:SenderID;constraint:OnDelete:SET NULL" json:"sender,omitempty"`
	Reply  *Message       `gorm:"foreignKey:ReplyTo;constraint:OnDelete:SET NULL" json:"reply,omitempty"`
	Status *MessageStatus `gorm:"foreignKey:MessageID;constraint:OnDelete:SET NULL" json:"status,omitempty"`
}

type MessageStatus struct {
	ID        uint64     `gorm:"primaryKey;autoIncrement" json:"id"`
	MessageID uint64     `gorm:"not null;index;uniqueIndex:idx_message_user" json:"message_id"`
	UserID    uint64     `gorm:"not null;index;uniqueIndex:idx_message_user" json:"user_id"`
	Status    ChatStatus `gorm:"type:varchar(10);default:'sent';check:status IN ('sent','delivered','readed')" json:"status"`
	UpdatedAt time.Time  `gorm:"autoUpdateTime" json:"updated_at"`

	// Relations
	Message *Message   `gorm:"foreignKey:MessageID;constraint:OnDelete:CASCADE" json:"message,omitempty"`
	User    *user.User `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE" json:"user,omitempty"`
}
