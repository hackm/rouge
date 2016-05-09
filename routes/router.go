package route

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"../engine"
  
	plugin0 "../plugins/user@hackm"
	plugin1 "../plugins/basic-auth@hackm"
	plugin2 "../plugins/paper@hackm"
)

func Init() *echo.Echo {
	e := echo.New()
	// Debug
	e.Debug()
	// Setup Middleware
	e.Use(middleware.Logger())  // Log HTTP requests
	e.Use(middleware.Recover()) // Recover from panics
	e.Use(middleware.Gzip())    // Send gzip HTTP response
	
  
  p0 := e.Group("/user@hackm")
  plugin0.Init(p0)
  
  p1 := e.Group("/basic-auth@hackm")
  plugin1.Init(p1)
  
  p2 := e.Group("/paper@hackm")
  plugin2.Init(p2)
  
  
  engine.AddTemplate("templates/footer.tmpl")
  engine.AddTemplate("templates/header.tmpl")
	e.SetRenderer(engine.NewRenderer())	
	e.Static("/statics", "statics")
	
	return e
}
