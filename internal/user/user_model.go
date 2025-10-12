package user

import (
	"time"
)

type User struct {
	ID             uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	Username       string    `gorm:"type:varchar(100);not null" json:"username"`
	Phone          string    `gorm:"type:char(13);not null;unique" json:"phone"`
	Password       string    `gorm:"type:varchar(250);not null" json:"-"`
	ProfilePicture *string   `gorm:"type:text" json:"profile_picture,omitempty"`
	Bio            string    `gorm:"type:varchar(250)" json:"bio,omitempty"`
	IsOnline       bool      `gorm:"default:false" json:"is_online"`
	LastOnline     time.Time `gorm:"autoCreateTime" json:"last_online"`

	// Chats   []chat.Chat      `gorm:"foreignKey:CreatedBy;constraint:OnDelete:CASCADE" json:"chats,omitempty"`
	// Session *session.Session `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE" json:"session,omitempty"`
}
