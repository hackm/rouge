package controller

import (
 	"github.com/labstack/echo"
	"net/http"
  "../model"
  "../../../engine"
)

func GetPapers(c echo.Context) error {
  return c.Render(http.StatusOK, "top:paper@hackm", model.Top {
    Meta: engine.Meta {
      Title: "Top",
      Description: "Rouge Paper",
    },
    Theme: engine.Theme {
      Name: "simple",
    },
  })
}
func GetPaper(c echo.Context) error {
  return c.Render(http.StatusOK, "paper:paper@hackm", model.Paper {
    Meta: engine.Meta {
      Title: "Top",
      Description: "Rouge Paper",
    },
    Theme: engine.Theme {
      Name: "simple",
    },
    Id: c.Param("id"),
  })
}