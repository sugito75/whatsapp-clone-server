package mock

import (
	"github.com/stretchr/testify/mock"
	"github.com/sugito75/chat-app-server/internal/user"
)

type MockUserRepository struct {
	mock.Mock
}

func (m *MockUserRepository) CreateUser(u user.User) (uint, error) {
	args := m.Called(u)
	return args.Get(0).(uint), args.Error(1)
}

func (m *MockUserRepository) GetUserByPhone(phone string) *user.User {
	args := m.Called(phone)
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).(*user.User)
}
