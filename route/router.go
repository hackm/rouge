package route

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"net/http"
	"../model/vm"
	"../api/content"
	"../api/paper"
	"../api/stock"
	"../api/tag"
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
		api.Get("/:username/papers", paper.GetUserPapers)
		api.Get("/:username/papers/:id", paper.GetPaper)
		api.Post("/:username/papers", paper.CreatePaper)
		api.Put("/:username/papers/:id", paper.UpdatePaper)
		api.Delete("/:username/papers/:id", paper.DeletePaper)

		//User Paper Conntents
		api.Get("/:username/papers/:id/contents", content.GetContents)

		//User Paper Tags
		api.Post("/:username/papers/:id/tags", paper.AddPaperTag)
		api.Delete("/:username/papers/:id/tags", paper.DeletePaperTag)

		// User Stocks
		api.Get("/:username/stocks", stock.GetStocks)
		api.Post("/:username/stocks/:paper_id", stock.CreateStock)
		api.Delete("/:username/stocks/:paper_id", stock.DeleteStock)

		// Tag
		api.Get("/tags", tag.GetTags)
		api.Get("/tags/:id", tag.GetTag)
		api.Post("/tags", tag.CreateTag)
		api.Put("/tags/:id", tag.UpdateTag)
		api.Delete("/tags/:id", tag.DeleteTag)
	}

	e.SetRenderer(CreateRenderer())
	e.Get("/:username/papers", func (c echo.Context) error{
		return c.Render(http.StatusOK, "papers", vm.ViewData{
			Meta: vm.Meta{
				Title: "Papers",
				Description: "This is paper list page.",
			},
			Theme: "simple"
		})
	})
	e.Get("/:username/papers/:id", func (c echo.Context) error{
		return c.Render(http.StatusOK, "paper", vm.ViewData{
			Meta: vm.Meta{
				Title: "Paper",
				Description: "This is paper page."
			},
			Theme: "simple",
		})
	})
	e.Get("/tags", func (c echo.Context) error{
		return c.Render(http.StatusOK, "tags", vm.ViewData{
			Meta: vm.Meta{
				Title: "Tags",
				Description: "This is tag list page."
			},
			Theme: "simple",
		})
	})
	e.Get("/tags/:id", func (c echo.Context) error{
		return c.Render(http.StatusOK, "tag", vm.ViewData{
			Meta: vm.Meta{
				Title: "Tag",
				Description: "This is tag page."
			},
			Theme: "simple",
		})
	})

	e.Static("/static", "static")
	return e
}
