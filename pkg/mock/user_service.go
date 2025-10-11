package mock

import (
	"github.com/stretchr/testify/mock"
	"github.com/sugito75/chat-app-server/internal/user"
)

type MockUserService struct {
	mock.Mock
}

func (m *MockUserService) CreateUser(u user.CreateUserDTO) (uint, error) {
	args := m.Called(u)
	return args.Get(0).(uint), args.Error(1)
}

func (m *MockUserService) GetUserCredentials(u user.GetUserCredentialsDTO) (*user.UserCredentialsDTO, error) {
	args := m.Called(u)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*user.UserCredentialsDTO), args.Error(1)
}
