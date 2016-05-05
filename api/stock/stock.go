package stock

import (
	"github.com/labstack/echo"
	"net/http"
)

//GetStocks retrieve stock by User
func GetStocks(c echo.Context) error {
	return c.String(http.StatusOK, c.Param("username") + "'s stocks")
}

//CreateStock create a new stock for paper
func CreateStock(c echo.Context) error {
	return c.String(http.StatusOK, c.Param("username")  + " new stock for paper(" + c.Param("paper_id") + ") is created")
}

//DeleteStock delete stock by user and paper
func DeleteStock(c echo.Context) error {
	return c.String(http.StatusOK, c.Param("username") + "'s stock for paper(" + c.Param("paper_id") + ") is deleted")
}
