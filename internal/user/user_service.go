package user

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sugito75/chat-app-server/pkg/session"
	"golang.org/x/crypto/bcrypt"
)

type userService struct {
	repo           UserRepository
	sessionService session.SessionService
}

func NewService(repo UserRepository, sessionService session.SessionService) UserService {
	return &userService{
		repo:           repo,
		sessionService: sessionService,
	}
}

func (s *userService) CreateUser(u CreateUserDTO) (uint, error) {
	if user := s.repo.GetUserByPhone(u.Phone); user != nil {
		return 0, fiber.NewError(fiber.StatusBadRequest, "number already in use!")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), 10)
	if err != nil {
		return 0, err
	}

	user := User{
		DisplayName:    u.DisplayName,
		Password:       string(hashedPassword),
		Phone:          u.Phone,
		Bio:            u.Bio,
		ProfilePicture: &u.ProfilePicture,
	}

	uid, err := s.repo.CreateUser(user)
	if err != nil {
		return 0, err
	}

	if err := s.sessionService.SaveSession(uid); err != nil {
		return 0, err
	}

	return uid, nil

}

func (s *userService) GetUserCredentials(u GetUserCredentialsDTO) (*UserCredentialsDTO, error) {
	user := s.repo.GetUserByPhone(u.Phone)
	if user == nil {
		return nil, fiber.NewError(fiber.StatusNotFound, "Invalid phone or password")
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(u.Password))
	if err != nil {
		return nil, fiber.NewError(fiber.StatusBadRequest, "Invalid phone or password")
	}

	sessionID, err := s.sessionService.GetSessionID(user.ID)
	if err != nil {
		return nil, fiber.NewError(fiber.StatusForbidden, "No session")
	}

	creds := UserCredentialsDTO{
		ID:             user.ID,
		Username:       user.DisplayName,
		ProfilePicture: user.ProfilePicture,
		Phone:          user.Phone,
		Bio:            user.Bio,
		SessionID:      sessionID,
	}

	return &creds, nil
}

func (s *userService) CheckIsNumberRegistered(phone string) bool {
	u := s.repo.GetUserByPhone(phone)

	return u != nil
}

func (s *userService) GetUserInfo(p string) (*GetUserInfoDTO, error) {
	user := s.repo.GetUserByPhone(p)
	if user == nil {
		return nil, fiber.NewError(fiber.StatusBadRequest, "user not exists")
	}

	dto := GetUserInfoDTO{
		ID:          user.ID,
		Phone:       user.Phone,
		DisplayName: user.DisplayName,
		Bio:         user.Bio,
		Icon:        user.ProfilePicture,
	}

	return &dto, nil
}
