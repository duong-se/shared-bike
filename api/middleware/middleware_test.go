package middleware

import (
	"net/http"
	"net/http/httptest"
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
	s.echo = e
}
func TestBikeHandlerTestSuite(t *testing.T) {
	suite.Run(t, new(BikeHandlerTestSuite))
}

func (s *BikeHandlerTestSuite) TestAuthorized_Failed() {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := s.echo.NewContext(req, rec)
	s.echo.ServeHTTP(rec, req)
	c.SetPath("/")
	respBody := "body output"
	s.Equal(http.StatusOK, rec.Code)
	s.Equal(respBody, rec.Body.String())
}
