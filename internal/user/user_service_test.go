package user_test

import (
	"errors"
	"mime/multipart"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"golang.org/x/crypto/bcrypt"

	"github.com/sugito75/chat-app-server/internal/user"
	mocks "github.com/sugito75/chat-app-server/pkg/mock"
)

func TestServiceCreateUser(t *testing.T) {
	t.Run("should success", func(t *testing.T) {
		repo := new(mocks.MockUserRepository)
		sessionSvc := new(mocks.MockSessionService)

		dto := user.CreateUserDTO{
			Username:       "John",
			Password:       "secret123",
			Phone:          "08123456789",
			Bio:            "Hi there!",
			ProfilePicture: nil,
		}

		repo.On("GetUserByPhone", dto.Phone).Return(nil)
		repo.On("CreateUser", mock.Anything).Return(uint(1), nil)
		sessionSvc.On("SaveSession", uint(1)).Return(nil)

		svc := user.NewService(repo, sessionSvc)

		uid, err := svc.CreateUser(dto)

		assert.NoError(t, err)
		assert.Equal(t, uint(1), uid)

		repo.AssertExpectations(t)
		sessionSvc.AssertExpectations(t)
	})

	t.Run("should fail when phone already used", func(t *testing.T) {

		repo := new(mocks.MockUserRepository)
		sessionSvc := new(mocks.MockSessionService)

		dto := user.CreateUserDTO{Username: "John", Password: "secret123", Phone: "081212121", ProfilePicture: &multipart.FileHeader{}}

		repo.On("GetUserByPhone", dto.Phone).Return(&user.User{})

		svc := user.NewService(repo, sessionSvc)

		uid, err := svc.CreateUser(dto)

		assert.Error(t, err)
		assert.Equal(t, uint(0), uid)
		repo.AssertExpectations(t)
	})

	t.Run("should fail when something wrong in repo", func(t *testing.T) {

		repo := new(mocks.MockUserRepository)
		sessionSvc := new(mocks.MockSessionService)

		dto := user.CreateUserDTO{Username: "John", Password: "secret123", ProfilePicture: &multipart.FileHeader{}}

		repo.On("GetUserByPhone", dto.Phone).Return(nil)
		repo.On("CreateUser", mock.Anything).Return(uint(0), errors.New("db error"))

		svc := user.NewService(repo, sessionSvc)

		uid, err := svc.CreateUser(dto)

		assert.Error(t, err)
		assert.Equal(t, uint(0), uid)
		repo.AssertExpectations(t)

	})

	t.Run("should fail when something  wrong in session service", func(t *testing.T) {

		repo := new(mocks.MockUserRepository)
		sessionSvc := new(mocks.MockSessionService)

		dto := user.CreateUserDTO{Username: "John", Password: "secret123", ProfilePicture: &multipart.FileHeader{}}

		repo.On("GetUserByPhone", dto.Phone).Return(nil)
		repo.On("CreateUser", mock.Anything).Return(uint(1), nil)
		sessionSvc.On("SaveSession", uint(1)).Return(errors.New("session error"))

		svc := user.NewService(repo, sessionSvc)

		uid, err := svc.CreateUser(dto)

		assert.Error(t, err)
		assert.Equal(t, uint(0), uid)
		repo.AssertExpectations(t)
		sessionSvc.AssertExpectations(t)

	})
}

func TestGetUserCredentials(t *testing.T) {
	t.Run("should success", func(t *testing.T) {
		repo := new(mocks.MockUserRepository)
		sessionSvc := new(mocks.MockSessionService)

		hashed, _ := bcrypt.GenerateFromPassword([]byte("secret123"), 10)
		mockUser := &user.User{
			ID:             1,
			Username:       "John",
			Password:       string(hashed),
			Phone:          "08123456789",
			Bio:            "Hi there!",
			ProfilePicture: nil,
		}

		repo.On("GetUserByPhone", "08123456789").Return(mockUser)
		sessionSvc.On("GetSessionID", uint(1)).Return("session-xyz", nil)

		svc := user.NewService(repo, sessionSvc)

		creds, err := svc.GetUserCredentials(user.GetUserCredentialsDTO{
			Phone:    "08123456789",
			Password: "secret123",
		})

		assert.NoError(t, err)
		assert.Equal(t, uint(1), creds.ID)
		assert.Equal(t, "session-xyz", creds.SessionID)
		repo.AssertExpectations(t)
		sessionSvc.AssertExpectations(t)
	})

	t.Run("should fail when no user found", func(t *testing.T) {
		repo := new(mocks.MockUserRepository)
		sessionSvc := new(mocks.MockSessionService)

		repo.On("GetUserByPhone", "08123456789").Return(nil)

		svc := user.NewService(repo, sessionSvc)

		creds, err := svc.GetUserCredentials(user.GetUserCredentialsDTO{
			Phone:    "08123456789",
			Password: "whatever",
		})

		assert.Nil(t, creds)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "Invalid phone or password")
		repo.AssertExpectations(t)
	})

	t.Run("should fail when password not match", func(t *testing.T) {
		repo := new(mocks.MockUserRepository)
		sessionSvc := new(mocks.MockSessionService)

		hashed, _ := bcrypt.GenerateFromPassword([]byte("secret123"), 10)
		mockUser := &user.User{Password: string(hashed)}

		repo.On("GetUserByPhone", "08123456789").Return(mockUser)

		svc := user.NewService(repo, sessionSvc)

		creds, err := svc.GetUserCredentials(user.GetUserCredentialsDTO{
			Phone:    "08123456789",
			Password: "wrongpass",
		})

		assert.Nil(t, creds)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "Invalid phone or password")
		repo.AssertExpectations(t)
	})

	t.Run("should fail when something wrong in session service", func(t *testing.T) {
		repo := new(mocks.MockUserRepository)
		sessionSvc := new(mocks.MockSessionService)

		hashed, _ := bcrypt.GenerateFromPassword([]byte("secret123"), 10)
		mockUser := &user.User{ID: 1, Password: string(hashed), Phone: "08123456789"}

		repo.On("GetUserByPhone", "08123456789").Return(mockUser)
		sessionSvc.On("GetSessionID", uint(1)).Return("", errors.New("no session"))

		svc := user.NewService(repo, sessionSvc)

		creds, err := svc.GetUserCredentials(user.GetUserCredentialsDTO{
			Phone:    "08123456789",
			Password: "secret123",
		})

		assert.Nil(t, creds)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "No session")
		repo.AssertExpectations(t)
		sessionSvc.AssertExpectations(t)
	})
}
