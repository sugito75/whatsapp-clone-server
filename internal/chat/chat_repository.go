package chat

import (
	"gorm.io/gorm"
)

type chatRepository struct {
	db *gorm.DB
}

func NewRepo(db *gorm.DB) ChatRepository {
	return &chatRepository{
		db: db,
	}
}

func (r *chatRepository) CreateChat(c Chat) (uint64, error) {
	result := r.db.Create(&c)

	if result.Error != nil {
		return 0, result.Error
	}

	return c.ID, nil
}

func (r *chatRepository) GetChats(uid uint64) ([]ChatMember, error) {
	var chats []ChatMember
	result := r.db.
		Preload("LastMessage").
		Preload("LastMessage.Status").
		Preload("Chat").
		Preload("Chat.Members", func(db *gorm.DB) *gorm.DB {
			return db.Where("user_id != $2", uid).Limit(1)
		}).
		Preload("Chat.Members.User").
		Order("created_at DESC").
		Find(&chats, "user_id = $1", uid)

	if result.Error != nil {
		return nil, result.Error
	}

	return chats, nil
}

func (r *chatRepository) AddChatMember(m ChatMember) error {
	result := r.db.Create(&m)

	return result.Error
}

func (r *chatRepository) EditMessage(m Message) error {
	return nil
}

func (r *chatRepository) DeleteMessage(id uint64) error {
	return nil
}
