package user

import (
	"context"
	"net/http"
	"net/http/httptest"
	"shared-bike/apperrors"
	"shared-bike/domain"
	"shared-bike/pkg/user/mocks"
	"strings"
	"testing"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/suite"
)

type UserHandlerTestSuite struct {
	suite.Suite
	mockUseCase *mocks.IUseCase
	echo        *echo.Echo
	handlerImpl *handlerImpl
}

func (s *UserHandlerTestSuite) SetupTest() {
	mockUseCase := &mocks.IUseCase{}
	s.mockUseCase = mockUseCase
	e := echo.New()
	s.echo = e
	handler := NewHandler(mockUseCase)
	s.handlerImpl = handler
}
func TestUserHandlerTestSuite(t *testing.T) {
	suite.Run(t, new(UserHandlerTestSuite))
}

func (s *UserHandlerTestSuite) TestLogin_Success() {
	var (
		mockContext = context.TODO()
		mockBody    = domain.LoginBody{
			Username: "testUsername",
			Password: "testPassword",
		}
		mockResult = domain.UserDTO{
			ID:       1,
			Username: "testUsername",
			Name:     "testName",
		}
		loginBody = `{"username":"testUsername","password":"testPassword"}`
	)
	s.mockUseCase.On("Login", mockContext, mockBody).Return(mockResult, nil)
	req := httptest.NewRequest(http.MethodPost, "/users/login", strings.NewReader(loginBody))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := s.echo.NewContext(req, rec)
	c.Set("_session_store", sessions.NewCookieStore([]byte("mockSecret")))
	respBody := `null
`
	c.SetPath("/users/login")
	s.NoError(s.handlerImpl.Login(c))
	s.Equal(http.StatusNoContent, rec.Code)
	s.Equal(respBody, rec.Body.String())
}

func (s *UserHandlerTestSuite) TestLogin_SessionInitError() {
	var (
		mockContext = context.TODO()
		mockBody    = domain.LoginBody{
			Username: "testUsername",
			Password: "testPassword",
		}
		mockResult = domain.UserDTO{
			ID:       1,
			Username: "testUsername",
			Name:     "testName",
		}
		loginBody = `{"username":"testUsername","password":"testPassword"}`
	)
	s.mockUseCase.On("Login", mockContext, mockBody).Return(mockResult, nil)
	req := httptest.NewRequest(http.MethodPost, "/users/login", strings.NewReader(loginBody))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := s.echo.NewContext(req, rec)
	respBody := `"internal server error"
`
	c.SetPath("/users/login")
	s.NoError(s.handlerImpl.Login(c))
	s.Equal(http.StatusInternalServerError, rec.Code)
	s.Equal(respBody, rec.Body.String())
}

func (s *UserHandlerTestSuite) TestLogin_InvalidBody() {
	var (
		mockContext = context.TODO()
		mockBody    = domain.LoginBody{
			Username: "testUsername",
			Password: "testPassword",
		}
		mockResult = domain.UserDTO{
			ID:       1,
			Username: "testUsername",
			Name:     "testName",
		}
		loginBody = `{"username":"testUsername","password":"testPassword",}`
	)
	s.mockUseCase.On("Login", mockContext, mockBody).Return(mockResult, nil)
	req := httptest.NewRequest(http.MethodPost, "/users/login", strings.NewReader(loginBody))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := s.echo.NewContext(req, rec)
	respBody := `"invalid body"
`
	c.SetPath("/users/login")
	s.NoError(s.handlerImpl.Login(c))
	s.Equal(http.StatusBadRequest, rec.Code)
	s.Equal(respBody, rec.Body.String())
}

func (s *UserHandlerTestSuite) TestLogin_InternalError() {
	var (
		mockContext = context.TODO()
		mockBody    = domain.LoginBody{
			Username: "testUsername",
			Password: "testPassword",
		}
		loginBody = `{"username":"testUsername","password":"testPassword"}`
	)
	s.mockUseCase.On("Login", mockContext, mockBody).Return(domain.UserDTO{}, apperrors.ErrInternalServerError)
	req := httptest.NewRequest(http.MethodPost, "/users/login", strings.NewReader(loginBody))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := s.echo.NewContext(req, rec)
	respBody := `"internal server error"
`
	c.SetPath("/users/login")
	s.NoError(s.handlerImpl.Login(c))
	s.Equal(http.StatusInternalServerError, rec.Code)
	s.Equal(respBody, rec.Body.String())
}

func (s *UserHandlerTestSuite) TestRegister_Success() {
	var (
		mockContext = context.TODO()
		mockBody    = domain.RegisterBody{
			Username: "testUsername",
			Password: "testPassword",
			Name:     "mockName",
		}
		mockResult = domain.UserDTO{
			ID:       1,
			Username: "testUsername",
			Name:     "testName",
		}
		registerBody = `{"username":"testUsername","password":"testPassword", "name":"mockName"}`
	)
	s.mockUseCase.On("Register", mockContext, mockBody).Return(mockResult, nil)
	req := httptest.NewRequest(http.MethodPost, "/users/register", strings.NewReader(registerBody))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := s.echo.NewContext(req, rec)
	c.Set("_session_store", sessions.NewCookieStore([]byte("mockSecret")))
	respBody := `null
`
	c.SetPath("/users/register")
	s.NoError(s.handlerImpl.Register(c))
	s.Equal(http.StatusCreated, rec.Code)
	s.Equal(respBody, rec.Body.String())
}

func (s *UserHandlerTestSuite) TestRegister_SessionError() {
	var (
		mockContext = context.TODO()
		mockBody    = domain.RegisterBody{
			Username: "testUsername",
			Password: "testPassword",
			Name:     "mockName",
		}
		mockResult = domain.UserDTO{
			ID:       1,
			Username: "testUsername",
			Name:     "testName",
		}
		registerBody = `{"username":"testUsername","password":"testPassword", "name":"mockName"}`
	)
	s.mockUseCase.On("Register", mockContext, mockBody).Return(mockResult, nil)
	req := httptest.NewRequest(http.MethodPost, "/users/register", strings.NewReader(registerBody))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := s.echo.NewContext(req, rec)
	respBody := `"internal server error"
`
	c.SetPath("/users/register")
	s.NoError(s.handlerImpl.Register(c))
	s.Equal(http.StatusInternalServerError, rec.Code)
	s.Equal(respBody, rec.Body.String())
}

func (s *UserHandlerTestSuite) TestRegister_InvalidBody() {
	var (
		mockContext = context.TODO()
		mockBody    = domain.RegisterBody{
			Username: "testUsername",
			Password: "testPassword",
			Name:     "mockName",
		}
		mockResult = domain.UserDTO{
			ID:       1,
			Username: "testUsername",
			Name:     "testName",
		}
		registerBody = `{"username":"testUsername","password":"testPassword","name":"mockName",}`
	)
	s.mockUseCase.On("Register", mockContext, mockBody).Return(mockResult, nil)
	req := httptest.NewRequest(http.MethodPost, "/users/register", strings.NewReader(registerBody))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := s.echo.NewContext(req, rec)
	respBody := `"invalid body"
`
	c.SetPath("/users/register")
	s.NoError(s.handlerImpl.Register(c))
	s.Equal(http.StatusBadRequest, rec.Code)
	s.Equal(respBody, rec.Body.String())
}

func (s *UserHandlerTestSuite) TestRegister_InternalError() {
	var (
		mockContext = context.TODO()
		mockBody    = domain.RegisterBody{
			Username: "testUsername",
			Password: "testPassword",
			Name:     "mockName",
		}
		registerBody = `{"username":"testUsername","password":"testPassword","name":"mockName"}`
	)
	s.mockUseCase.On("Register", mockContext, mockBody).Return(domain.UserDTO{}, apperrors.ErrInternalServerError)
	req := httptest.NewRequest(http.MethodPost, "/users/register", strings.NewReader(registerBody))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := s.echo.NewContext(req, rec)
	respBody := `"internal server error"
`
	c.SetPath("/users/register")
	s.NoError(s.handlerImpl.Register(c))
	s.Equal(http.StatusInternalServerError, rec.Code)
	s.Equal(respBody, rec.Body.String())
}
