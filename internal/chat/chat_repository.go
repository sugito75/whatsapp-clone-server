package chat

import "gorm.io/gorm"

type chatRepository struct {
	db *gorm.DB
}

func NewRepo(db *gorm.DB) ChatRepository {
	return &chatRepository{
		db: db,
	}
}

func (r *chatRepository) CreateChat(c Chat) (uint64, error) {
	return 0, nil
}

func (r *chatRepository) GetChats(uid uint64) ([]Chat, error) {
	return nil, nil
}

func (r *chatRepository) AddChatMember(uid uint64, chatID uint64) error {
	return nil
}

func (r *chatRepository) EditMessage(m Message) error {
	return nil
}

func (r *chatRepository) DeleteMessage(id uint64) error {
	return nil
}
