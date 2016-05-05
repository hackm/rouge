package content

import (
	"github.com/labstack/echo"
	"net/http"
)

//GetPaperContents retrieve paper's contents list
func GetContents(c echo.Context) error {
	return c.String(http.StatusOK, "paper(" + c.Param("id") + ") contents")
}
