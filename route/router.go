package route

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"net/http"
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
	v1 := e.Group("/api/v1")
	{
		// CRUD
		v1.Post("/papers", func(c echo.Context) error {
			return c.String(http.StatusOK, "papers")
		})
	}
	return e
}
