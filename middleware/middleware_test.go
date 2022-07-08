package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/suite"
)

type BikeHandlerTestSuite struct {
	suite.Suite
	echo *echo.Echo
}

func (s *BikeHandlerTestSuite) SetupTest() {
	e := echo.New()
	e.Use(session.Middleware(sessions.NewCookieStore([]byte("my-secret"))))
	e.Use(Authorize)
	// e.Use(AddHeaderXRequestID)
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
	respBody := `"unauthorized"
`
	s.Equal(http.StatusUnauthorized, rec.Code)
	s.Equal(respBody, rec.Body.String())
}
