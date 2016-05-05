package top

import (
  "github.com/labstack/echo"
  "net/http"
  "../base"
)

func GetTop(c echo.Context) error {
  return c.Render(http.StatusOK, "top", TopVM {
    Meta: base.Meta {
      Title: "Top",
      Description: "This is top page.",
    },
    Theme: base.Theme {
      Name: "simple",
    },
  })
}