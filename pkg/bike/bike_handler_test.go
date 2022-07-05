package bike

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/duong-se/shared-bike/apperrors"
	"github.com/duong-se/shared-bike/domain"
	"github.com/duong-se/shared-bike/middleware"
	"github.com/duong-se/shared-bike/pkg/bike/mocks"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/suite"
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
	s.echo = echo.New()
	handler := NewHandler(mockUseCase)
	s.handlerImpl = handler
}
func TestBikeHandlerTestSuite(t *testing.T) {
	suite.Run(t, new(BikeHandlerTestSuite))
}

func (s *BikeHandlerTestSuite) TestGetAll_Success() {
	var (
		mockContext = context.TODO()
		mockResult  = []domain.Bike{
			{
				ID:     1,
				Lat:    "50.119504",
				Long:   "8.638137",
				Status: domain.BikeStatusAvailable,
			},
			{
				ID:     1,
				Lat:    "50.119229",
				Long:   "8.640020",
				Status: domain.BikeStatusRented,
			},
			{
				ID:     1,
				Lat:    "50.120452",
				Long:   "8.650507",
				Status: domain.BikeStatusAvailable,
			},
		}
	)
	s.mockUseCase.On("GetAllBike", mockContext).Return(mockResult, nil)
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := s.echo.NewContext(req, rec)
	respBody := `[{"id":1,"lat":"50.119504","long":"8.638137","status":"available","userId":null,"createdAt":"0001-01-01T00:00:00Z","updatedAt":"0001-01-01T00:00:00Z","deletedAt":null},{"id":1,"lat":"50.119229","long":"8.640020","status":"rented","userId":null,"createdAt":"0001-01-01T00:00:00Z","updatedAt":"0001-01-01T00:00:00Z","deletedAt":null},{"id":1,"lat":"50.120452","long":"8.650507","status":"available","userId":null,"createdAt":"0001-01-01T00:00:00Z","updatedAt":"0001-01-01T00:00:00Z","deletedAt":null}]
`
	c.SetPath("/bikes")
	s.NoError(s.handlerImpl.GetAllBike(c))
	s.Equal(http.StatusOK, rec.Code)
	s.Equal(respBody, rec.Body.String())
}

func (s *BikeHandlerTestSuite) TestGetAll_Failed() {
	var (
		mockContext = context.TODO()
		mockResult  = []domain.Bike{}
	)
	s.mockUseCase.On("GetAllBike", mockContext).Return(mockResult, apperrors.ErrInternalServerError)
	req := httptest.NewRequest(http.MethodGet, "/", nil)
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
		mockResult  = domain.Bike{
			ID:     1,
			Lat:    "50.119504",
			Long:   "8.638137",
			Status: domain.BikeStatusRented,
			UserID: &userID,
		}
		mockInput = domain.RentOrReturnRequestPayload{
			UserID: 1,
			ID:     1,
		}
	)
	s.mockUseCase.On("Rent", mockContext, mockInput).Return(&mockResult, nil)
	req := httptest.NewRequest(http.MethodPatch, "/1/rent", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := s.echo.NewContext(req, rec)
	c.Set(middleware.UserIDKey, int64(1))
	respBody := `{"id":1,"lat":"50.119504","long":"8.638137","status":"rented","userId":1,"createdAt":"0001-01-01T00:00:00Z","updatedAt":"0001-01-01T00:00:00Z","deletedAt":null}
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
		mockResult  = domain.Bike{}
		mockInput   = domain.RentOrReturnRequestPayload{
			UserID: 1,
			ID:     1,
		}
	)
	s.mockUseCase.On("Rent", mockContext, mockInput).Return(&mockResult, apperrors.ErrBikeNotFound)
	req := httptest.NewRequest(http.MethodPatch, "/1/rent", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := s.echo.NewContext(req, rec)
	c.Set(middleware.UserIDKey, int64(1))
	respBody := `"your request bike not found"
`
	c.SetPath("/bikes/:id/rent")
	c.SetParamNames("id")
	c.SetParamValues("1")
	s.NoError(s.handlerImpl.Rent(c))
	s.Equal(http.StatusNotFound, rec.Code)
	s.Equal(respBody, rec.Body.String())
}

func (s *BikeHandlerTestSuite) TestRent_FailedParams() {
	req := httptest.NewRequest(http.MethodPatch, "/1/rent", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := s.echo.NewContext(req, rec)
	c.Set(middleware.UserIDKey, int64(1))
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
		mockResult  = domain.Bike{
			ID:     1,
			Lat:    "50.119504",
			Long:   "8.638137",
			Status: domain.BikeStatusAvailable,
			UserID: nil,
		}
		mockInput = domain.RentOrReturnRequestPayload{
			UserID: 1,
			ID:     1,
		}
	)
	s.mockUseCase.On("Return", mockContext, mockInput).Return(&mockResult, nil)
	req := httptest.NewRequest(http.MethodPatch, "/1/return", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := s.echo.NewContext(req, rec)
	c.Set(middleware.UserIDKey, int64(1))
	respBody := `{"id":1,"lat":"50.119504","long":"8.638137","status":"available","userId":null,"createdAt":"0001-01-01T00:00:00Z","updatedAt":"0001-01-01T00:00:00Z","deletedAt":null}
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
		mockResult  = domain.Bike{}
		mockInput   = domain.RentOrReturnRequestPayload{
			UserID: 1,
			ID:     1,
		}
	)
	s.mockUseCase.On("Return", mockContext, mockInput).Return(&mockResult, apperrors.ErrBikeNotFound)
	req := httptest.NewRequest(http.MethodPatch, "/1/return", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := s.echo.NewContext(req, rec)
	c.Set(middleware.UserIDKey, int64(1))
	respBody := `"your request bike not found"
`
	c.SetPath("/bikes/:id/return")
	c.SetParamNames("id")
	c.SetParamValues("1")
	s.NoError(s.handlerImpl.Return(c))
	s.Equal(http.StatusNotFound, rec.Code)
	s.Equal(respBody, rec.Body.String())
}

func (s *BikeHandlerTestSuite) TestReturn_FailedParams() {
	req := httptest.NewRequest(http.MethodPatch, "/1/return", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := s.echo.NewContext(req, rec)
	c.Set(middleware.UserIDKey, int64(1))
	respBody := `"invalid bike id"
`
	c.SetPath("/bikes/:id/return")
	c.SetParamNames("id")
	c.SetParamValues("abc")
	s.NoError(s.handlerImpl.Return(c))
	s.Equal(http.StatusBadRequest, rec.Code)
	s.Equal(respBody, rec.Body.String())
}
