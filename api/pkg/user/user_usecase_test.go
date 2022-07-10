package user

import (
	"context"
	"testing"
	"time"

	"shared-bike/apperrors"
	"shared-bike/domain"
	"shared-bike/pkg/user/mocks"

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
	s.mockLogger.On("Info", mock.Anything, mock.Anything).Return()
	s.mockLogger.On("Error", mock.Anything, mock.Anything).Return()
	useCase := NewUseCase(mockLogger, mockRepository)
	s.useCaseImpl = useCase
}
func TestUserUseCaseTestSuite(t *testing.T) {
	suite.Run(t, new(UserUseCaseTestSuite))
}

func (s *UserUseCaseTestSuite) TestLogin_Success() {
	mockContext := context.TODO()
	mockPayload := domain.LoginBody{
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
	actual, err := s.useCaseImpl.Login(context.TODO(), mockPayload)
	s.Nil(err)
	s.Equal(mockUserResult.ToDTO(), actual)
}

func (s *UserUseCaseTestSuite) TestLogin_FailedByUserNotFound() {
	mockContext := context.TODO()
	mockPayload := domain.LoginBody{
		Username: "testUsername",
		Password: "testPassword",
	}
	s.mockRepository.On("GetByUsername", mockContext, mockPayload.Username).Return(nil, gorm.ErrRecordNotFound)
	actual, err := s.useCaseImpl.Login(context.TODO(), mockPayload)
	s.Equal(apperrors.ErrUserLoginNotFound, err)
	s.Equal(domain.UserDTO{}, actual)
}

func (s *UserUseCaseTestSuite) TestLogin_FailedByPassword() {
	mockContext := context.TODO()
	mockPayload := domain.LoginBody{
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
	actual, err := s.useCaseImpl.Login(context.TODO(), mockPayload)
	s.Equal(apperrors.ErrUserLoginNotFound, err)
	s.Equal(domain.UserDTO{}, actual)
}

func (s *UserUseCaseTestSuite) TestLogin_FailedByInternalError() {
	mockContext := context.TODO()
	mockPayload := domain.LoginBody{
		Username: "testUsername",
		Password: "testPassword",
	}
	s.mockRepository.On("GetByUsername", mockContext, mockPayload.Username).Return(nil, gorm.ErrInvalidValue)
	actual, err := s.useCaseImpl.Login(context.TODO(), mockPayload)
	s.Equal(apperrors.ErrInternalServerError, err)
	s.Equal(domain.UserDTO{}, actual)
}

func (s *UserUseCaseTestSuite) TestRegister_Success() {
	mockContext := context.TODO()
	mockPayload := domain.RegisterBody{
		Username: "testUsername",
		Password: "testPassword",
		Name:     "testName",
	}
	mockUserResult := domain.UserDTO{
		ID:       0,
		Username: "testUsername",
		Name:     "testName",
	}
	s.mockRepository.On("Create", mockContext, mock.Anything).Return(nil)
	actual, err := s.useCaseImpl.Register(context.TODO(), mockPayload)
	s.Nil(err)
	s.Equal(mockUserResult, actual)
}

func (s *UserUseCaseTestSuite) TestRegister_Failed() {
	mockContext := context.TODO()
	mockPayload := domain.RegisterBody{
		Username: "testUsername",
		Password: "testPassword",
		Name:     "testName",
	}
	s.mockRepository.On("Create", mockContext, mock.Anything).Return(gorm.ErrInvalidDB)
	actual, err := s.useCaseImpl.Register(context.TODO(), mockPayload)
	s.Equal(apperrors.ErrInternalServerError, err)
	s.Equal(domain.UserDTO{}, actual)
}
