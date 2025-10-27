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

func (r *chatRepository) CreateChat(c Chat, phones []string) (uint64, error) {
	result := r.db.Create(&c)

	if result.Error != nil {
		return 0, result.Error
	}

	chatMembers := []ChatMember{}
	for _, phone := range phones {
		m := ChatMember{ChatID: c.ID, UserPhone: phone}
		chatMembers = append(chatMembers, m)
	}

	result = r.db.CreateInBatches(chatMembers, len(phones))
	if result.Error != nil {
		return 0, result.Error
	}

	return c.ID, nil
}

// needs optimations
func (r *chatRepository) GetChats(phone string) ([]ChatMember, error) {
	var chats []ChatMember
	result := r.db.
		Preload("LastMessage").
		Preload("LastMessage.Status").
		Preload("Chat").
		Preload("Chat.Members", func(db *gorm.DB) *gorm.DB {
			return db.Where("user_phone != $2", phone).Limit(1)
		}).
		Preload("Chat.Members.User").
		Order("joined_at DESC").
		Find(&chats, "user_phone = $1", phone)

	if result.Error != nil {
		return nil, result.Error
	}

	return chats, nil
}

func (r *chatRepository) GetChat(id uint64) *Chat {
	var chat Chat
	result := r.db.Find(&chat, "id  = $1", id)

	if result.Error != nil {
		return nil
	}

	return &chat
}

func (r *chatRepository) RemoveChatMember(userPhone string, chatId uint64) error {

	result := r.db.Delete(ChatMember{UserPhone: userPhone, ChatID: chatId})

	return result.Error
}

func (r *chatRepository) AddChatMember(m ChatMember) error {
	result := r.db.Create(&m)

	return result.Error
}

func (r *chatRepository) SaveMessage(m *Message) error {
	m.Status = &MessageStatus{
		Status: StatusSent,
	}

	result := r.db.Create(m)
	if result.Error != nil {
		return result.Error
	}

	result = r.db.Table("chat_members").Where("chat_id = $2", m.ChatID).Update("last_message_id", m.ID)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *chatRepository) EditMessage(id uint64, m Message) error {
	return nil
}

func (r *chatRepository) DeleteMessage(id uint64) error {
	return nil
}

func (r chatRepository) SetMessageStatus(id uint64, status ChatStatus) error {
	result := r.db.Table("message_statuses").Where("message_id = $2", id).Update("status", status)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
