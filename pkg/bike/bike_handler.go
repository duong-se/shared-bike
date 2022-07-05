package bike

import (
	"net/http"
	"strconv"

	"github.com/duong-se/shared-bike/apperrors"
	"github.com/duong-se/shared-bike/domain"
	"github.com/duong-se/shared-bike/middleware"
	"github.com/labstack/echo"
)

type handlerImpl struct {
	useCase IUseCase
}

func NewHandler(useCase IUseCase) *handlerImpl {
	return &handlerImpl{
		useCase: useCase,
	}
}

func (h *handlerImpl) GetAllBike(c echo.Context) error {
	ctx := c.Request().Context()
	bikes, err := h.useCase.GetAllBike(ctx)
	if err != nil {
		return c.JSON(apperrors.GetStatusCode(err), err.Error())
	}
	return c.JSON(http.StatusOK, bikes)
}

func (h *handlerImpl) Rent(c echo.Context) error {
	var (
		ctx    = c.Request().Context()
		bikeID int64
		err    error
	)
	bikeIDStr := c.Param("id")
	if bikeID, err = strconv.ParseInt(bikeIDStr, 10, 64); err != nil {
		return c.JSON(http.StatusBadRequest, "invalid bike id")
	}
	userID := c.Get(middleware.UserIDKey).(int64)
	request := domain.RentOrReturnRequestPayload{
		ID:     bikeID,
		UserID: userID,
	}
	bikes, err := h.useCase.Rent(ctx, request)
	if err != nil {
		return c.JSON(apperrors.GetStatusCode(err), err.Error())
	}
	return c.JSON(http.StatusOK, bikes)
}

func (h *handlerImpl) Return(c echo.Context) error {
	var (
		ctx    = c.Request().Context()
		bikeID int64
		err    error
	)
	bikeIDStr := c.Param("id")
	if bikeID, err = strconv.ParseInt(bikeIDStr, 10, 64); err != nil {
		return c.JSON(http.StatusBadRequest, "invalid bike id")
	}
	userID := c.Get(middleware.UserIDKey).(int64)
	request := domain.RentOrReturnRequestPayload{
		ID:     bikeID,
		UserID: userID,
	}
	bikes, err := h.useCase.Return(ctx, request)
	if err != nil {
		return c.JSON(apperrors.GetStatusCode(err), err.Error())
	}
	return c.JSON(http.StatusOK, bikes)
}
