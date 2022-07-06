package domain

import (
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
)

type UserDomainTestSuite struct {
	suite.Suite
	user *User
}

func (s *UserDomainTestSuite) SetupTest() {
	mockTime := time.Time{}
	user := User{
		ID:        1,
		Username:  "testUsername",
		Password:  "$2a$10$Mjx4fmq9ykGxlqlT/l9yGuojZ0FLV8QmrDhGwxmdE3QdkaXQgCcMG",
		Name:      "testName",
		CreatedAt: mockTime,
		UpdatedAt: mockTime,
		DeletedAt: gorm.DeletedAt{Valid: false},
	}
	s.user = &user
}

func TestUserDomainTestSuite(t *testing.T) {
	suite.Run(t, new(UserDomainTestSuite))
}

func (s *UserDomainTestSuite) TestHashPassword_Success() {
	hashedPassword, err := s.user.HashPassword("testPassword")
	s.Nil(err)
	s.NotNil(hashedPassword)
}

func (s *UserDomainTestSuite) TestCheckPasswordIsEqual_Success() {
	isValid := s.user.ValidatePassword("testPassword")
	s.True(isValid)
}

func (s *UserDomainTestSuite) TestCheckPasswordIsEqual_Failed() {
	isValid := s.user.ValidatePassword("testPassword1")
	s.False(isValid)
}

func (s *UserDomainTestSuite) TestTableName_Success() {
	tableName := s.user.TableName()
	s.Equal("user", tableName)
}
