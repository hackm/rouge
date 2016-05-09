package plugin

import (
 	"github.com/labstack/echo"
   "../../engine"
)

const PLUGIN_NAME string = "user@hackm"

func Init(g *echo.Group) error {
  //a := g.Group("/api") 
  // {
  //   // a.Post("/users", api.CreateUser)   
  // }
  engine.AddPluginTemplate(PLUGIN_NAME, "template/register.tmpl")
  return nil
}