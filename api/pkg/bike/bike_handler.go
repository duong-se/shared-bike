package bike

import (
	"fmt"
	"net/http"
	"strconv"

	"shared-bike/apperrors"
	"shared-bike/domain"
	"shared-bike/middleware"

	"github.com/golang-jwt/jwt"
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

// GetAllBike godoc
// @Summary      Get all bikes
// @Description  API for getting all bikes
// @Tags         bikes
// @Accept       json
// @Produce      json
// @Success      200  {array}   []domain.BikeDTO "Success"
// @Failure      500  {string}  string 	"internal server error"
// @Router       /bikes [get]
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

// Rent godoc
// @Summary      Rent a bike
// @Description  API for renting a bike
// @Tags         bikes
// @Accept       json
// @Produce      json
// @Param 			 id 	path  		string 		true 								"bike id"
// @Success      200  {object}  domain.BikeDTO 							  "Success"
// @Failure      400  {string}  string 												"invalid bike id | cannot rent because you have already rented a bike | user not exists or inactive | bike not found | cannot rent because bike is rented"
// @Failure      500  {string}  string 												"internal server error"
// @Router       /bikes/{id}/rent [patch]
func (h *handlerImpl) Rent(c echo.Context) error {
	var (
		ctx    = c.Request().Context()
		bikeID int64
		err    error
	)
	bikeIDStr := c.Param("id")
	if bikeID, err = strconv.ParseInt(bikeIDStr, 10, 64); err != nil {
		c.Logger().Error(fmt.Sprintf("[BikeHandler.Rent] invalid bike %s", bikeIDStr), err)
		return c.JSON(apperrors.GetStatusCode(apperrors.ErrInvalidBikeID), apperrors.ErrInvalidBikeID.Error())
	}
	user := c.Get(middleware.UserKey).(*jwt.Token)
	claims := user.Claims.(*domain.Claims)
	userID := claims.ID
	request := domain.RentOrReturnRequestPayload{
		ID:     bikeID,
		UserID: userID,
	}
	c.Logger().Info(fmt.Sprintf("[BikeHandler.Rent] user %d is renting bike %s", userID, bikeIDStr))
	bikes, err := h.useCase.Rent(ctx, request)
	if err != nil {
		c.Logger().Error(fmt.Sprintf("[BikeHandler.Rent] user %d rent bike %s failed", userID, bikeIDStr), err)
		return c.JSON(apperrors.GetStatusCode(err), err.Error())
	}
	c.Logger().Info(fmt.Sprintf("[BikeHandler.Rent] user %d rent bike %s success", userID, bikeIDStr))
	return c.JSON(http.StatusOK, bikes)
}

// Return godoc
// @Summary      Return a bike
// @Description  API for returning a bike
// @Tags         bikes
// @Accept       json
// @Produce      json
// @Param 			 id 	path  		string 		true 								"bike id"
// @Success      200  {object}  []domain.BikeDTO 							"Success"
// @Failure      400  {string}  string 												"invalid bike id | bike not found | cannot return because bike is available | cannot return because bike is not yours"
// @Failure      500  {string}  string 												"internal server error"
// @Router       /bikes/{id}/return [patch]
func (h *handlerImpl) Return(c echo.Context) error {
	var (
		ctx    = c.Request().Context()
		bikeID int64
		err    error
	)
	bikeIDStr := c.Param("id")
	if bikeID, err = strconv.ParseInt(bikeIDStr, 10, 64); err != nil {
		c.Logger().Error(fmt.Sprintf("[BikeHandler.Return] invalid bike id %s", bikeIDStr), err)
		return c.JSON(apperrors.GetStatusCode(apperrors.ErrInvalidBikeID), apperrors.ErrInvalidBikeID.Error())
	}
	user := c.Get(middleware.UserKey).(*jwt.Token)
	claims := user.Claims.(*domain.Claims)
	userID := claims.ID
	request := domain.RentOrReturnRequestPayload{
		ID:     bikeID,
		UserID: userID,
	}
	c.Logger().Info(fmt.Sprintf("[BikeHandler.Return] user %d is returning bike %s", userID, bikeIDStr))
	bikes, err := h.useCase.Return(ctx, request)
	if err != nil {
		c.Logger().Error(fmt.Sprintf("[BikeHandler.Return] user %d is return bike %s failed", userID, bikeIDStr), err)
		return c.JSON(apperrors.GetStatusCode(err), err.Error())
	}
	c.Logger().Info(fmt.Sprintf("[BikeHandler.Return] user %d return bike %s success", userID, bikeIDStr))
	return c.JSON(http.StatusOK, bikes)
}
