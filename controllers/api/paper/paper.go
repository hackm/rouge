package paper

import ()
import (
	"github.com/labstack/echo"
	"net/http"
	"strconv"
	"../../../models/paper_service"
)

//GetUserPapers retrieves papers
func GetUserPapers(c echo.Context) error {
	username, include := c.QueryParams()["username"]
	if include {
		return c.String(http.StatusOK, username[0] + "'s papers")
	} else {
		return c.String(http.StatusOK, "all papers")
	}
}

//GetUserPaper retrieves papers by username & ID
func GetPaper(c echo.Context) error {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	s := paper_service.NewPaperService()
	p := s.GetPaper(id)
	return c.JSON(http.StatusOK, p)
}

//CreateUserPaper creates a new paper and content record
func CreatePaper(c echo.Context) error {
	return c.String(http.StatusCreated, "create a new paper for " + c.Get("username").(string))
}

//UpdateUserPaper update a paper for username & ID
func UpdatePaper(c echo.Context) error {
	return c.String(http.StatusOK, "paper(" + c.Param("id") + ") is updated!")
}

//DeleteUserPaper delete a paper for username & ID
func DeletePaper(c echo.Context) error {
	return c.String(http.StatusOK, "paper(" + c.Param("id") + ") is Deleted!")
}

//AddPaperTag add new tag for paper
func AddPaperTag(c echo.Context) error {
	return c.String(http.StatusOK, "add tag to paper(" + c.Param("id") + ")")
}

//DeletePaperTag delete a tag by paper id
func DeletePaperTag(c echo.Context) error {
	return c.String(http.StatusOK, "delete tag to paper(" + c.Param("id") + ")")
}
