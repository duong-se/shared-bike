package middleware

import (
	"shared-bike/apperrors"

	"github.com/google/uuid"
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

func AddHeaderXRequestID(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		reqId := c.Request().Header.Get(echo.HeaderXRequestID)
		if len(reqId) == 0 {
			reqId = uuid.NewString()
			c.Request().Header.Set(echo.HeaderXRequestID, reqId)
		}
		c.Logger().SetPrefix(reqId)
		c.Response().Header().Set(echo.HeaderXRequestID, reqId)
		return next(c)
	}
}
