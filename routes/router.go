package route

import (
	"net/http"

	"../controllers/api/content"
	"../controllers/api/paper"
	"../controllers/api/stock"
	"../controllers/api/tag"
	"../models/vm"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func Init() *echo.Echo {
	e := echo.New()
	// Debug
	e.Debug()
	// Setup Middleware
	e.Use(middleware.Logger())  // Log HTTP requests
	e.Use(middleware.Recover()) // Recover from panics
	e.Use(middleware.Gzip())    // Send gzip HTTP response
	// API Version name
	api := e.Group("/api")
	{
		//User Papers
		api.Get("/papers", paper.GetUserPapers)
		api.Get("/papers/:id", paper.GetPaper)
		api.Post("/papers", paper.CreatePaper)
		api.Put("/papers/:id", paper.UpdatePaper)
		api.Delete("/papers/:id", paper.DeletePaper)

		//User Paper Conntents
		api.Get("/papers/:id/contents", content.GetContents)

		//User Paper Tags
		api.Post("/papers/:id/tags", paper.AddPaperTag)
		api.Delete("/papers/:id/tags", paper.DeletePaperTag)

		// User Stocks
		api.Get("/stocks", stock.GetStocks)
		api.Post("/stocks/:paper_id", stock.CreateStock)
		api.Delete("/stocks/:paper_id", stock.DeleteStock)

		// Tag
		api.Get("/tags", tag.GetTags)
		api.Get("/tags/:id", tag.GetTag)
		api.Post("/tags", tag.CreateTag)
		api.Put("/tags/:id", tag.UpdateTag)
		api.Delete("/tags/:id", tag.DeleteTag)
	}

	e.SetRenderer(CreateRenderer())
	e.Get("/", func(c echo.Context) error {
		return c.Render(http.StatusOK, "top", vm.ViewData{
			Meta: vm.Meta{
				Title:       "Top",
				Description: "This is top page.",
			},
			Theme: "simple",
		})
	})
	e.Get("/:username/papers", func(c echo.Context) error {
		return c.Render(http.StatusOK, "papers", vm.ViewData{
			Meta: vm.Meta{
				Title:       "Papers",
				Description: "This is paper list page.",
			},
			Theme: "simple",
		})
	})
	e.Get("/:username/papers/:id", func(c echo.Context) error {
		return c.Render(http.StatusOK, "paper", vm.ViewData{
			Meta: vm.Meta{
				Title:       "Paper",
				Description: "This is paper page.",
			},
			Theme: "simple",
		})
	})
	e.Get("/tags/:id", func(c echo.Context) error {
		return c.Render(http.StatusOK, "tag", vm.ViewData{
			Meta: vm.Meta{
				Title:       "Tag",
				Description: "This is tag page.",
			},
			Theme: "simple",
		})
	})
	e.Get("/search", func(c echo.Context) error {
		return c.Render(http.StatusOK, "search", vm.ViewData{
			Meta: vm.Meta{
				Title:       "Search",
				Description: "This is search page.",
			},
			Theme: "simple",
		})
	})

	e.Static("/statics", "statics")
	return e
}
