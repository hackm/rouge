package engine

import (
  "github.com/labstack/echo"
)

var authorize func() echo.MiddlewareFunc = func() echo.MiddlewareFunc {
  return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			return next(c)
		}
  }
}

func SetAuthorize(fn func() echo.MiddlewareFunc) {
  authorize = fn
}
func Authorize() echo.MiddlewareFunc {
  return authorize()
}