package plugin

import (
 	"github.com/labstack/echo"
   "../../engine"
   "../user@hackm/service"
)


func Init(g *echo.Group) error {
  engine.SetAuthorize(func() echo.MiddlewareFunc {
    return BasicAuth(func(username, password string) bool {
      userManager := service.NewUserManager()
      return userManager.ValidateUser(username, password)
    }, func(c echo.Context, username string) {
      c.Set("username", username)
    })
  })
  return nil
}