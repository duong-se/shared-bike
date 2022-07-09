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
// @Description  API for logining
// @Tags         users
// @Accept       json
// @Produce      json
// @Param    		 request  body      domain.LoginBody  true  "Login body"
// @Success      204
// @Failure      400  {string}  string 	"invalid body"
// @Failure      404  {string}  string 	"username or password is wrong"
// @Failure      500  {string}  string 	"internal server error"
// @Router       /users/login [post]
func (h *handlerImpl) Login(c echo.Context) error {
	ctx := c.Request().Context()
	body := domain.LoginBody{}
	if err := c.Bind(&body); err != nil {
		c.Logger().Error("[UserHandler.Login] invalid body", err)
		return c.JSON(http.StatusBadRequest, "invalid body")
	}
	c.Logger().Info("[UserHandler.Login] logging")
	user, err := h.usecase.Login(ctx, body)
	if err != nil {
		c.Logger().Error("[UserHandler.Login] login failed", err)
		return c.JSON(apperrors.GetStatusCode(err), err.Error())
	}
	c.Logger().Info("[UserHandler.Login] login success")
	err = h.setSession(user, c)
	if err != nil {
		c.Logger().Error("[UserHandler.Login] set session error", err)
		return c.JSON(http.StatusInternalServerError, "internal server error")
	}
	return c.JSON(http.StatusNoContent, nil)
}

func (h *handlerImpl) setSession(user domain.UserDTO, c echo.Context) error {
	sess, err := session.Get("session", c)
	if err != nil {
		return err
	}
	day := 86400
	sess.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   day * 7,
		HttpOnly: false,
		Secure:   false,
	}
	sess.Values[middleware.UserIDKey] = user.ID
	sess.Save(c.Request(), c.Response())
	return nil
}

// Register godoc
// @Summary      Register new user
// @Description  API for registering new user
// @Tags         users
// @Accept       json
// @Produce      json
// @Param    		 request  body      domain.RegisterBody  true  "Register body"
// @Success      204
// @Failure      400  {string}  string 	"invalid body"
// @Failure      500  {string}  string 	"internal server error"
// @Router       /users/register [post]
func (h *handlerImpl) Register(c echo.Context) error {
	c.Logger().Info("[UserHandler.Register] register is starting")
	ctx := c.Request().Context()
	body := domain.RegisterBody{}
	if err := c.Bind(&body); err != nil {
		c.Logger().Error("[UserHandler.Register] invalid body", err)
		return c.JSON(http.StatusBadRequest, "invalid body")
	}
	user, err := h.usecase.Register(ctx, body)
	if err != nil {
		c.Logger().Error("[UserHandler.Register] register failed", err)
		return c.JSON(apperrors.GetStatusCode(err), err.Error())
	}
	c.Logger().Info("[UserHandler.Register] register success")
	err = h.setSession(user, c)
	if err != nil {
		c.Logger().Error("[UserHandler.Register] set session error", err)
		return c.JSON(http.StatusInternalServerError, "internal server error")
	}
	return c.JSON(http.StatusCreated, nil)
}
