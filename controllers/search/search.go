package search

import (
  "github.com/labstack/echo"
  "net/http"
  "../base"
)

func GetSearch(c echo.Context) error {
  return c.Render(http.StatusOK, "search", SearchVM {
    Meta: base.Meta {
      Title: "Search",
      Description: "This is search result page.",
    },
    Theme: base.Theme {
      Name: "simple",
    },
    Keyword: c.QueryParam("keyword"),
    Tag: c.QueryParam("tag"),
  })
}