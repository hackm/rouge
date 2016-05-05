package tag

import (
  "github.com/labstack/echo"
  "net/http"
  "../base"
)

func GetTag(c echo.Context) error {
  return c.Render(http.StatusOK, "tag", TagVM {
    Meta: base.Meta {
      Title: "Tag",
      Description: "This is tag page.",
    },
    Theme: base.Theme {
      Name: "simple",
    },
    Tag: c.Param("name"),
  })
}