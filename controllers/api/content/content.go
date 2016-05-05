package content

import (
	"github.com/labstack/echo"
	"net/http"
	"strconv"
	"github.com/pkg/errors"
	"../../../models/repositories/content"
)

//GetPaperContents retrieve paper's contents list
func GetContents(c echo.Context) error {
	r := content.NewContentRepository()
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return errors.Wrap(err, "pager_id is not integer")
	}
	return c.JSON(http.StatusOK, r.GetContent(id))
}
