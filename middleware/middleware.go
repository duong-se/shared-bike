package middleware

import (
	"shared-bike/apperrors"

	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

var (
	UserIDKey = "userID"
)

type CustomContext struct {
	echo.Context
	UserID int64
}

func Authorize(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		unauthorizedErr := apperrors.ErrUnauthorizeError
		sess, err := session.Get("session", c)
		if err != nil {
			c.Logger().Error("[Middlware.Authorize] error on get session", err)
			return c.JSON(apperrors.GetStatusCode(unauthorizedErr), unauthorizedErr.Error())
		}
		if sess.Values[UserIDKey] != nil {
			c.Logger().Info("[Middlware.Authorize] authorized request")
			c.Set(UserIDKey, sess.Values[UserIDKey])
			return next(c)
		}
		c.Logger().Info("[Middlware.Authorize] unauthorized request")
		return c.JSON(apperrors.GetStatusCode(unauthorizedErr), unauthorizedErr.Error())
	}
}

func AddLoggerContext(contextLogger CustomLogger) func(next echo.HandlerFunc) echo.HandlerFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			req := c.Request()
			res := c.Response()
			id := req.Header.Get(echo.HeaderXRequestID)
			if id == "" {
				id = res.Header().Get(echo.HeaderXRequestID)
			}
			contextLogger.SetRequestID(id)
			c.SetLogger(contextLogger)
			return next(c)
		}
	}
}
