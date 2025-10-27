package auth

import (
	"errors"

	"github.com/gofiber/fiber/v2"

	"github.com/sugito75/chat-app-server/internal/user"
	"github.com/sugito75/chat-app-server/pkg/jwt"
	"github.com/sugito75/chat-app-server/pkg/session"
	"golang.org/x/crypto/bcrypt"
)

type authService struct {
	authRepo       AuthRepository
	userRepo       user.UserRepository
	sessionService session.SessionService
	jwtService     *jwt.JwtService
}

func NewService(authRepo AuthRepository, userRepo user.UserRepository, sessionService session.SessionService, jwtService *jwt.JwtService) AuthService {
	return &authService{
		authRepo:       authRepo,
		userRepo:       userRepo,
		sessionService: sessionService,
		jwtService:     jwtService,
	}
}

func (s *authService) Register(dto RegisterDTO) (uint, error) {
	if user := s.userRepo.GetUserByPhone(dto.Phone); user != nil {
		return 0, fiber.NewError(fiber.StatusBadRequest, "number already in use!")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(dto.Password), 10)
	if err != nil {
		return 0, err
	}

	user := user.User{
		DisplayName:    dto.DisplayName,
		Password:       string(hashedPassword),
		Phone:          dto.Phone,
		Bio:            dto.Bio,
		ProfilePicture: &dto.ProfilePicture,
	}

	uid, err := s.userRepo.CreateUser(user)
	if err != nil {
		return 0, err
	}

	if err := s.sessionService.SaveSession(uid); err != nil {
		return 0, err
	}

	return uid, nil
}

func (s *authService) Login(dto LoginDTO) (*UserCredentialsDTO, error) {
	user := s.userRepo.GetUserByPhone(dto.Phone)
	if user == nil {
		return nil, fiber.NewError(fiber.StatusNotFound, "Invalid phone or password")
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(dto.Password))
	if err != nil {
		return nil, fiber.NewError(fiber.StatusBadRequest, "Invalid phone or password")
	}

	sessionID, err := s.sessionService.GetSessionID(user.ID)
	if err != nil {
		return nil, fiber.NewError(fiber.StatusForbidden, "No session")
	}

	refreshToken, accessToken := s.generateTokens(jwt.UserData{ID: user.ID, Phone: user.Phone})

	if err := s.authRepo.SaveToken(uint64(user.ID), refreshToken); err != nil {
		return nil, err
	}

	return &UserCredentialsDTO{
		ID:             user.ID,
		Username:       user.DisplayName,
		ProfilePicture: user.ProfilePicture,
		Phone:          user.Phone,
		Bio:            user.Bio,
		SessionID:      sessionID,
		RefreshToken:   refreshToken,
		AccessToken:    accessToken,
	}, nil
}

func (s *authService) Verify(token string) error {
	return nil
}

func (s *authService) GenerateAccessToken(refreshToken string) (string, error) {
	user, err := s.jwtService.Verify(refreshToken, false)
	if err != nil {
		if errors.Is(err, jwt.ErrTokenExpired) {
			s.authRepo.RemoveToken(refreshToken)
		}

		return "", err
	}

	if t := s.authRepo.GetToken(uint64(user.ID), refreshToken); t == "" {
		return "", fiber.NewError(fiber.StatusUnauthorized, "invalid token")
	}

	accessToken := s.jwtService.Generate(*user, true)

	return accessToken, nil

}

func (s *authService) Logout(token string) error {
	return s.authRepo.RemoveToken(token)

}

func (s *authService) generateTokens(payload jwt.UserData) (string, string) {
	accessToken := s.jwtService.Generate(payload, true)
	refreshToken := s.jwtService.Generate(payload, false)

	return refreshToken, accessToken
}
