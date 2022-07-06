package user

import (
	"net/http"

	"shared-bike/domain"
	"shared-bike/middleware"

	"shared-bike/apperrors"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
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
// @Description  Login by username and password
// @Tags         users
// @Accept       json
// @Produce      json
// @Param    		 request  body      domain.LoginPayload  true  "Login payload"
// @Success      204
// @Failure      400  {string}  string 	"invalid payload"
// @Failure      404  {string}  string 	"username or password is wrong"
// @Failure      500  {string}  string 	"internal server error"
// @Router       /users/login [post]
func (h *handlerImpl) Login(c echo.Context) error {
	ctx := c.Request().Context()
	payload := domain.LoginPayload{}
	if err := c.Bind(&payload); err != nil {
		c.Logger().Error("[UserHandler.Login] invalid payload", err)
		return c.JSON(http.StatusBadRequest, "invalid payload")
	}
	c.Logger().Info("[UserHandler.Login] logging")
	user, err := h.usecase.Login(ctx, payload)
	if err != nil {
		c.Logger().Error("[UserHandler.Login] login failed", err)
		return c.JSON(apperrors.GetStatusCode(err), err.Error())
	}
	c.Logger().Info("[UserHandler.Login] login success")
	h.setSession(user, c)
	return c.JSON(http.StatusNoContent, nil)
}

func (h *handlerImpl) setSession(user domain.User, c echo.Context) error {
	sess, err := session.Get("session", c)
	if err != nil {
		return err
	}
	day := 86400
	sess.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   day * 7,
		HttpOnly: true,
		Secure:   false,
	}
	sess.Values[middleware.UserIDKey] = user.ID
	sess.Save(c.Request(), c.Response())
	return nil
}

// Register godoc
// @Summary      Register
// @Description  Register by username and password
// @Tags         users
// @Accept       json
// @Produce      json
// @Param    		 request  body      domain.RegisterPayload  true  "Register payload"
// @Success      204
// @Failure      400  {string}  string 	"invalid payload"
// @Failure      500  {string}  string 	"internal server error"
// @Router       /users/register [post]
func (h *handlerImpl) Register(c echo.Context) error {
	c.Logger().Info("[UserHandler.Register] register is starting")
	ctx := c.Request().Context()
	payload := domain.RegisterPayload{}
	if err := c.Bind(&payload); err != nil {
		c.Logger().Error("[UserHandler.Register] invalid payload", err)
		return c.JSON(http.StatusBadRequest, "invalid payload")
	}
	user, err := h.usecase.Register(ctx, payload)
	if err != nil {
		c.Logger().Error("[UserHandler.Register] register failed", err)
		return c.JSON(apperrors.GetStatusCode(err), err.Error())
	}
	c.Logger().Info("[UserHandler.Register] register success")
	h.setSession(user, c)
	return c.JSON(http.StatusCreated, nil)
}
