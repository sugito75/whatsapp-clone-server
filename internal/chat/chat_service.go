package chat

import (
	"errors"

	"github.com/gofiber/fiber/v2"
)

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

func (s *chatService) CreatePrivateChat(c CreatePrivateChatDTO) (uint64, error) {
	chat := Chat{ChatType: ChatTypePrivate}

	id, err := s.repo.CreateChat(chat, c.Members)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (s *chatService) CreateGroupChat(c CreateGroupChatDTO) (uint64, error) {
	chat := Chat{
		ChatType:    ChatTypeGroup,
		Title:       &c.Title,
		Icon:        c.Icon,
		Description: c.Description,
	}

	id, err := s.repo.CreateChat(chat, c.Members)
	if err != nil {
		return 0, err
	}

	return id, nil

}

func (s *chatService) JoinGroupChat(userPhone string, groupId uint64) error {
	c := s.repo.GetChat(groupId)

	if c == nil {
		return fiber.NewError(fiber.StatusNotFound, "no group found")
	}

	if c.ChatType != ChatTypeGroup {
		return fiber.NewError(fiber.StatusBadRequest, "cannot join to non-group chat")
	}

	return s.repo.AddChatMember(ChatMember{UserPhone: userPhone, ChatID: groupId})

}

func (s *chatService) GetChats(uid uint64) ([]GetChatsDTO, error) {

	result, err := s.repo.GetChats(uid)
	if err != nil {
		return nil, err
	}

	chats := ChatModelToDTO(uid, result)

	return chats, nil
}

func (s *chatService) LeaveGroup(userPhone string, groupId uint64) error {
	c := s.repo.GetChat(groupId)

	if c == nil {
		return errors.New("no chat found")
	}

	if c.ChatType != ChatTypeGroup {
		return errors.New("cannot leave from non-group chat")
	}
	return s.repo.RemoveChatMember(userPhone, groupId)

}
