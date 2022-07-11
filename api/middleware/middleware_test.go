package middleware

import (
	"net/http"
	"net/http/httptest"
	"shared-bike/apperrors"
	"shared-bike/customlogger"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/suite"
)

type BikeHandlerTestSuite struct {
	suite.Suite
	echo *echo.Echo
}

func (s *BikeHandlerTestSuite) SetupTest() {
	e := echo.New()
	contextLogger := customlogger.NewContextLogger(e.Logger)
	e.Use(AddLoggerContext(contextLogger))
	e.GET("/", func(c echo.Context) error {
		return c.HTML(http.StatusOK, "body output")
	})
	e.GET("/api/v1/users/login", func(c echo.Context) error {
		return c.HTML(http.StatusOK, "login ok")
	})
	e.GET("/api/v1/users/register", func(c echo.Context) error {
		return c.HTML(http.StatusOK, "register ok")
	})
	e.GET("/health", func(c echo.Context) error {
		return c.HTML(http.StatusOK, "health ok")
	})
	e.GET("/swagger/index.html", func(c echo.Context) error {
		return c.HTML(http.StatusOK, "swagger ok")
	})
	s.echo = e
}
func TestBikeHandlerTestSuite(t *testing.T) {
	suite.Run(t, new(BikeHandlerTestSuite))
}

func (s *BikeHandlerTestSuite) TestAddLoggerContext_Success() {
	req := httptest.NewRequest(http.MethodGet, "/api/v1/users/login", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := s.echo.NewContext(req, rec)
	s.echo.ServeHTTP(rec, req)
	c.SetPath("/api/v1/users/login")
	respBody := "login ok"
	s.Equal(http.StatusOK, rec.Code)
	s.Equal(respBody, rec.Body.String())
}

func (s *BikeHandlerTestSuite) TestWhiteListAPI_False() {
	req := httptest.NewRequest(http.MethodGet, "/api/v2", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := s.echo.NewContext(req, rec)
	c.SetPath("/api/v2")
	result := WhiteListAPI(c)
	s.echo.ServeHTTP(rec, req)
	s.False(result)
}

func (s *BikeHandlerTestSuite) TestWhiteListAPI_TrueLogin() {
	req := httptest.NewRequest(http.MethodGet, "/api/v1/users/login", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := s.echo.NewContext(req, rec)
	c.SetPath("/api/v1/users/login")
	result := WhiteListAPI(c)
	s.echo.ServeHTTP(rec, req)
	s.True(result)
}

func (s *BikeHandlerTestSuite) TestWhiteListAPI_TrueRegister() {
	req := httptest.NewRequest(http.MethodGet, "/api/v1/users/register", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := s.echo.NewContext(req, rec)
	c.SetPath("/api/v1/users/register")
	result := WhiteListAPI(c)
	s.echo.ServeHTTP(rec, req)
	s.True(result)
}

func (s *BikeHandlerTestSuite) TestWhiteListAPI_TrueHealth() {
	req := httptest.NewRequest(http.MethodGet, "/health", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := s.echo.NewContext(req, rec)
	c.SetPath("/health")
	result := WhiteListAPI(c)
	s.echo.ServeHTTP(rec, req)
	s.True(result)
}

func (s *BikeHandlerTestSuite) TestWhiteListAPI_TrueSwagger() {
	req := httptest.NewRequest(http.MethodGet, "/swagger/index.html", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := s.echo.NewContext(req, rec)
	c.SetPath("/swagger/index.html")
	result := WhiteListAPI(c)
	s.echo.ServeHTTP(rec, req)
	s.True(result)
}

func (s *BikeHandlerTestSuite) TestCustomJWTError_Success() {
	req := httptest.NewRequest(http.MethodGet, "/swagger/index.html", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := s.echo.NewContext(req, rec)
	c.SetPath("/swagger/index.html")
	CustomJWTError(apperrors.ErrInternalServerError, c)
	s.echo.ServeHTTP(rec, req)
	s.Equal(http.StatusUnauthorized, rec.Code)
}
