package bike

import (
	"net/http"
	"strconv"

	"shared-bike/apperrors"
	"shared-bike/domain"
	"shared-bike/middleware"

	"github.com/labstack/echo/v4"
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
	c.Logger().Info("[BikeHandler.GetAllBike] starting")
	ctx := c.Request().Context()
	bikes, err := h.useCase.GetAllBike(ctx)
	if err != nil {
		c.Logger().Error("[BikeHandler.GetAllBike] cannot get all bikes", err)
		return c.JSON(apperrors.GetStatusCode(err), err.Error())
	}
	c.Logger().Info("[BikeHandler.GetAllBike] success")
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
		c.Logger().Errorf("[BikeHandler.Rent] invalid bike %s", bikeIDStr, err)
		return c.JSON(http.StatusBadRequest, "invalid bike id")
	}
	userID := c.Get(middleware.UserIDKey).(int64)
	request := domain.RentOrReturnRequestPayload{
		ID:     bikeID,
		UserID: userID,
	}
	c.Logger().Infof("[BikeHandler.Rent] user %d is renting bike %s", userID, bikeIDStr)
	bikes, err := h.useCase.Rent(ctx, request)
	if err != nil {
		c.Logger().Errorf("[BikeHandler.Rent] user %d rent bike %s failed", userID, bikeIDStr, err)
		return c.JSON(apperrors.GetStatusCode(err), err.Error())
	}
	c.Logger().Infof("[BikeHandler.Rent] user %d rent bike %s success", userID, bikeIDStr)
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
		c.Logger().Errorf("[BikeHandler.Return] invalid bike id %s", bikeIDStr, err)
		return c.JSON(http.StatusBadRequest, "invalid bike id")
	}
	userID := c.Get(middleware.UserIDKey).(int64)
	request := domain.RentOrReturnRequestPayload{
		ID:     bikeID,
		UserID: userID,
	}
	c.Logger().Infof("[BikeHandler.Return] user %d is returning bike %s", userID, bikeIDStr)
	bikes, err := h.useCase.Return(ctx, request)
	if err != nil {
		c.Logger().Errorf("[BikeHandler.Return] user %d is return bike %s failed", userID, bikeIDStr, err)
		return c.JSON(apperrors.GetStatusCode(err), err.Error())
	}
	c.Logger().Infof("[BikeHandler.Return] user %d return bike %s success", userID, bikeIDStr)
	return c.JSON(http.StatusOK, bikes)
}
