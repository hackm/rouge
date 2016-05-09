package plugin

import (
 	"github.com/labstack/echo"
   "./api"
   "./service"
   "./controller"
   "../../engine"
)

const PLUGIN_NAME string = "paper@hackm"

func PaperContextFactory(next echo.HandlerFunc) echo.HandlerFunc {
  return func(c echo.Context) error {
    c.Set(service.Key(), service.NewService())
    return next(c)
  }
}

func Init(g *echo.Group) error {
  g.Use(PaperContextFactory)
  a := g.Group("/api") 
  {
    a.Get("/papers", api.GetPapers)
    a.Get("/papers/:id", api.GetPaper)
    a.Get("/papers/:id/olders", api.GetPaperContents, engine.Authorize())
    a.Post("/papers", api.PostPapers, engine.Authorize())
    a.Put("/papers/:id", api.PutPaper, engine.Authorize())
    a.Delete("/papers/:id", api.DeletePaper, engine.Authorize())    
  }
  engine.AddTemplate(PLUGIN_NAME, "template/footer.tmpl")
  engine.AddTemplate(PLUGIN_NAME, "template/header.tmpl")
  engine.AddTemplate(PLUGIN_NAME, "template/paper.tmpl")
  engine.AddTemplate(PLUGIN_NAME, "template/top.tmpl")
  g.Get("/", controller.GetPapers)
  g.Get("/papers/:id", controller.GetPaper)
  return nil
}