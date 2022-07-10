package middleware

import (
	"github.com/labstack/echo/v4"
)

var (
	UserKey = "user"
)

type CustomContext struct {
	echo.Context
	UserID int64
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
