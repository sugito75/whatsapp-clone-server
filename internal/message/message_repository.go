package message

import "gorm.io/gorm"

type messageRepository struct {
	db *gorm.DB
}

func NewRepo(db *gorm.DB) MessageRepository {
	return &messageRepository{
		db: db,
	}
}

func (r *messageRepository) SaveMessage(m Message) error {
	result := r.db.Create(&m)
	if result.Error != nil {
		return result.Error
	}

	result = r.db.Table("chat_members").Update("last_message_id", m.ID).Where("chat_id = $1", m.ChatID)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *messageRepository) EditMessage(m Message) error {
	result := r.db.Save(&m)

	return result.Error
}

func (r *messageRepository) DeleteMessage(id uint64) error {
	result := r.db.Delete(&Message{ID: id})

	return result.Error
}
