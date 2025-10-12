package chat

type chatService struct {
	repo ChatRepository
}

func NewService(repo ChatRepository) ChatService {
	return &chatService{
		repo: repo,
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
