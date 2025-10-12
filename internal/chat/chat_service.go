package chat

import "github.com/sugito75/chat-app-server/pkg/mq"

type chatService struct {
	repo ChatRepository
	mq   mq.MessageQueue
}

func NewService(repo ChatRepository, mq mq.MessageQueue) ChatService {
	return &chatService{
		repo: repo,
		mq:   mq,
	}
}

func (s *chatService) CreatePrivateChat(c CreatePrivateChatDTO) error {
	return nil
}

func (s *chatService) CreateGroupChat(c CreateGroupChatDTO) error {
	return nil
}

func (s *chatService) JoinGroupChat(g JoinGroupDTO) error {
	return nil
}

func (s *chatService) GetChats(uid uint) error {
	return nil
}

func (s *chatService) SendMessage(m MessageDTO) error {
	return nil
}

func (s *chatService) ReadMessage(id uint) error {
	return nil
}

func (s *chatService) EditMessage(m EditMessageDTO) error {
	return nil
}

func (s *chatService) DeleteMessage(id uint) error {
	return nil
}
