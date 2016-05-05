package auth

import (
	"github.com/labstack/echo"
	userManager "../../models/user_manager"
)

func Auth() echo.MiddlewareFunc {
	return BasicAuth(func(username, password string) bool {
		userManager := userManager.NewUserManager()
		return userManager.ValidateUser(username, password)
	}, func(c echo.Context, username string) {
		c.Set("username", username)
	})
}
