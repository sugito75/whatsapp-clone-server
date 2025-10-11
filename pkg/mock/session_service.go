package mock

import "github.com/stretchr/testify/mock"

type MockSessionService struct {
	mock.Mock
}

func (m *MockSessionService) SaveSession(uid uint) error {
	args := m.Called(uid)
	return args.Error(0)
}

func (m *MockSessionService) GetSessionID(uid uint) (string, error) {
	args := m.Called(uid)
	return args.String(0), args.Error(1)
}
