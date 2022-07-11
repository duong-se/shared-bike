package middleware

import (
	"regexp"
	"shared-bike/apperrors"

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

func WhiteListAPI(c echo.Context) bool {
	requestPath := c.Request().URL.Path
	c.Logger().Debug("request ========>", requestPath)
	return requestPath == "/api/v1/users/login" || requestPath == "/api/v1/users/register" || requestPath == "/health" || regexp.MustCompile(`\/swagger\/[a-zA-Z0-9]+.[a-zA-Z0-9]+`).MatchString(requestPath)
}

func CustomJWTError(err error, c echo.Context) error {
	c.Logger().Error("[JWTValidate] error", err)
	return c.JSON(apperrors.GetStatusCode(apperrors.ErrUnauthorizeError), apperrors.ErrUnauthorizeError.Error())
}
