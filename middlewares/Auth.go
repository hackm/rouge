package middlewares

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func BasicAuth() echo.MiddlewareFunc {
	return middleware.BasicAuth(func(username, password string) bool {
		// TODO getting user information from username
		return username == "test" && password == "hogehoge"
	})
}
