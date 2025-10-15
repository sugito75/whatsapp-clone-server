package chat

type chatService struct {
	repo ChatRepository
	// mq   mq.MessageQueue
}

func NewService(repo ChatRepository) ChatService {
	return &chatService{
		repo: repo,
		// mq:   mq,
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

func (s *chatService) GetChats(uid uint64) ([]GetChatsDTO, error) {

	result, err := s.repo.GetChats(uid)
	if err != nil {
		return nil, err
	}

	chats := ChatModelToDTO(uid, result)

	return chats, nil
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
