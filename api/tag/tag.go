package tag

import (
	"github.com/labstack/echo"
	"net/http"
)

//GetTags retrieve all tags
func GetTags(c echo.Context) error {
	return c.String(http.StatusOK, "tag list")
}

//GetTag retrieve a tag by ID
func GetTag(c echo.Context) error {
	return c.String(http.StatusOK, "tag(" + c.Param("id") + ")")
}

//CreateTag create a new tag for name
func CreateTag(c echo.Context) error {
	return c.String(http.StatusOK, "tag(" + c.Param("name") + ") created")
}

//UpdateTag update a tag by ID
func UpdateTag(c echo.Context) error {
	return c.String(http.StatusOK, "tag(" + c.Param("id") + ") updated")
}

//DeleteTag delete a tag by ID
func DeleteTag(c echo.Context) error {
	return c.String(http.StatusOK, "tag(" + c.Param("id") + ") deleted")
}
