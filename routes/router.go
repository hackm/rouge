package route

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	apiContent "../controllers/api/content"
	apiPaper "../controllers/api/paper"
	apiTag "../controllers/api/tag"
	apiStock "../controllers/api/stock"
	apiUser "../controllers/api/user"
	"../controllers/top"
	"../controllers/paper"
	"../controllers/tag"
	"../controllers/search"
	//"../middlewares/auth"
)

func Init() *echo.Echo {
	e := echo.New()
	// Debug
	e.Debug()
	// Setup Middleware
	e.Use(middleware.Logger())  // Log HTTP requests
	e.Use(middleware.Recover()) // Recover from panics
	e.Use(middleware.Gzip())    // Send gzip HTTP response
	// e.Use(auth.Auth())
	// API Version name
	api := e.Group("/api")
	{
		// User
		api.Post("users", apiUser.CreateUser)

		//User Papers
		api.Get("/papers", apiPaper.GetUserPapers)
		api.Get("/papers/:id", apiPaper.GetPaper)
		api.Post("/papers", apiPaper.CreatePaper)
		api.Put("/papers/:id", apiPaper.UpdatePaper)
		api.Delete("/papers/:id", apiPaper.DeletePaper)

		//User Paper Conntents
		api.Get("/papers/:id/contents", apiContent.GetContents)

		//User Paper Tags
		api.Post("/papers/:id/tags", apiPaper.AddPaperTag)
		api.Delete("/papers/:id/tags", apiPaper.DeletePaperTag)

		// User Stocks
		api.Get("/stocks", apiStock.GetStocks)
		api.Post("/stocks/:paper_id", apiStock.CreateStock)
		api.Delete("/stocks/:paper_id", apiStock.DeleteStock)

		// Tag
		api.Get("/tags", apiTag.GetTags)
		api.Get("/tags/:id", apiTag.GetTag)
		api.Post("/tags", apiTag.CreateTag)
		api.Put("/tags/:id", apiTag.UpdateTag)
		api.Delete("/tags/:id", apiTag.DeleteTag)
	}

	e.SetRenderer(CreateRenderer())
	e.Get("/", top.GetTop)
	e.Get("/:username/papers", paper.GetPapers)
	e.Get("/:username/papers/:id", paper.GetPaper)
	e.Get("/tags/:name", tag.GetTag)
	e.Get("/search", search.GetSearch)

	e.Static("/statics", "statics")
	return e
}
