package paper

import (
  "github.com/labstack/echo"
  "net/http"
  "../base"
)

func GetPaper(c echo.Context) error {
  return c.Render(http.StatusOK, "paper", PaperVM {
    Meta: base.Meta {
      Title: "Paper",
      Description: "This is paper page.",
    },
    Theme: base.Theme {
      Name: "simple",
    },
    Id: c.Param("id"),
  })
}
func GetPapers(c echo.Context) error {
  return c.Render(http.StatusOK, "papers", PapersVM {
    Meta: base.Meta {
      Title: "Papers",
      Description: "This is paper list page.",
    },
    Theme: base.Theme {
      Name: "simple",
    },
    Username: c.Param("username"),
  })
}