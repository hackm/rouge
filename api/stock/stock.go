package stock

import (
	"github.com/labstack/echo"
	"net/http"
)

//GetStocks retrieve stock by User
func GetStocks(c echo.Context) error {
	return c.String(http.StatusOK, "get stocks")
}

//CreateStock create a new stock for paper
func CreateStock(c echo.Context) error {
	return c.String(http.StatusOK, "new stock for paper(" + c.Param("paper_id") + ") is created")
}

//DeleteStock delete stock by user and paper
func DeleteStock(c echo.Context) error {
	return c.String(http.StatusOK, "stock for paper(" + c.Param("paper_id") + ") is deleted")
}
