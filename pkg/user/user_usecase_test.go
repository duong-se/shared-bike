package user

import (
	"context"
	"testing"
	"time"

	"github.com/duong-se/shared-bike/apperrors"
	"github.com/duong-se/shared-bike/domain"
	"github.com/duong-se/shared-bike/pkg/user/mocks"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
)

type UserUseCaseTestSuite struct {
	suite.Suite
	mockRepository *mocks.IRepository
	mockLogger     *mocks.ILogger
	useCaseImpl    *useCaseImpl
}

func (s *UserUseCaseTestSuite) SetupTest() {
	mockRepository := &mocks.IRepository{}
	mockLogger := &mocks.ILogger{}
	s.mockRepository = mockRepository
	s.mockLogger = mockLogger
	s.mockLogger.On("Info", mock.Anything, mock.Anything, mock.Anything).Return()
	s.mockLogger.On("Error", mock.Anything, mock.Anything, mock.Anything).Return()
	s.mockLogger.On("Infof", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return()
	s.mockLogger.On("Errorf", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return()
	useCase := NewUseCase(mockLogger, mockRepository)
	s.useCaseImpl = useCase
}
func TestUserUseCaseTestSuite(t *testing.T) {
	suite.Run(t, new(UserUseCaseTestSuite))
}

func (s *UserUseCaseTestSuite) TestLogin_Success() {
	mockContext := context.TODO()
	mockPayload := domain.LoginPayload{
		Username: "testUsername",
		Password: "testPassword",
	}
	mockTime := time.Time{}
	mockUserResult := domain.User{
		ID:        1,
		Username:  "testUsername",
		Password:  "$2a$10$Mjx4fmq9ykGxlqlT/l9yGuojZ0FLV8QmrDhGwxmdE3QdkaXQgCcMG",
		Name:      "testName",
		CreatedAt: mockTime,
		UpdatedAt: mockTime,
		DeletedAt: gorm.DeletedAt{Valid: false},
	}
	s.mockRepository.On("GetByUsername", mockContext, mockPayload.Username).Return(&mockUserResult, nil)
	result, err := s.useCaseImpl.Login(context.TODO(), mockPayload)
	s.Nil(err)
	s.True(result)
}

func (s *UserUseCaseTestSuite) TestLogin_FailedByUserNotFound() {
	mockContext := context.TODO()
	mockPayload := domain.LoginPayload{
		Username: "testUsername",
		Password: "testPassword",
	}
	s.mockRepository.On("GetByUsername", mockContext, mockPayload.Username).Return(nil, gorm.ErrRecordNotFound)
	result, err := s.useCaseImpl.Login(context.TODO(), mockPayload)
	s.Equal(apperrors.ErrUserNotFound, err)
	s.False(result)
}

func (s *UserUseCaseTestSuite) TestLogin_FailedByPassword() {
	mockContext := context.TODO()
	mockPayload := domain.LoginPayload{
		Username: "testUsername",
		Password: "testPassword1",
	}
	mockTime := time.Time{}
	mockUserResult := domain.User{
		ID:        1,
		Username:  "testUsername",
		Password:  "$2a$10$Mjx4fmq9ykGxlqlT/l9yGuojZ0FLV8QmrDhGwxmdE3QdkaXQgCcMG",
		Name:      "testName",
		CreatedAt: mockTime,
		UpdatedAt: mockTime,
		DeletedAt: gorm.DeletedAt{Valid: false},
	}
	s.mockRepository.On("GetByUsername", mockContext, mockPayload.Username).Return(&mockUserResult, nil)
	result, err := s.useCaseImpl.Login(context.TODO(), mockPayload)
	s.Equal(apperrors.ErrUserNotFound, err)
	s.False(result)
}

func (s *UserUseCaseTestSuite) TestLogin_FailedByInternalError() {
	mockContext := context.TODO()
	mockPayload := domain.LoginPayload{
		Username: "testUsername",
		Password: "testPassword",
	}
	s.mockRepository.On("GetByUsername", mockContext, mockPayload.Username).Return(nil, gorm.ErrInvalidValue)
	result, err := s.useCaseImpl.Login(context.TODO(), mockPayload)
	s.Equal(apperrors.ErrInternalServerError, err)
	s.False(result)
}

func (s *UserUseCaseTestSuite) TestRegister_Success() {
	mockContext := context.TODO()
	mockPayload := domain.RegisterPayload{
		Username: "testUsername",
		Password: "testPassword",
		Name:     "testName",
	}
	s.mockRepository.On("Create", mockContext, mock.Anything).Return(nil)
	err := s.useCaseImpl.Register(context.TODO(), mockPayload)
	s.Nil(err)
}

func (s *UserUseCaseTestSuite) TestRegister_Failed() {
	mockContext := context.TODO()
	mockPayload := domain.RegisterPayload{
		Username: "testUsername",
		Password: "testPassword",
		Name:     "testName",
	}
	s.mockRepository.On("Create", mockContext, mock.Anything).Return(gorm.ErrInvalidDB)
	err := s.useCaseImpl.Register(context.TODO(), mockPayload)
	s.Equal(apperrors.ErrInternalServerError, err)
}
