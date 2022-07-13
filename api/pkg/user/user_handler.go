package user

import (
	"net/http"
	"os"
	"time"

	"shared-bike/domain"

	"shared-bike/apperrors"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

type handlerImpl struct {
	usecase IUseCase
}

func NewHandler(usecase IUseCase) *handlerImpl {
	return &handlerImpl{
		usecase: usecase,
	}
}

// Login godoc
// @Summary      Login
// @Description  API for logining
// @Tags         users
// @Accept       json
// @Produce      json
// @Param    		 request  body      domain.LoginBody  true  "Login body"
// @Success      200	{object}  domain.Credentials 	"success"
// @Failure      400  {string}  string 							"invalid body"
// @Failure      404  {string}  string 							"username or password is wrong"
// @Failure      500  {string}  string 							"internal server error"
// @Router       /users/login [post]
func (h *handlerImpl) Login(c echo.Context) error {
	ctx := c.Request().Context()
	body := domain.LoginBody{}
	if err := c.Bind(&body); err != nil {
		c.Logger().Error("[UserHandler.Login] invalid body", err)
		return c.JSON(apperrors.GetStatusCode(apperrors.ErrInvalidBody), apperrors.ErrInvalidBody.Error())
	}
	c.Logger().Info("[UserHandler.Login] logging")
	user, err := h.usecase.Login(ctx, body)
	if err != nil {
		c.Logger().Error("[UserHandler.Login] login failed", err)
		return c.JSON(apperrors.GetStatusCode(err), err.Error())
	}
	c.Logger().Info("[UserHandler.Login] login success")
	token, err := h.signToken(user, 300)
	if err != nil {
		c.Logger().Error("[UserHandler.Login] sign token error", err)
		return c.JSON(apperrors.GetStatusCode(apperrors.ErrInternalServerError), apperrors.ErrInternalServerError.Error())
	}
	return c.JSON(http.StatusOK, domain.Credentials{AccessToken: token})
}

func (h *handlerImpl) signToken(user domain.UserDTO, duration time.Duration) (string, error) {
	expiresAt := time.Now().Add(time.Minute * time.Duration(duration)).Unix()
	// Set custom claims
	claims := &domain.Claims{
		ID:       user.ID,
		Username: user.Username,
		Name:     user.Name,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expiresAt,
		},
	}
	secret := os.Getenv("SECRET")
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	return t, nil
}

// Register godoc
// @Summary      Register new user
// @Description  API for registering new user
// @Tags         users
// @Accept       json
// @Produce      json
// @Param    		 request  body      domain.RegisterBody  true  "Register body"
// @Success      200	{object}  domain.Credentials 	"success"
// @Failure      400  {string}  string 							"invalid body"
// @Failure      500  {string}  string 							"internal server error"
// @Router       /users/register [post]
func (h *handlerImpl) Register(c echo.Context) error {
	c.Logger().Info("[UserHandler.Register] register is starting")
	ctx := c.Request().Context()
	body := domain.RegisterBody{}
	if err := c.Bind(&body); err != nil {
		c.Logger().Error("[UserHandler.Register] invalid body", err)
		return c.JSON(apperrors.GetStatusCode(apperrors.ErrInvalidBody), apperrors.ErrInvalidBody.Error())
	}
	user, err := h.usecase.Register(ctx, body)
	if err != nil {
		c.Logger().Error("[UserHandler.Register] register failed", err)
		return c.JSON(apperrors.GetStatusCode(err), err.Error())
	}
	c.Logger().Info("[UserHandler.Register] register success")
	token, err := h.signToken(user, 300)
	if err != nil {
		c.Logger().Error("[UserHandler.Register] sign token error", err)
		return c.JSON(apperrors.GetStatusCode(apperrors.ErrInternalServerError), apperrors.ErrInternalServerError.Error())
	}
	return c.JSON(http.StatusCreated, domain.Credentials{AccessToken: token})
}
