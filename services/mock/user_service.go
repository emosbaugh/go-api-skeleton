package mock

import (
	"fmt"
	"time"

	"github.com/replicatedcom/gin-example/models"

	"github.com/stretchr/testify/mock"
)

type MockUserService struct {
	mock.Mock
}

func (m *MockUserService) Get(email string) (*models.User, error) {
	args := m.Called(email)
	return ArgsUser(args, 0), args.Error(1)
}

func (m *MockUserService) List() ([]*models.User, error) {
	args := m.Called()
	return ArgsUserSlice(args, 0), args.Error(1)
}

func (m *MockUserService) Create(user *models.User, password string) error {
	args := m.Called(user, password)
	user.PasswordHash = fmt.Sprintf("hash(%s)", password)
	now := time.Now()
	user.CreatedAt = now
	user.UpdatedAt = now
	user.PasswordUpdatedAt = now
	return args.Error(0)
}

func (m *MockUserService) UpdatePassword(user *models.User, password string) error {
	args := m.Called(password)
	return args.Error(0)
}

func ArgsUser(args mock.Arguments, index int) *models.User {
	var user *models.User
	var ok bool
	if user, ok = args.Get(0).(*models.User); !ok {
		panic(fmt.Sprintf("assert: arguments: *models.User(%d) failed because object wasn't correct type: %v", 0, args.Get(0)))
	}
	return user
}

func ArgsUserSlice(args mock.Arguments, index int) []*models.User {
	var users []*models.User
	var ok bool
	if users, ok = args.Get(0).([]*models.User); !ok {
		panic(fmt.Sprintf("assert: arguments: []*models.User(%d) failed because object wasn't correct type: %v", 0, args.Get(0)))
	}
	return users
}
