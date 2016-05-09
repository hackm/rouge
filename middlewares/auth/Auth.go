package auth

import (
	"github.com/labstack/echo"
	"../../models/services"
)

func Auth() echo.MiddlewareFunc {
	return BasicAuth(func(username, password string) bool {
		userManager := services.NewUserManager()
		return userManager.ValidateUser(username, password)
	}, func(c echo.Context, username string) {
		c.Set("username", username)
	})
}
