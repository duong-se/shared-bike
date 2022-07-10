package bike

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"shared-bike/apperrors"
	"shared-bike/domain"
	"shared-bike/middleware"
	"shared-bike/pkg/bike/mocks"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
)

type BikeHandlerTestSuite struct {
	suite.Suite
	mockUseCase *mocks.IUseCase
	echo        *echo.Echo
	handlerImpl *handlerImpl
}

func (s *BikeHandlerTestSuite) SetupTest() {
	mockUseCase := &mocks.IUseCase{}
	s.mockUseCase = mockUseCase
	e := echo.New()
	e.Use(session.Middleware(sessions.NewCookieStore([]byte("mockSecret"))))
	s.echo = e
	handler := NewHandler(mockUseCase)
	s.handlerImpl = handler
}
func TestBikeHandlerTestSuite(t *testing.T) {
	suite.Run(t, new(BikeHandlerTestSuite))
}

func (s *BikeHandlerTestSuite) TestGetAll_Success() {
	var (
		mockContext    = context.TODO()
		mockTime       = time.Time{}
		mockUserID     = int64(1)
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
				ID:               1,
				Lat:              "50.119504",
				Long:             "8.638137",
				Status:           domain.BikeStatusAvailable,
				UserID:           0,
				NameOfRenter:     "",
				UsernameOfRenter: "",
			},
			{
				ID:               1,
				Lat:              "50.119229",
				Long:             "8.640020",
				Status:           domain.BikeStatusRented,
				UserID:           mockUserID,
				NameOfRenter:     mockUserResult[0].Name,
				UsernameOfRenter: mockUserResult[0].Username,
			},
			{
				ID:               1,
				Lat:              "50.120452",
				Long:             "8.650507",
				Status:           domain.BikeStatusAvailable,
				UserID:           0,
				NameOfRenter:     "",
				UsernameOfRenter: "",
			},
		}
	)
	s.mockUseCase.On("GetAllBike", mockContext).Return(mockResult, nil)
	req := httptest.NewRequest(http.MethodGet, "/bikes", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := s.echo.NewContext(req, rec)
	respBody := `[{"id":1,"lat":"50.119504","long":"8.638137","status":"available","userId":0,"nameOfRenter":"","usernameOfRenter":""},{"id":1,"lat":"50.119229","long":"8.640020","status":"rented","userId":1,"nameOfRenter":"testName","usernameOfRenter":"testUsername"},{"id":1,"lat":"50.120452","long":"8.650507","status":"available","userId":0,"nameOfRenter":"","usernameOfRenter":""}]
`
	c.SetPath("/bikes")
	s.NoError(s.handlerImpl.GetAllBike(c))
	s.Equal(http.StatusOK, rec.Code)
	s.Equal(respBody, rec.Body.String())
}

func (s *BikeHandlerTestSuite) TestGetAll_Failed() {
	var (
		mockContext = context.TODO()
	)
	s.mockUseCase.On("GetAllBike", mockContext).Return(nil, apperrors.ErrInternalServerError)
	req := httptest.NewRequest(http.MethodGet, "/bikes", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := s.echo.NewContext(req, rec)
	respBody := `"internal server error"
`
	c.SetPath("/bikes")
	s.handlerImpl.GetAllBike(c)
	s.Equal(http.StatusInternalServerError, rec.Code)
	s.Equal(respBody, rec.Body.String())
}

func (s *BikeHandlerTestSuite) TestRent_Success() {
	var (
		userID      = int64(1)
		mockContext = context.TODO()
		name        = "mockName"
		username    = "mockUserName"
		mockResult  = domain.BikeDTO{
			ID:               1,
			Lat:              "50.119504",
			Long:             "8.638137",
			Status:           domain.BikeStatusRented,
			UserID:           userID,
			NameOfRenter:     name,
			UsernameOfRenter: username,
		}
		mockInput = domain.RentOrReturnRequestPayload{
			UserID: 1,
			ID:     1,
		}
	)
	s.mockUseCase.On("Rent", mockContext, mockInput).Return(mockResult, nil)
	req := httptest.NewRequest(http.MethodPatch, "/bikes/1/rent", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := s.echo.NewContext(req, rec)
	c.Set(middleware.UserKey, domain.Claims{ID: 1, Name: "TestUser", Username: "TestUserName"})
	respBody := `{"id":1,"lat":"50.119504","long":"8.638137","status":"rented","userId":1,"nameOfRenter":"mockName","usernameOfRenter":"mockUserName"}
`
	c.SetPath("/bikes/:id/rent")
	c.SetParamNames("id")
	c.SetParamValues("1")
	s.NoError(s.handlerImpl.Rent(c))
	s.Equal(http.StatusOK, rec.Code)
	s.Equal(respBody, rec.Body.String())
}

func (s *BikeHandlerTestSuite) TestRent_FailedUseCase() {
	var (
		mockContext = context.TODO()
		mockResult  = domain.BikeDTO{}
		mockInput   = domain.RentOrReturnRequestPayload{
			UserID: 1,
			ID:     1,
		}
	)
	s.mockUseCase.On("Rent", mockContext, mockInput).Return(mockResult, apperrors.ErrBikeNotFound)
	req := httptest.NewRequest(http.MethodPatch, "/bikes/1/rent", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := s.echo.NewContext(req, rec)
	c.Set(middleware.UserKey, domain.Claims{ID: 1, Name: "TestUser", Username: "TestUserName"})
	respBody := `"bike not found"
`
	c.SetPath("/bikes/:id/rent")
	c.SetParamNames("id")
	c.SetParamValues("1")
	s.NoError(s.handlerImpl.Rent(c))
	s.Equal(http.StatusNotFound, rec.Code)
	s.Equal(respBody, rec.Body.String())
}

func (s *BikeHandlerTestSuite) TestRent_FailedParams() {
	req := httptest.NewRequest(http.MethodPatch, "/bikes/1/rent", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := s.echo.NewContext(req, rec)
	c.Set(middleware.UserKey, domain.Claims{ID: 1, Name: "TestUser", Username: "TestUserName"})
	respBody := `"invalid bike id"
`
	c.SetPath("/bikes/:id/rent")
	c.SetParamNames("id")
	c.SetParamValues("abc")
	s.NoError(s.handlerImpl.Rent(c))
	s.Equal(http.StatusBadRequest, rec.Code)
	s.Equal(respBody, rec.Body.String())
}

func (s *BikeHandlerTestSuite) TestReturn_Success() {
	var (
		mockContext = context.TODO()
		mockResult  = domain.BikeDTO{
			ID:               1,
			Lat:              "50.119504",
			Long:             "8.638137",
			Status:           domain.BikeStatusAvailable,
			UserID:           0,
			NameOfRenter:     "",
			UsernameOfRenter: "",
		}
		mockInput = domain.RentOrReturnRequestPayload{
			UserID: 1,
			ID:     1,
		}
	)
	s.mockUseCase.On("Return", mockContext, mockInput).Return(mockResult, nil)
	req := httptest.NewRequest(http.MethodPatch, "/bikes/1/return", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := s.echo.NewContext(req, rec)
	c.Set(middleware.UserKey, domain.Claims{ID: 1, Name: "TestUser", Username: "TestUserName"})
	respBody := `{"id":1,"lat":"50.119504","long":"8.638137","status":"available","userId":0,"nameOfRenter":"","usernameOfRenter":""}
`
	c.SetPath("/bikes/:id/return")
	c.SetParamNames("id")
	c.SetParamValues("1")
	s.NoError(s.handlerImpl.Return(c))
	s.Equal(http.StatusOK, rec.Code)
	s.Equal(respBody, rec.Body.String())
}

func (s *BikeHandlerTestSuite) TestReturn_FailedUseCase() {
	var (
		mockContext = context.TODO()
		mockResult  = domain.BikeDTO{}
		mockInput   = domain.RentOrReturnRequestPayload{
			UserID: 1,
			ID:     1,
		}
	)
	s.mockUseCase.On("Return", mockContext, mockInput).Return(mockResult, apperrors.ErrBikeNotFound)
	req := httptest.NewRequest(http.MethodPatch, "/bikes/1/return", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := s.echo.NewContext(req, rec)
	c.Set(middleware.UserKey, domain.Claims{ID: 1, Name: "TestUser", Username: "TestUserName"})
	respBody := `"bike not found"
`
	c.SetPath("/bikes/:id/return")
	c.SetParamNames("id")
	c.SetParamValues("1")
	s.NoError(s.handlerImpl.Return(c))
	s.Equal(http.StatusNotFound, rec.Code)
	s.Equal(respBody, rec.Body.String())
}

func (s *BikeHandlerTestSuite) TestReturn_FailedParams() {
	req := httptest.NewRequest(http.MethodPatch, "/bikes/1/return", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := s.echo.NewContext(req, rec)
	c.Set(middleware.UserKey, domain.Claims{ID: 1, Name: "TestUser", Username: "TestUserName"})
	respBody := `"invalid bike id"
`
	c.SetPath("/bikes/:id/return")
	c.SetParamNames("id")
	c.SetParamValues("abc")
	s.NoError(s.handlerImpl.Return(c))
	s.Equal(http.StatusBadRequest, rec.Code)
	s.Equal(respBody, rec.Body.String())
}
