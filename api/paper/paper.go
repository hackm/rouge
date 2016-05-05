package paper

import ()
import (
	"github.com/labstack/echo"
	"net/http"
)

//GetUserPapers retrieves papers by username
func GetUserPapers(c echo.Context) error {
	return c.String(http.StatusOK, c.Param("username") + "'s papers")
}

//GetUserPaper retrieves papers by username & ID
func GetPaper(c echo.Context) error {
	return c.String(http.StatusOK, c.Param("username") + "'s paper(" + c.Param("id") + ")")
}

//CreateUserPaper creates a new paper and content record
func CreatePaper(c echo.Context) error {
	return c.String(http.StatusCreated, c.Param("username") + "'s new paper")
}

//UpdateUserPaper update a paper for username & ID
func UpdatePaper(c echo.Context) error {
	return c.String(http.StatusOK, c.Param("username") + "'s paper(" + c.Param("id") + ") is updated!")
}

//DeleteUserPaper delete a paper for username & ID
func DeletePaper(c echo.Context) error {
	return c.String(http.StatusOK, c.Param("username") + "'s paper(" + c.Param("id") + ") is Deleted!")
}

//GetPaperConntents retrieve paper's contents list
func GetPaperContents(c echo.Context) error {
	return c.String(http.StatusOK, "add tag to " + c.Param("username") + "'s paper(" + c.Param("id") + ")")
}

//AddPaperTag add new tag for paper
func AddPaperTag(c echo.Context) error {
	return c.String(http.StatusOK, "add tag to " + c.Param("username") + "'s paper(" + c.Param("id") + ")")
}

//DeletePaperTag delete a tag by paper id
func DeletePaperTag(c echo.Context) error {
	return c.String(http.StatusOK, "delete tag to " + c.Param("username") + "'s paper(" + c.Param("id") + ")")
}
