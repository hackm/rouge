package user

import (
	"github.com/labstack/echo"
	"net/http"
)

func CreateUser(c echo.Context) error {
	return c.String(http.StatusOK, "create user")
}
