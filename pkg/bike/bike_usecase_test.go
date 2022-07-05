package bike

import (
	"context"
	"testing"

	"github.com/duong-se/shared-bike/apperrors"
	"github.com/duong-se/shared-bike/domain"
	"github.com/duong-se/shared-bike/pkg/bike/mocks"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
)

type BikeUseCaseTestSuite struct {
	suite.Suite
	mockRepository *mocks.IRepository
	mockLogger     *mocks.ILogger
	useCaseImpl    *useCaseImpl
}

func (s *BikeUseCaseTestSuite) SetupTest() {
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
func TestBikeUseCaseTestSuite(t *testing.T) {
	suite.Run(t, new(BikeUseCaseTestSuite))
}

func (s *BikeUseCaseTestSuite) TestGetAllBike_Success() {
	var (
		mockContext = context.TODO()
		mockResult  = []domain.Bike{
			{
				ID:     1,
				Lat:    "50.119504",
				Long:   "8.638137",
				Status: domain.BikeStatusAvailable,
			},
			{
				ID:     1,
				Lat:    "50.119229",
				Long:   "8.640020",
				Status: domain.BikeStatusRented,
			},
			{
				ID:     1,
				Lat:    "50.120452",
				Long:   "8.650507",
				Status: domain.BikeStatusAvailable,
			},
		}
	)
	s.mockRepository.On("GetList", mockContext).Return(&mockResult, nil)
	actual, err := s.useCaseImpl.GetAllBike(mockContext)
	s.Equal(mockResult, actual)
	s.Nil(err)
}

func (s *BikeUseCaseTestSuite) TestGetAllBike_RecordNotFound() {
	var (
		mockContext = context.TODO()
		mockResult  = []domain.Bike{}
	)
	s.mockRepository.On("GetList", mockContext).Return(&mockResult, gorm.ErrRecordNotFound)
	actual, err := s.useCaseImpl.GetAllBike(mockContext)
	s.Equal(mockResult, actual)
	s.Nil(err)
}

func (s *BikeUseCaseTestSuite) TestGetAllBike_InternalServerError() {
	var (
		mockContext = context.TODO()
		mockResult  = []domain.Bike{}
	)
	s.mockRepository.On("GetList", mockContext).Return(&mockResult, gorm.ErrInvalidData)
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
		mockExistRecord = domain.Bike{
			ID:     1,
			Lat:    "50.119504",
			Long:   "8.638137",
			Status: domain.BikeStatusAvailable,
			UserID: nil,
		}
		mockResult = domain.Bike{
			ID:     1,
			Lat:    "50.119504",
			Long:   "8.638137",
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
		mockExistRecord = domain.Bike{
			ID:     1,
			Lat:    "50.119504",
			Long:   "8.638137",
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
		mockExistRecord = domain.Bike{
			ID:     1,
			Lat:    "50.119504",
			Long:   "8.638137",
			Status: domain.BikeStatusAvailable,
			UserID: nil,
		}
		mockUpdateInput = domain.Bike{
			ID:     1,
			Lat:    "50.119504",
			Long:   "8.638137",
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
		mockExistRecord = domain.Bike{
			ID:     1,
			Lat:    "50.119504",
			Long:   "8.638137",
			Status: domain.BikeStatusRented,
			UserID: &userID,
		}
		mockResult = domain.Bike{
			ID:     1,
			Lat:    "50.119504",
			Long:   "8.638137",
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
		mockExistRecord = domain.Bike{
			ID:     1,
			Lat:    "50.119504",
			Long:   "8.638137",
			Status: domain.BikeStatusRented,
			UserID: &userID,
		}
		mockResult = domain.Bike{
			ID:     1,
			Lat:    "50.119504",
			Long:   "8.638137",
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
		mockExistRecord = domain.Bike{
			ID:     1,
			Lat:    "50.119504",
			Long:   "8.638137",
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
		mockExistRecord = domain.Bike{
			ID:     1,
			Lat:    "50.119504",
			Long:   "8.638137",
			Status: domain.BikeStatusRented,
			UserID: &userID,
		}
	)
	s.mockRepository.On("GetByID", mockContext, mockInput.ID).Return(&mockExistRecord, nil)
	actual, err := s.useCaseImpl.Return(mockContext, mockInput)
	s.Nil(actual)
	s.Equal(apperrors.ErrBikeNotYours, err)
}
