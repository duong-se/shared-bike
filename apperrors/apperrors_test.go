package apperrors

import (
	"errors"
	"net/http"
	"testing"

	"github.com/stretchr/testify/suite"
)

type AppErrorsTestSuite struct {
	suite.Suite
}

func (s *AppErrorsTestSuite) SetupTest() {}

func TestAppErrorsTestSuite(t *testing.T) {
	suite.Run(t, new(AppErrorsTestSuite))
}

func (s *AppErrorsTestSuite) TestGetStatusCode_Return500() {
	err := ErrInternalServerError
	s.Equal(http.StatusInternalServerError, GetStatusCode(err))
}

func (s *AppErrorsTestSuite) TestGetStatusCode_Return401() {
	err := ErrUnauthorizeError
	s.Equal(http.StatusUnauthorized, GetStatusCode(err))
}

func (s *AppErrorsTestSuite) TestGetStatusCode_ReturnBikeRented() {
	err := ErrBikeRented
	s.Equal(http.StatusBadRequest, GetStatusCode(err))
}

func (s *AppErrorsTestSuite) TestGetStatusCode_ReturnBikeNotFound() {
	err := ErrBikeNotFound
	s.Equal(http.StatusNotFound, GetStatusCode(err))
}

func (s *AppErrorsTestSuite) TestGetStatusCode_Accept() {
	s.Equal(http.StatusOK, GetStatusCode(nil))
}

func (s *AppErrorsTestSuite) TestGetStatusCode_ReturnUserHasBikeAlready() {
	err := ErrUserHasBikeAlready
	s.Equal(http.StatusBadRequest, GetStatusCode(err))
}

func (s *AppErrorsTestSuite) TestGetStatusCode_ReturnBikeAvailable() {
	err := ErrBikeAvailable
	s.Equal(http.StatusBadRequest, GetStatusCode(err))
}

func (s *AppErrorsTestSuite) TestGetStatusCode_ReturnUserNotExisted() {
	err := ErrUserNotExisted
	s.Equal(http.StatusBadRequest, GetStatusCode(err))
}

func (s *AppErrorsTestSuite) TestGetStatusCode_ReturnUserLoginNotFound() {
	err := ErrUserLoginNotFound
	s.Equal(http.StatusNotFound, GetStatusCode(err))
}

func (s *AppErrorsTestSuite) TestGetStatusCode_ReturnFallback() {
	err := errors.New("mock")
	s.Equal(http.StatusInternalServerError, GetStatusCode(err))
}
