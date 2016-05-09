package controller

import (
 	"github.com/labstack/echo"
	"net/http"
  "../model"
)

func GetPapers(c echo.Context) error {
  return c.Render(http.StatusOK, "top:paper@hackm", model.Top {
    Meta: model.Meta {
      Title: "Top",
      Description: "Rouge Paper",
    },
    Theme: model.Theme {
      Name: "simple",
    },
  })
}
func GetPaper(c echo.Context) error {
  return c.Render(http.StatusOK, "paper:paper@hackm", model.Paper {
    Meta: model.Meta {
      Title: "Top",
      Description: "Rouge Paper",
    },
    Theme: model.Theme {
      Name: "simple",
    },
    Id: c.Param(":id"),
  })
}