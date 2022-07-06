package bike

import (
	"context"
	"testing"
	"time"

	"shared-bike/apperrors"
	"shared-bike/domain"
	"shared-bike/pkg/bike/mocks"

	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
)

type BikeUseCaseTestSuite struct {
	suite.Suite
	mockRepository     *mocks.IRepository
	mockUserRepository *mocks.IUserRepository
	mockLogger         *mocks.ILogger
	useCaseImpl        *useCaseImpl
}

func (s *BikeUseCaseTestSuite) SetupTest() {
	mockRepository := &mocks.IRepository{}
	s.mockRepository = mockRepository
	mockLogger := &mocks.ILogger{}
	s.mockLogger = mockLogger
	mockUserRepository := &mocks.IUserRepository{}
	s.mockUserRepository = mockUserRepository
	s.mockLogger.On("Info", mock.Anything, mock.Anything, mock.Anything).Return()
	s.mockLogger.On("Error", mock.Anything, mock.Anything, mock.Anything).Return()
	s.mockLogger.On("Infof", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return()
	s.mockLogger.On("Errorf", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return()
	useCase := NewUseCase(mockLogger, mockRepository, mockUserRepository)
	s.useCaseImpl = useCase
}
func TestBikeUseCaseTestSuite(t *testing.T) {
	suite.Run(t, new(BikeUseCaseTestSuite))
}

func (s *BikeUseCaseTestSuite) TestGetAllBike_Success() {
	var (
		mockContext = context.TODO()
		mockUserID  = int64(1)
		lat         = decimal.NewFromFloat(50.119504)
		long        = decimal.NewFromFloat(8.638137)
		mockBike    = []domain.Bike{
			{
				ID:     1,
				Lat:    &lat,
				Long:   &long,
				Status: domain.BikeStatusAvailable,
				UserID: nil,
			},
			{
				ID:     1,
				Lat:    &lat,
				Long:   &long,
				Status: domain.BikeStatusRented,
				UserID: &mockUserID,
			},
			{
				ID:     1,
				Lat:    &lat,
				Long:   &long,
				Status: domain.BikeStatusAvailable,
				UserID: nil,
			},
		}
		mockTime       = time.Time{}
		mockUserResult = []domain.User{
			{
				ID:        1,
				Username:  "testUsername",
				Password:  "testPassword",
				Name:      "testName",
				CreatedAt: mockTime,
				UpdatedAt: mockTime,
				DeletedAt: gorm.DeletedAt{Valid: false},
			},
		}
		mockResult = []domain.GetAllBikeResponse{
			{
				ID:               1,
				Lat:              "50.119504",
				Long:             "8.638137",
				Status:           domain.BikeStatusAvailable,
				UserID:           nil,
				NameOfRenter:     nil,
				UsernameOfRenter: nil,
			},
			{
				ID:               1,
				Lat:              "50.119504",
				Long:             "8.638137",
				Status:           domain.BikeStatusRented,
				UserID:           &mockUserID,
				NameOfRenter:     &mockUserResult[0].Name,
				UsernameOfRenter: &mockUserResult[0].Username,
			},
			{
				ID:               1,
				Lat:              "50.119504",
				Long:             "8.638137",
				Status:           domain.BikeStatusAvailable,
				UserID:           nil,
				NameOfRenter:     nil,
				UsernameOfRenter: nil,
			},
		}
	)
	s.mockRepository.On("GetList", mockContext).Return(&mockBike, nil)
	s.mockUserRepository.On("GetListByIDs", mockContext, []int64{1}).Return(&mockUserResult, nil)
	actual, err := s.useCaseImpl.GetAllBike(mockContext)
	s.Equal(mockResult, actual)
	s.Nil(err)
}

func (s *BikeUseCaseTestSuite) TestGetAllBike_RecordNotFound() {
	var (
		mockContext = context.TODO()
		mockResult  = []domain.GetAllBikeResponse{}
	)
	s.mockRepository.On("GetList", mockContext).Return(nil, gorm.ErrRecordNotFound)
	actual, err := s.useCaseImpl.GetAllBike(mockContext)
	s.Equal(mockResult, actual)
	s.Nil(err)
}

func (s *BikeUseCaseTestSuite) TestGetAllBike_InternalServerError() {
	var (
		mockContext = context.TODO()
		mockResult  = []domain.GetAllBikeResponse{}
	)
	s.mockRepository.On("GetList", mockContext).Return(nil, gorm.ErrInvalidData)
	actual, err := s.useCaseImpl.GetAllBike(mockContext)
	s.Equal(mockResult, actual)
	s.Equal(apperrors.ErrInternalServerError, err)
}

func (s *BikeUseCaseTestSuite) TestRent_Success() {
	var (
		mockContext = context.TODO()
		userID      = int64(1)
		mockInput   = domain.RentOrReturnRequestPayload{
			ID:     1,
			UserID: 1,
		}
		lat             = decimal.NewFromFloat(50.119504)
		long            = decimal.NewFromFloat(8.638137)
		mockExistRecord = domain.Bike{
			ID:     1,
			Lat:    &lat,
			Long:   &long,
			Status: domain.BikeStatusAvailable,
			UserID: nil,
		}
		mockResult = domain.Bike{
			ID:     1,
			Lat:    &lat,
			Long:   &long,
			Status: domain.BikeStatusRented,
			UserID: &userID,
		}
	)
	s.mockRepository.On("GetByID", mockContext, mockInput.ID).Return(&mockExistRecord, nil)
	s.mockRepository.On("Update", mockContext, &mockResult).Return(nil)
	actual, err := s.useCaseImpl.Rent(mockContext, mockInput)
	s.Equal(mockResult, *actual)
	s.Nil(err)
}

func (s *BikeUseCaseTestSuite) TestRent_NotFoundData() {
	var (
		mockContext = context.TODO()
		mockInput   = domain.RentOrReturnRequestPayload{
			ID:     1,
			UserID: 1,
		}
	)
	s.mockRepository.On("GetByID", mockContext, mockInput.ID).Return(nil, gorm.ErrRecordNotFound)
	actual, err := s.useCaseImpl.Rent(mockContext, mockInput)
	s.Nil(actual)
	s.Equal(apperrors.ErrBikeNotFound, err)
}

func (s *BikeUseCaseTestSuite) TestRent_NotAvailableByUserIDAndStatus() {
	var (
		userID      = int64(2)
		mockContext = context.TODO()
		mockInput   = domain.RentOrReturnRequestPayload{
			ID:     1,
			UserID: 1,
		}
		lat             = decimal.NewFromFloat(50.119504)
		long            = decimal.NewFromFloat(8.638137)
		mockExistRecord = domain.Bike{
			ID:     1,
			Lat:    &lat,
			Long:   &long,
			Status: domain.BikeStatusRented,
			UserID: &userID,
		}
	)
	s.mockRepository.On("GetByID", mockContext, mockInput.ID).Return(&mockExistRecord, nil)
	actual, err := s.useCaseImpl.Rent(mockContext, mockInput)
	s.Nil(actual)
	s.Equal(apperrors.ErrBikeRented, err)
}

func (s *BikeUseCaseTestSuite) TestRent_InternalServerErrorWhenFetchCurrentBike() {
	var (
		mockContext = context.TODO()
		mockInput   = domain.RentOrReturnRequestPayload{
			ID:     1,
			UserID: 1,
		}
	)
	s.mockRepository.On("GetByID", mockContext, mockInput.ID).Return(nil, gorm.ErrInvalidData)
	actual, err := s.useCaseImpl.Rent(mockContext, mockInput)
	s.Nil(actual)
	s.Equal(apperrors.ErrInternalServerError, err)
}

func (s *BikeUseCaseTestSuite) TestRent_InternalServerErrorWhenUpdate() {
	var (
		userID      = int64(1)
		mockContext = context.TODO()
		mockInput   = domain.RentOrReturnRequestPayload{
			ID:     1,
			UserID: 1,
		}
		lat             = decimal.NewFromFloat(50.119504)
		long            = decimal.NewFromFloat(8.638137)
		mockExistRecord = domain.Bike{
			ID:     1,
			Lat:    &lat,
			Long:   &long,
			Status: domain.BikeStatusAvailable,
			UserID: nil,
		}
		mockUpdateInput = domain.Bike{
			ID:     1,
			Lat:    &lat,
			Long:   &long,
			Status: domain.BikeStatusRented,
			UserID: &userID,
		}
	)
	s.mockRepository.On("GetByID", mockContext, mockInput.ID).Return(&mockExistRecord, nil)
	s.mockRepository.On("Update", mockContext, &mockUpdateInput).Return(gorm.ErrInvalidData)
	actual, err := s.useCaseImpl.Rent(mockContext, mockInput)
	s.Nil(actual)
	s.Equal(apperrors.ErrInternalServerError, err)
}

func (s *BikeUseCaseTestSuite) TestReturn_Success() {
	var (
		userID      = int64(1)
		mockContext = context.TODO()
		mockInput   = domain.RentOrReturnRequestPayload{
			ID:     1,
			UserID: 1,
		}
		lat             = decimal.NewFromFloat(50.119504)
		long            = decimal.NewFromFloat(8.638137)
		mockExistRecord = domain.Bike{
			ID:     1,
			Lat:    &lat,
			Long:   &long,
			Status: domain.BikeStatusRented,
			UserID: &userID,
		}
		mockResult = domain.Bike{
			ID:     1,
			Lat:    &lat,
			Long:   &long,
			Status: domain.BikeStatusAvailable,
			UserID: nil,
		}
	)
	s.mockRepository.On("GetByID", mockContext, mockInput.ID).Return(&mockExistRecord, nil)
	s.mockRepository.On("Update", mockContext, &mockResult).Return(nil)
	actual, err := s.useCaseImpl.Return(mockContext, mockInput)
	s.Equal(mockResult, *actual)
	s.Nil(err)
}

func (s *BikeUseCaseTestSuite) TestReturn_InternalServerErrorWhenUpdate() {
	var (
		userID      = int64(1)
		mockContext = context.TODO()
		mockInput   = domain.RentOrReturnRequestPayload{
			ID:     1,
			UserID: 1,
		}
		lat             = decimal.NewFromFloat(50.119504)
		long            = decimal.NewFromFloat(8.638137)
		mockExistRecord = domain.Bike{
			ID:     1,
			Lat:    &lat,
			Long:   &long,
			Status: domain.BikeStatusRented,
			UserID: &userID,
		}
		mockResult = domain.Bike{
			ID:     1,
			Lat:    &lat,
			Long:   &long,
			Status: domain.BikeStatusAvailable,
			UserID: nil,
		}
	)
	s.mockRepository.On("GetByID", mockContext, mockInput.ID).Return(&mockExistRecord, nil)
	s.mockRepository.On("Update", mockContext, &mockResult).Return(gorm.ErrEmptySlice)
	actual, err := s.useCaseImpl.Return(mockContext, mockInput)
	s.Nil(actual)
	s.Equal(apperrors.ErrInternalServerError, err)
}

func (s *BikeUseCaseTestSuite) TestReturn_InternalServerErrorWhenGetByID() {
	var (
		mockContext = context.TODO()
		mockInput   = domain.RentOrReturnRequestPayload{
			ID:     1,
			UserID: 1,
		}
	)
	s.mockRepository.On("GetByID", mockContext, mockInput.ID).Return(nil, gorm.ErrInvalidDB)
	actual, err := s.useCaseImpl.Return(mockContext, mockInput)
	s.Nil(actual)
	s.Equal(apperrors.ErrInternalServerError, err)
}

func (s *BikeUseCaseTestSuite) TestReturn_NotFoundBike() {
	var (
		mockContext = context.TODO()
		mockInput   = domain.RentOrReturnRequestPayload{
			ID:     1,
			UserID: 1,
		}
	)
	s.mockRepository.On("GetByID", mockContext, mockInput.ID).Return(nil, gorm.ErrRecordNotFound)
	actual, err := s.useCaseImpl.Return(mockContext, mockInput)
	s.Nil(actual)
	s.Equal(apperrors.ErrBikeNotFound, err)
}

func (s *BikeUseCaseTestSuite) TestReturn_BikeNotAvailable() {
	var (
		mockContext = context.TODO()
		mockInput   = domain.RentOrReturnRequestPayload{
			ID:     1,
			UserID: 1,
		}
		lat             = decimal.NewFromFloat(50.119504)
		long            = decimal.NewFromFloat(8.638137)
		mockExistRecord = domain.Bike{
			ID:     1,
			Lat:    &lat,
			Long:   &long,
			Status: domain.BikeStatusAvailable,
			UserID: nil,
		}
	)
	s.mockRepository.On("GetByID", mockContext, mockInput.ID).Return(&mockExistRecord, nil)
	actual, err := s.useCaseImpl.Return(mockContext, mockInput)
	s.Nil(actual)
	s.Equal(apperrors.ErrBikeAvailable, err)
}

func (s *BikeUseCaseTestSuite) TestReturn_BikeIsNotYours() {
	var (
		userID      = int64(1)
		mockContext = context.TODO()
		mockInput   = domain.RentOrReturnRequestPayload{
			ID:     1,
			UserID: 2,
		}
		lat             = decimal.NewFromFloat(50.119504)
		long            = decimal.NewFromFloat(8.638137)
		mockExistRecord = domain.Bike{
			ID:     1,
			Lat:    &lat,
			Long:   &long,
			Status: domain.BikeStatusRented,
			UserID: &userID,
		}
	)
	s.mockRepository.On("GetByID", mockContext, mockInput.ID).Return(&mockExistRecord, nil)
	actual, err := s.useCaseImpl.Return(mockContext, mockInput)
	s.Nil(actual)
	s.Equal(apperrors.ErrBikeNotYours, err)
}
