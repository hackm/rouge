package route

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"net/http"
	"../model/vm"
)

func Init() *echo.Echo {
	e := echo.New()
	// Debug
	e.Debug()
	// Setup Middleware
	e.Use(middleware.Logger()) // Log HTTP requests
	e.Use(middleware.Recover()) // Recover from panics
	e.Use(middleware.Gzip()) // Send gzip HTTP response
	// API Version name
	api := e.Group("/api")
	{
		//User Papers
		api.Get("/:username/papers", func (c echo.Context) error {
			return c.String(http.StatusOK, c.Param("username") + "'s papers")
		})
		api.Get("/:username/papers/:id", func (c echo.Context) error {
			return c.String(http.StatusOK, c.Param("username") + "'s paper(" + c.Param("id") + ")")
		})
		api.Post("/:username/papers", func(c echo.Context) error {
			return c.String(http.StatusCreated, c.Param("username") + "'s new paper")
		})
		api.Put("/:username/papers/:id", func(c echo.Context) error {
			return c.String(http.StatusOK, c.Param("username") + "'s paper(" + c.Param("id") + ") is updated!")
		})
		api.Delete("/:username/papers/:id", func(c echo.Context) error {
			return c.String(http.StatusOK, c.Param("username") + "'s paper(" + c.Param("id") + ") is Deleted!")
		})
		
		//User Paper Conntents		
		api.Get("/:username/papers/:id/contents", func(c echo.Context) error {
			return c.String(http.StatusOK, "add tag to " + c.Param("username") + "'s paper(" + c.Param("id") + ")")
		})
		
		//User Paper Tags
		api.Post("/:username/papers/:id/tags", func(c echo.Context) error {
			return c.String(http.StatusOK, "add tag to " + c.Param("username") + "'s paper(" + c.Param("id") + ")")
		})
		api.Delete("/:username/papers/:id/tags", func(c echo.Context) error {
			return c.String(http.StatusOK, "delete tag to " + c.Param("username") + "'s paper(" + c.Param("id") + ")")
		})
		
		// User Stocks
		api.Get("/:username/stocks", func (c echo.Context) error {
			return c.String(http.StatusOK, c.Param("username") + "'s stocks")
		})
		api.Delete("/:username/stocks/:paper_id", func(c echo.Context) error {
			return c.String(http.StatusOK, c.Param("username") + "'s stock for paper(" + c.Param("paper_id") + ") is deleted")
		})
		
		
		// Tag List
		api.Get("/tags", func (c echo.Context) error {
			return c.String(http.StatusOK, "tag list")
		})
		api.Get("/tags/:id", func (c echo.Context) error {
			return c.String(http.StatusOK, "tag(" + c.Param("id") + ")")
		})
	}
	
	e.SetRenderer(CreateRenderer())
	e.Get("/:username/papers", func (c echo.Context) error{
		return c.Render(http.StatusOK, "papers", vm.Meta{ 
			Title: "Papers", 
			Description: "This is paper list page."})
	})
	e.Get("/:username/papers/:id", func (c echo.Context) error{
		return c.Render(http.StatusOK, "paper", vm.Meta{ 
			Title: "Paper", 
			Description: "This is paper page."})
	})
	e.Get("/tags", func (c echo.Context) error{
		return c.Render(http.StatusOK, "tags", vm.Meta{ 
			Title: "Tags", 
			Description: "This is tag list page."})
	})
	e.Get("/tags/:id", func (c echo.Context) error{
		return c.Render(http.StatusOK, "tag", vm.Meta{ 
			Title: "Tag", 
			Description: "This is tag page."})
	})
	
	e.Static("/static", "static")
	return e
}
