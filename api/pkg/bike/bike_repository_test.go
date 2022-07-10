package bike

import (
	"context"
	"database/sql"
	"regexp"
	"testing"
	"time"

	"shared-bike/domain"

	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/suite"
	sqlmock "gopkg.in/DATA-DOG/go-sqlmock.v1"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type BikeRepositoryTestSuite struct {
	suite.Suite
	mockDB         sqlmock.Sqlmock
	repositoryImpl *repositoryImpl
}

func (s *BikeRepositoryTestSuite) SetupTest() {
	db, mock, err := sqlmock.New()
	if err != nil {
		s.Error(err, "Failed to open mock sql db, got error")
	}

	if db == nil {
		s.Error(nil, "mock db is null")
	}

	if mock == nil {
		s.Error(nil, "sqlmock is null")
	}

	gormDB, err := gorm.Open(mysql.New(mysql.Config{
		Conn:                      db,
		SkipInitializeWithVersion: true,
	}), &gorm.Config{})
	if err != nil {
		s.Error(err, "gormDB init err")
	}
	s.mockDB = mock
	repositoryImpl := NewRepository(gormDB)
	s.repositoryImpl = repositoryImpl
}
func TestBikeRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(BikeRepositoryTestSuite))
}

func (s *BikeRepositoryTestSuite) TestGetList_Success() {
	mockUserID := sql.NullInt64{
		Valid: true,
		Int64: 1,
	}
	mockNilUserID := sql.NullInt64{
		Valid: false,
		Int64: 0,
	}
	mockTime := time.Time{}
	lat := decimal.NewFromFloat(50.119504)
	long := decimal.NewFromFloat(8.638137)
	mockBikes := []domain.Bike{
		{
			ID:        1,
			Lat:       &lat,
			Long:      &long,
			Status:    domain.BikeStatusAvailable,
			UserID:    mockNilUserID,
			CreatedAt: mockTime,
			UpdatedAt: mockTime,
			DeletedAt: gorm.DeletedAt{
				Valid: false,
				Time:  time.Time{},
			},
		},
		{
			ID:        1,
			Lat:       &lat,
			Long:      &long,
			Status:    domain.BikeStatusRented,
			UserID:    mockUserID,
			CreatedAt: mockTime,
			UpdatedAt: mockTime,
			DeletedAt: gorm.DeletedAt{
				Valid: false,
				Time:  time.Time{},
			},
		},
		{
			ID:        1,
			Lat:       &lat,
			Long:      &long,
			Status:    domain.BikeStatusAvailable,
			UserID:    mockNilUserID,
			CreatedAt: mockTime,
			UpdatedAt: mockTime,
			DeletedAt: gorm.DeletedAt{
				Valid: false,
				Time:  time.Time{},
			},
		},
	}

	rows := sqlmock.NewRows([]string{"id", "lat", "long", "status", "user_id", "created_at", "updated_at", "deleted_at"}).
		AddRow(mockBikes[0].ID, mockBikes[0].Lat, mockBikes[0].Long,
			mockBikes[0].Status, nil, mockBikes[0].CreatedAt, mockBikes[0].UpdatedAt, nil).
		AddRow(mockBikes[1].ID, mockBikes[1].Lat, mockBikes[1].Long,
			mockBikes[1].Status, 1, mockBikes[1].CreatedAt, mockBikes[1].UpdatedAt, nil).
		AddRow(mockBikes[2].ID, mockBikes[2].Lat, mockBikes[2].Long,
			mockBikes[2].Status, nil, mockBikes[2].CreatedAt, mockBikes[2].UpdatedAt, nil)

	query := regexp.QuoteMeta("SELECT * FROM `bike`")

	s.mockDB.ExpectQuery(query).WillReturnRows(rows)
	actual, err := s.repositoryImpl.GetList(context.TODO())
	s.Equal(mockBikes, *actual)
	s.Nil(err)
}

func (s *BikeRepositoryTestSuite) TestGetList_Failed() {
	query := regexp.QuoteMeta("SELECT * FROM `bike`")

	s.mockDB.ExpectQuery(query).WillReturnError(gorm.ErrRecordNotFound)
	actual, err := s.repositoryImpl.GetList(context.TODO())
	s.Nil(actual)
	s.Equal(gorm.ErrRecordNotFound, err)
}

func (s *BikeRepositoryTestSuite) TestGetByID_Success() {
	mockTime := time.Time{}
	lat := decimal.NewFromFloat(50.119504)
	long := decimal.NewFromFloat(8.638137)
	mockNilUserID := sql.NullInt64{
		Valid: false,
		Int64: 0,
	}
	mockBike := domain.Bike{
		ID:        1,
		Lat:       &lat,
		Long:      &long,
		UserID:    mockNilUserID,
		Status:    domain.BikeStatusAvailable,
		CreatedAt: mockTime,
		UpdatedAt: mockTime,
		DeletedAt: gorm.DeletedAt{
			Valid: false,
			Time:  time.Time{},
		},
	}
	query := regexp.QuoteMeta("SELECT * FROM `bike` WHERE id = ?")
	row := sqlmock.NewRows([]string{"id", "lat", "long", "status", "user_id", "created_at", "updated_at", "deleted_at"}).
		AddRow(mockBike.ID, mockBike.Lat, mockBike.Long,
			mockBike.Status, nil, mockBike.CreatedAt, mockBike.UpdatedAt, nil)
	s.mockDB.ExpectQuery(query).WithArgs(sqlmock.AnyArg()).WillReturnRows(row)
	actual, err := s.repositoryImpl.GetByID(context.TODO(), int64(1))
	s.Equal(mockBike, *actual)
	s.Nil(err)
}

func (s *BikeRepositoryTestSuite) TestGetByID_Failed() {
	query := regexp.QuoteMeta("SELECT * FROM `bike` WHERE id = ?")
	s.mockDB.ExpectQuery(query).WithArgs(sqlmock.AnyArg()).WillReturnError(gorm.ErrRecordNotFound)
	actual, err := s.repositoryImpl.GetByID(context.TODO(), int64(1))
	s.Nil(actual)
	s.Equal(gorm.ErrRecordNotFound, err)
}

func (s *BikeRepositoryTestSuite) TestUpdate_Success() {
	mockUserID := sql.NullInt64{
		Valid: true,
		Int64: 1,
	}
	updatedVariables := domain.Bike{
		Status: domain.BikeStatusRented,
		UserID: mockUserID,
	}
	query := regexp.QuoteMeta("UPDATE `bike` SET `status`=?,`user_id`=?,`updated_at`=? WHERE id = ?")
	s.mockDB.ExpectBegin()
	s.mockDB.ExpectExec(query).WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg()).WillReturnResult(sqlmock.NewResult(1, 1))
	s.mockDB.ExpectCommit()
	err := s.repositoryImpl.UpdateStatusAndUserID(context.TODO(), &updatedVariables)
	s.Nil(err)
}

func (s *BikeRepositoryTestSuite) TestUpdate_Failed() {
	mockUserID := sql.NullInt64{
		Valid: true,
		Int64: 1,
	}
	updatedVariables := domain.Bike{
		Status: domain.BikeStatusRented,
		UserID: mockUserID,
	}
	query := regexp.QuoteMeta("UPDATE `bike` SET `status`=?,`user_id`=?,`updated_at`=? WHERE id = ?")
	s.mockDB.ExpectBegin()
	s.mockDB.ExpectExec(query).WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg()).WillReturnError(gorm.ErrRecordNotFound)
	s.mockDB.ExpectRollback()
	err := s.repositoryImpl.UpdateStatusAndUserID(context.TODO(), &updatedVariables)
	s.Equal(gorm.ErrRecordNotFound, err)
}

func (s *BikeRepositoryTestSuite) TestCountByUserID_Success() {
	query := regexp.QuoteMeta("SELECT count(*) FROM `bike` WHERE user_id = ? AND `bike`.`deleted_at` IS NULL")
	row := sqlmock.NewRows([]string{"count"}).
		AddRow(int64(1))
	s.mockDB.ExpectQuery(query).WithArgs(sqlmock.AnyArg()).WillReturnRows(row)
	actual, err := s.repositoryImpl.CountByUserID(context.TODO(), int64(1))
	s.Equal(int64(1), actual)
	s.Nil(err)
}

func (s *BikeRepositoryTestSuite) TestCountByUserID_Failed() {
	query := regexp.QuoteMeta("SELECT count(*) FROM `bike` WHERE user_id = ? AND `bike`.`deleted_at` IS NULL")
	s.mockDB.ExpectQuery(query).WithArgs(sqlmock.AnyArg()).WillReturnError(gorm.ErrRecordNotFound)
	actual, err := s.repositoryImpl.CountByUserID(context.TODO(), int64(1))
	s.Equal(int64(0), actual)
	s.Equal(gorm.ErrRecordNotFound, err)
}
