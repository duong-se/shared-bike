package bike

import (
	"context"
	"database/sql"
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
	s.mockLogger.On("Info", mock.Anything, mock.Anything).Return()
	s.mockLogger.On("Error", mock.Anything, mock.Anything).Return()
	useCase := NewUseCase(mockLogger, mockRepository, mockUserRepository)
	s.useCaseImpl = useCase
}
func TestBikeUseCaseTestSuite(t *testing.T) {
	suite.Run(t, new(BikeUseCaseTestSuite))
}

func (s *BikeUseCaseTestSuite) TestGetAllBike_Success() {
	var (
		mockContext = context.TODO()
		mockUserID  = sql.NullInt64{
			Valid: true,
			Int64: 1,
		}
		mockNilUserID = sql.NullInt64{
			Valid: false,
			Int64: 0,
		}
		lat      = decimal.NewFromFloat(50.119504)
		long     = decimal.NewFromFloat(8.638137)
		mockBike = []domain.Bike{
			{
				ID:     1,
				Lat:    &lat,
				Long:   &long,
				Name:   "testName",
				Status: domain.BikeStatusAvailable,
				UserID: mockNilUserID,
			},
			{
				ID:     2,
				Lat:    &lat,
				Long:   &long,
				Name:   "testName",
				Status: domain.BikeStatusRented,
				UserID: mockUserID,
			},
			{
				ID:     3,
				Lat:    &lat,
				Long:   &long,
				Name:   "testName",
				Status: domain.BikeStatusAvailable,
				UserID: mockNilUserID,
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
		mockResult = []domain.BikeDTO{
			{
				ID:           1,
				Lat:          "50.119504",
				Long:         "8.638137",
				Name:         "testName",
				Status:       domain.BikeStatusAvailable,
				UserID:       0,
				NameOfRenter: "",
			},
			{
				ID:           2,
				Lat:          "50.119504",
				Long:         "8.638137",
				Name:         "testName",
				Status:       domain.BikeStatusRented,
				UserID:       mockUserID.Int64,
				NameOfRenter: mockUserResult[0].Name,
			},
			{
				ID:           3,
				Lat:          "50.119504",
				Long:         "8.638137",
				Name:         "testName",
				Status:       domain.BikeStatusAvailable,
				UserID:       0,
				NameOfRenter: "",
			},
		}
	)
	s.mockRepository.On("GetList", mockContext).Return(&mockBike, nil)
	s.mockUserRepository.On("GetListByIDs", mockContext, []int64{1}).Return(&mockUserResult, nil)
	actual, err := s.useCaseImpl.GetAllBike(mockContext)
	s.Equal(mockResult, actual)
	s.Nil(err)
}

func (s *BikeUseCaseTestSuite) TestGetAllBike_FetchUserListEmpty() {
	var (
		mockContext = context.TODO()
		lat         = decimal.NewFromFloat(50.119504)
		long        = decimal.NewFromFloat(8.638137)
		mockUserID  = sql.NullInt64{
			Valid: true,
			Int64: 1,
		}
		mockNilUserID = sql.NullInt64{
			Valid: false,
			Int64: 0,
		}
		mockBike = []domain.Bike{
			{
				ID:     1,
				Lat:    &lat,
				Long:   &long,
				Name:   "testName",
				Status: domain.BikeStatusAvailable,
				UserID: mockNilUserID,
			},
			{
				ID:     2,
				Lat:    &lat,
				Long:   &long,
				Name:   "testName",
				Status: domain.BikeStatusRented,
				UserID: mockUserID,
			},
			{
				ID:     3,
				Lat:    &lat,
				Long:   &long,
				Name:   "testName",
				Status: domain.BikeStatusAvailable,
				UserID: mockNilUserID,
			},
		}
		mockResult = []domain.BikeDTO{
			{
				ID:           1,
				Lat:          "50.119504",
				Long:         "8.638137",
				Name:         "testName",
				Status:       domain.BikeStatusAvailable,
				UserID:       0,
				NameOfRenter: "",
			},
			{
				ID:           2,
				Lat:          "50.119504",
				Long:         "8.638137",
				Name:         "testName",
				Status:       domain.BikeStatusRented,
				UserID:       mockUserID.Int64,
				NameOfRenter: "",
			},
			{
				ID:           3,
				Lat:          "50.119504",
				Long:         "8.638137",
				Status:       domain.BikeStatusAvailable,
				UserID:       0,
				Name:         "testName",
				NameOfRenter: "",
			},
		}
	)
	s.mockRepository.On("GetList", mockContext).Return(&mockBike, nil)
	s.mockUserRepository.On("GetListByIDs", mockContext, []int64{1}).Return(nil, gorm.ErrRecordNotFound)
	actual, err := s.useCaseImpl.GetAllBike(mockContext)
	s.Equal(mockResult, actual)
	s.Nil(err)
}

func (s *BikeUseCaseTestSuite) TestGetAllBike_CannotFetchUserList() {
	var (
		mockContext = context.TODO()
		lat         = decimal.NewFromFloat(50.119504)
		long        = decimal.NewFromFloat(8.638137)
		mockUserID  = sql.NullInt64{
			Valid: true,
			Int64: 1,
		}
		mockNilUserID = sql.NullInt64{
			Valid: false,
			Int64: 0,
		}
		mockBike = []domain.Bike{
			{
				ID:     1,
				Lat:    &lat,
				Long:   &long,
				Name:   "testName",
				Status: domain.BikeStatusAvailable,
				UserID: mockNilUserID,
			},
			{
				ID:     2,
				Lat:    &lat,
				Long:   &long,
				Name:   "testName",
				Status: domain.BikeStatusRented,
				UserID: mockUserID,
			},
			{
				ID:     3,
				Lat:    &lat,
				Long:   &long,
				Name:   "testName",
				Status: domain.BikeStatusAvailable,
				UserID: mockNilUserID,
			},
		}
	)
	s.mockRepository.On("GetList", mockContext).Return(&mockBike, nil)
	s.mockUserRepository.On("GetListByIDs", mockContext, []int64{1}).Return(nil, gorm.ErrDryRunModeUnsupported)
	actual, err := s.useCaseImpl.GetAllBike(mockContext)
	s.Equal([]domain.BikeDTO{}, actual)
	s.Error(apperrors.ErrInternalServerError, err)
}

func (s *BikeUseCaseTestSuite) TestGetAllBike_RecordNotFound() {
	var (
		mockContext = context.TODO()
		mockResult  = []domain.BikeDTO{}
	)
	s.mockRepository.On("GetList", mockContext).Return(nil, gorm.ErrRecordNotFound)
	actual, err := s.useCaseImpl.GetAllBike(mockContext)
	s.Equal(mockResult, actual)
	s.Nil(err)
}

func (s *BikeUseCaseTestSuite) TestGetAllBike_InternalServerError() {
	var (
		mockContext = context.TODO()
		mockResult  = []domain.BikeDTO{}
	)
	s.mockRepository.On("GetList", mockContext).Return(nil, gorm.ErrInvalidData)
	actual, err := s.useCaseImpl.GetAllBike(mockContext)
	s.Equal(mockResult, actual)
	s.Equal(apperrors.ErrInternalServerError, err)
}

func (s *BikeUseCaseTestSuite) TestRent_Success() {
	var (
		mockContext = context.TODO()
		mockUserID  = sql.NullInt64{
			Valid: true,
			Int64: 1,
		}
		mockNilUserID = sql.NullInt64{
			Valid: false,
			Int64: 0,
		}
		mockInput = domain.RentOrReturnRequestPayload{
			ID:     1,
			UserID: 1,
		}
		lat             = decimal.NewFromFloat(50.119504)
		long            = decimal.NewFromFloat(8.638137)
		mockExistRecord = domain.Bike{
			ID:     1,
			Lat:    &lat,
			Long:   &long,
			Name:   "testName",
			Status: domain.BikeStatusAvailable,
			UserID: mockNilUserID,
		}
		mockTime       = time.Time{}
		mockUserResult = domain.User{
			ID:        1,
			Username:  "testUsername",
			Password:  "$2a$10$Mjx4fmq9ykGxlqlT/l9yGuojZ0FLV8QmrDhGwxmdE3QdkaXQgCcMG",
			Name:      "testName",
			CreatedAt: mockTime,
			UpdatedAt: mockTime,
			DeletedAt: gorm.DeletedAt{Valid: false},
		}
		mockUpdateInput = domain.Bike{
			ID:     1,
			Lat:    &lat,
			Long:   &long,
			Name:   "testName",
			Status: domain.BikeStatusRented,
			UserID: mockUserID,
		}
		expected = domain.BikeDTO{
			ID:           1,
			Lat:          lat.String(),
			Long:         long.String(),
			Name:         "testName",
			Status:       domain.BikeStatusRented,
			UserID:       mockUserResult.ID,
			NameOfRenter: mockUserResult.Name,
		}
	)
	s.mockRepository.On("CountByUserID", mockContext, mockInput.UserID).Return(int64(0), nil)
	s.mockUserRepository.On("GetByID", mockContext, mockInput.ID).Return(&mockUserResult, nil)
	s.mockRepository.On("GetByID", mockContext, mockInput.ID).Return(&mockExistRecord, nil)
	s.mockRepository.On("UpdateStatusAndUserID", mockContext, &mockUpdateInput).Return(nil)
	actual, err := s.useCaseImpl.Rent(mockContext, mockInput)
	s.Equal(expected, actual)
	s.Nil(err)
}

func (s *BikeUseCaseTestSuite) TestRent_FailedByAlreadyRented() {
	var (
		mockContext = context.TODO()
		mockInput   = domain.RentOrReturnRequestPayload{
			ID:     1,
			UserID: 1,
		}
	)
	s.mockRepository.On("CountByUserID", mockContext, mockInput.UserID).Return(int64(1), nil)
	actual, err := s.useCaseImpl.Rent(mockContext, mockInput)
	s.Equal(domain.BikeDTO{}, actual)
	s.Equal(apperrors.ErrUserHasBikeAlready, err)
}

func (s *BikeUseCaseTestSuite) TestRent_FailedByCountByUserID() {
	var (
		mockContext = context.TODO()
		mockInput   = domain.RentOrReturnRequestPayload{
			ID:     1,
			UserID: 1,
		}
	)
	s.mockRepository.On("CountByUserID", mockContext, mockInput.UserID).Return(int64(0), gorm.ErrInvalidValue)
	actual, err := s.useCaseImpl.Rent(mockContext, mockInput)
	s.Equal(domain.BikeDTO{}, actual)
	s.Equal(apperrors.ErrInternalServerError, err)
}

func (s *BikeUseCaseTestSuite) TestRent_NotFoundCurrentUser() {
	var (
		mockContext = context.TODO()
		mockInput   = domain.RentOrReturnRequestPayload{
			ID:     1,
			UserID: 1,
		}
	)
	s.mockRepository.On("CountByUserID", mockContext, mockInput.UserID).Return(int64(0), nil)
	s.mockUserRepository.On("GetByID", mockContext, mockInput.ID).Return(nil, gorm.ErrRecordNotFound)
	actual, err := s.useCaseImpl.Rent(mockContext, mockInput)
	s.Equal(domain.BikeDTO{}, actual)
	s.Equal(apperrors.ErrUserNotExisted, err)
}

func (s *BikeUseCaseTestSuite) TestRent_ErrorOnFetchCurrentUser() {
	var (
		mockContext = context.TODO()
		mockInput   = domain.RentOrReturnRequestPayload{
			ID:     1,
			UserID: 1,
		}
	)
	s.mockRepository.On("CountByUserID", mockContext, mockInput.UserID).Return(int64(0), nil)
	s.mockUserRepository.On("GetByID", mockContext, mockInput.ID).Return(nil, gorm.ErrInvalidField)
	actual, err := s.useCaseImpl.Rent(mockContext, mockInput)
	s.Equal(domain.BikeDTO{}, actual)
	s.Equal(apperrors.ErrInternalServerError, err)
}

func (s *BikeUseCaseTestSuite) TestRent_NotFoundData() {
	var (
		mockContext = context.TODO()
		mockInput   = domain.RentOrReturnRequestPayload{
			ID:     1,
			UserID: 1,
		}
		mockTime       = time.Time{}
		mockUserResult = domain.User{
			ID:        1,
			Username:  "testUsername",
			Password:  "$2a$10$Mjx4fmq9ykGxlqlT/l9yGuojZ0FLV8QmrDhGwxmdE3QdkaXQgCcMG",
			Name:      "testName",
			CreatedAt: mockTime,
			UpdatedAt: mockTime,
			DeletedAt: gorm.DeletedAt{Valid: false},
		}
	)
	s.mockRepository.On("CountByUserID", mockContext, mockInput.UserID).Return(int64(0), nil)
	s.mockUserRepository.On("GetByID", mockContext, mockInput.ID).Return(&mockUserResult, nil)
	s.mockRepository.On("GetByID", mockContext, mockInput.ID).Return(nil, gorm.ErrRecordNotFound)
	actual, err := s.useCaseImpl.Rent(mockContext, mockInput)
	s.Equal(domain.BikeDTO{}, actual)
	s.Equal(apperrors.ErrBikeNotFound, err)
}

func (s *BikeUseCaseTestSuite) TestRent_NotAvailableByUserIDAndStatus() {
	var (
		mockUserID = sql.NullInt64{
			Valid: true,
			Int64: 1,
		}
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
			Name:   "testName",
			Status: domain.BikeStatusRented,
			UserID: mockUserID,
		}
		mockTime       = time.Time{}
		mockUserResult = domain.User{
			ID:        1,
			Username:  "testUsername",
			Password:  "$2a$10$Mjx4fmq9ykGxlqlT/l9yGuojZ0FLV8QmrDhGwxmdE3QdkaXQgCcMG",
			Name:      "testName",
			CreatedAt: mockTime,
			UpdatedAt: mockTime,
			DeletedAt: gorm.DeletedAt{Valid: false},
		}
	)
	s.mockRepository.On("CountByUserID", mockContext, mockInput.UserID).Return(int64(0), nil)
	s.mockUserRepository.On("GetByID", mockContext, mockInput.ID).Return(&mockUserResult, nil)
	s.mockRepository.On("GetByID", mockContext, mockInput.ID).Return(&mockExistRecord, nil)
	actual, err := s.useCaseImpl.Rent(mockContext, mockInput)
	s.Equal(domain.BikeDTO{}, actual)
	s.Equal(apperrors.ErrBikeRented, err)
}

func (s *BikeUseCaseTestSuite) TestRent_InternalServerErrorWhenFetchCurrentBike() {
	var (
		mockContext = context.TODO()
		mockInput   = domain.RentOrReturnRequestPayload{
			ID:     1,
			UserID: 1,
		}
		mockTime       = time.Time{}
		mockUserResult = domain.User{
			ID:        1,
			Username:  "testUsername",
			Password:  "$2a$10$Mjx4fmq9ykGxlqlT/l9yGuojZ0FLV8QmrDhGwxmdE3QdkaXQgCcMG",
			Name:      "testName",
			CreatedAt: mockTime,
			UpdatedAt: mockTime,
			DeletedAt: gorm.DeletedAt{Valid: false},
		}
	)
	s.mockRepository.On("CountByUserID", mockContext, mockInput.UserID).Return(int64(0), nil)
	s.mockUserRepository.On("GetByID", mockContext, mockInput.ID).Return(&mockUserResult, nil)
	s.mockRepository.On("GetByID", mockContext, mockInput.ID).Return(nil, gorm.ErrInvalidData)
	actual, err := s.useCaseImpl.Rent(mockContext, mockInput)
	s.Equal(domain.BikeDTO{}, actual)
	s.Equal(apperrors.ErrInternalServerError, err)
}

func (s *BikeUseCaseTestSuite) TestRent_InternalServerErrorWhenUpdate() {
	var (
		mockContext = context.TODO()
		mockInput   = domain.RentOrReturnRequestPayload{
			ID:     1,
			UserID: 1,
		}
		mockUserID = sql.NullInt64{
			Valid: true,
			Int64: 1,
		}
		mockNilUserID = sql.NullInt64{
			Valid: false,
			Int64: 0,
		}
		lat             = decimal.NewFromFloat(50.119504)
		long            = decimal.NewFromFloat(8.638137)
		mockExistRecord = domain.Bike{
			ID:     1,
			Lat:    &lat,
			Long:   &long,
			Name:   "testName",
			Status: domain.BikeStatusAvailable,
			UserID: mockNilUserID,
		}
		mockUpdateInput = domain.Bike{
			ID:     1,
			Lat:    &lat,
			Long:   &long,
			Name:   "testName",
			Status: domain.BikeStatusRented,
			UserID: mockUserID,
		}
		mockTime       = time.Time{}
		mockUserResult = domain.User{
			ID:        1,
			Username:  "testUsername",
			Password:  "$2a$10$Mjx4fmq9ykGxlqlT/l9yGuojZ0FLV8QmrDhGwxmdE3QdkaXQgCcMG",
			Name:      "testName",
			CreatedAt: mockTime,
			UpdatedAt: mockTime,
			DeletedAt: gorm.DeletedAt{Valid: false},
		}
	)
	s.mockRepository.On("CountByUserID", mockContext, mockInput.UserID).Return(int64(0), nil)
	s.mockUserRepository.On("GetByID", mockContext, mockInput.ID).Return(&mockUserResult, nil)
	s.mockRepository.On("GetByID", mockContext, mockInput.ID).Return(&mockExistRecord, nil)
	s.mockRepository.On("UpdateStatusAndUserID", mockContext, &mockUpdateInput).Return(gorm.ErrInvalidData)
	actual, err := s.useCaseImpl.Rent(mockContext, mockInput)
	s.Equal(domain.BikeDTO{}, actual)
	s.Equal(apperrors.ErrInternalServerError, err)
}

func (s *BikeUseCaseTestSuite) TestReturn_Success() {
	var (
		mockUserID = sql.NullInt64{
			Valid: true,
			Int64: 1,
		}
		mockNilUserID = sql.NullInt64{
			Valid: false,
			Int64: 0,
		}
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
			Name:   "testName",
			Status: domain.BikeStatusRented,
			UserID: mockUserID,
		}
		mockResult = domain.Bike{
			ID:     1,
			Lat:    &lat,
			Long:   &long,
			Name:   "testName",
			Status: domain.BikeStatusAvailable,
			UserID: mockNilUserID,
		}
	)
	s.mockRepository.On("GetByID", mockContext, mockInput.ID).Return(&mockExistRecord, nil)
	s.mockRepository.On("UpdateStatusAndUserID", mockContext, &mockResult).Return(nil)
	actual, err := s.useCaseImpl.Return(mockContext, mockInput)
	s.Equal(mockResult.ToDTO(), actual)
	s.Nil(err)
}

func (s *BikeUseCaseTestSuite) TestReturn_InternalServerErrorWhenUpdate() {
	var (
		mockUserID = sql.NullInt64{
			Valid: true,
			Int64: 1,
		}
		mockNilUserID = sql.NullInt64{
			Valid: false,
			Int64: 0,
		}
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
			Name:   "testName",
			Status: domain.BikeStatusRented,
			UserID: mockUserID,
		}
		mockResult = domain.Bike{
			ID:     1,
			Lat:    &lat,
			Long:   &long,
			Name:   "testName",
			Status: domain.BikeStatusAvailable,
			UserID: mockNilUserID,
		}
	)
	s.mockRepository.On("GetByID", mockContext, mockInput.ID).Return(&mockExistRecord, nil)
	s.mockRepository.On("UpdateStatusAndUserID", mockContext, &mockResult).Return(gorm.ErrEmptySlice)
	actual, err := s.useCaseImpl.Return(mockContext, mockInput)
	s.Equal(domain.BikeDTO{}, actual)
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
	s.Equal(domain.BikeDTO{}, actual)
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
	s.Equal(domain.BikeDTO{}, actual)
	s.Equal(apperrors.ErrBikeNotFound, err)
}

func (s *BikeUseCaseTestSuite) TestReturn_BikeNotAvailable() {
	var (
		mockContext   = context.TODO()
		mockNilUserID = sql.NullInt64{
			Valid: false,
			Int64: 0,
		}
		mockInput = domain.RentOrReturnRequestPayload{
			ID:     1,
			UserID: 1,
		}
		lat             = decimal.NewFromFloat(50.119504)
		long            = decimal.NewFromFloat(8.638137)
		mockExistRecord = domain.Bike{
			ID:     1,
			Lat:    &lat,
			Long:   &long,
			Name:   "testName",
			Status: domain.BikeStatusAvailable,
			UserID: mockNilUserID,
		}
	)
	s.mockRepository.On("GetByID", mockContext, mockInput.ID).Return(&mockExistRecord, nil)
	actual, err := s.useCaseImpl.Return(mockContext, mockInput)
	s.Equal(domain.BikeDTO{}, actual)
	s.Equal(apperrors.ErrBikeAvailable, err)
}

func (s *BikeUseCaseTestSuite) TestReturn_BikeIsNotYours() {
	var (
		mockUserID = sql.NullInt64{
			Valid: true,
			Int64: 1,
		}
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
			Name:   "testName",
			Status: domain.BikeStatusRented,
			UserID: mockUserID,
		}
	)
	s.mockRepository.On("GetByID", mockContext, mockInput.ID).Return(&mockExistRecord, nil)
	actual, err := s.useCaseImpl.Return(mockContext, mockInput)
	s.Equal(domain.BikeDTO{}, actual)
	s.Equal(apperrors.ErrBikeNotYours, err)
}
