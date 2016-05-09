package api

import (
 	"github.com/labstack/echo"
	"net/http"
	"strconv"
	"github.com/pkg/errors"
  "../entity"
  "../service"
)

func GetPapers(c echo.Context) error {
  s := c.Get(service.Key()).(*service.Service)
  skip, errSkip := strconv.Atoi(c.QueryParam("skip"))
  take, errTake := strconv.Atoi(c.QueryParam("top"))
  if errSkip != nil {
    skip = 0
  }
  if errTake != nil {
    take = 10
  }
  return c.JSON(http.StatusOK, s.GetPapers(skip, take))
}

func GetPaper(c echo.Context) error {
  s := c.Get(service.Key()).(*service.Service)
  id, err4id := strconv.ParseInt(c.Param("id"), 10, 64)
  if err4id != nil {
		return errors.Wrap(err4id, "pager_id is not integer")
	}
  return c.JSON(http.StatusOK, s.GetPaper(id))
}

func GetPaperContents(c echo.Context) error {
  s := c.Get(service.Key()).(*service.Service)
  id, err4id := strconv.ParseInt(c.Param("id"), 10, 64)
  if err4id != nil {
		return errors.Wrap(err4id, "pager_id is not integer")
	}
  return c.JSON(http.StatusOK, s.GetPaperContents(id))
}

func PostPapers(c echo.Context) error {
  s := c.Get(service.Key()).(*service.Service)
  p := new(entity.Paper)
  if err4model := c.Bind(p); err4model != nil {
    return errors.Wrap(err4model, "invalid model")
  }
  return c.JSON(http.StatusOK, s.CreatePaper(p))
}

func PutPaper(c echo.Context) error {
  s := c.Get(service.Key()).(*service.Service)
  id, err4id := strconv.ParseInt(c.Param("id"), 10, 64)
  if err4id != nil {
		return errors.Wrap(err4id, "pager_id is not integer")
	}
  p := new(entity.Paper)
  if err4model := c.Bind(p); err4model != nil || p.Id != id {
    return errors.Wrap(err4model, "invalid model")
  }
  return c.JSON(http.StatusOK, s.UpdatePaper(p))
}

func DeletePaper(c echo.Context) error {
  s := c.Get(service.Key()).(*service.Service)
  id, err := strconv.ParseInt(c.Param("id"), 10, 64)
  if err != nil {
		return errors.Wrap(err, "pager_id is not integer")
	}
  return c.JSON(http.StatusOK, s.DeletePaper(id))
}