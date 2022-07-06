package user

import (
	"context"
	"regexp"
	"testing"
	"time"

	"shared-bike/domain"

	"github.com/stretchr/testify/suite"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type UserRepositoryTestSuite struct {
	suite.Suite
	mockDB         sqlmock.Sqlmock
	repositoryImpl *repositoryImpl
}

func (s *UserRepositoryTestSuite) SetupTest() {
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
func TestUserRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(UserRepositoryTestSuite))
}

func (s *UserRepositoryTestSuite) TestGetByUsername_Success() {
	mockTime := time.Time{}
	mockUser := domain.User{
		ID:        1,
		Username:  "testUsername",
		Password:  "testPassword",
		Name:      "testName",
		CreatedAt: mockTime,
		UpdatedAt: mockTime,
		DeletedAt: gorm.DeletedAt{Valid: false},
	}

	rows := sqlmock.NewRows([]string{"id", "username", "password", "name", "created_at", "updated_at"}).
		AddRow(mockUser.ID, mockUser.Username, mockUser.Password,
			mockUser.Name, mockUser.CreatedAt, mockUser.UpdatedAt)

	query := regexp.QuoteMeta("SELECT * FROM `user` WHERE username = ? AND `user`.`deleted_at` IS NULL ORDER BY `user`.`id` LIMIT 1")
	s.mockDB.ExpectQuery(query).WithArgs(sqlmock.AnyArg()).WillReturnRows(rows)
	actual, err := s.repositoryImpl.GetByUsername(context.TODO(), "testUsername")
	s.Equal(mockUser, *actual)
	s.Nil(err)
}

func (s *UserRepositoryTestSuite) TestGetByUsername_Failed() {
	query := regexp.QuoteMeta("SELECT * FROM `user` WHERE username = ? AND `user`.`deleted_at` IS NULL ORDER BY `user`.`id` LIMIT 1")
	s.mockDB.ExpectQuery(query).WithArgs(sqlmock.AnyArg()).WillReturnError(gorm.ErrRecordNotFound)
	actual, err := s.repositoryImpl.GetByUsername(context.TODO(), "testUsername")
	s.Nil(actual)
	s.Equal(gorm.ErrRecordNotFound, err)
}

func (s *UserRepositoryTestSuite) TestGetListByIDs_Success() {
	mockTime := time.Time{}
	mockUsers := []domain.User{
		{
			ID:        1,
			Username:  "testUsername",
			Password:  "testPassword",
			Name:      "testName",
			CreatedAt: mockTime,
			UpdatedAt: mockTime,
			DeletedAt: gorm.DeletedAt{Valid: false},
		},
		{
			ID:        2,
			Username:  "testUsername",
			Password:  "testPassword",
			Name:      "testName",
			CreatedAt: mockTime,
			UpdatedAt: mockTime,
			DeletedAt: gorm.DeletedAt{Valid: false},
		},
		{
			ID:        3,
			Username:  "testUsername",
			Password:  "testPassword",
			Name:      "testName",
			CreatedAt: mockTime,
			UpdatedAt: mockTime,
			DeletedAt: gorm.DeletedAt{Valid: false},
		},
	}

	rows := sqlmock.NewRows([]string{"id", "username", "password", "name", "created_at", "updated_at"}).
		AddRow(mockUsers[0].ID, mockUsers[0].Username, mockUsers[0].Password,
			mockUsers[0].Name, mockUsers[0].CreatedAt, mockUsers[0].UpdatedAt).
		AddRow(mockUsers[1].ID, mockUsers[1].Username, mockUsers[1].Password,
			mockUsers[1].Name, mockUsers[1].CreatedAt, mockUsers[1].UpdatedAt).
		AddRow(mockUsers[2].ID, mockUsers[2].Username, mockUsers[2].Password,
			mockUsers[2].Name, mockUsers[2].CreatedAt, mockUsers[2].UpdatedAt)

	query := regexp.QuoteMeta("SELECT * FROM `user` WHERE id IN (?,?,?) AND `user`.`deleted_at` IS NULL")
	s.mockDB.ExpectQuery(query).WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg()).WillReturnRows(rows)
	actual, err := s.repositoryImpl.GetListByIDs(context.TODO(), []int64{1, 2, 3})
	s.Equal(mockUsers, *actual)
	s.Nil(err)
}

func (s *UserRepositoryTestSuite) TestGetListByIDs_Failed() {
	query := regexp.QuoteMeta("SELECT * FROM `user` WHERE id IN (?,?,?) AND `user`.`deleted_at` IS NULL")
	s.mockDB.ExpectQuery(query).WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg()).WillReturnError(gorm.ErrRecordNotFound)
	actual, err := s.repositoryImpl.GetListByIDs(context.TODO(), []int64{1, 2, 3})
	s.Nil(actual)
	s.Equal(gorm.ErrRecordNotFound, err)
}

func (s *UserRepositoryTestSuite) TestCreate_Success() {
	userID := int64(1)
	mockTime := time.Time{}
	mockNewUser := domain.User{
		ID:        userID,
		Username:  "testUsername",
		Password:  "testPassword",
		Name:      "testName",
		CreatedAt: mockTime,
		UpdatedAt: mockTime,
		DeletedAt: gorm.DeletedAt{Valid: false},
	}
	query := regexp.QuoteMeta("INSERT INTO `user` (`username`,`password`,`name`,`created_at`,`updated_at`,`deleted_at`,`id`) VALUES (?,?,?,?,?,?,?)")
	s.mockDB.ExpectBegin()
	s.mockDB.ExpectExec(query).WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg()).WillReturnResult(sqlmock.NewResult(1, 1))
	s.mockDB.ExpectCommit()
	err := s.repositoryImpl.Create(context.TODO(), &mockNewUser)
	s.Nil(err)
}

func (s *UserRepositoryTestSuite) TestCreate_Failed() {
	userID := int64(1)
	mockTime := time.Time{}
	mockNewUser := domain.User{
		ID:        userID,
		Username:  "testUsername",
		Password:  "testPassword",
		Name:      "testName",
		CreatedAt: mockTime,
		UpdatedAt: mockTime,
		DeletedAt: gorm.DeletedAt{Valid: false},
	}
	query := regexp.QuoteMeta("INSERT INTO `user` (`username`,`password`,`name`,`created_at`,`updated_at`,`deleted_at`,`id`) VALUES (?,?,?,?,?,?,?)")
	s.mockDB.ExpectBegin()
	s.mockDB.ExpectExec(query).WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg()).WillReturnError(gorm.ErrRecordNotFound)
	s.mockDB.ExpectRollback()
	err := s.repositoryImpl.Create(context.TODO(), &mockNewUser)
	s.Equal(gorm.ErrRecordNotFound, err)
}
