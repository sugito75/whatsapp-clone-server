package auth

import "github.com/sugito75/chat-app-server/internal/user"

type AuthToken struct {
	ID     uint64 `gorm:"primaryKey;autoIncrement"`
	UserID uint64 `gorm:"not null;index"`
	Token  string `gorm:"type:text;not null"`

	User user.User `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE;"`
}
