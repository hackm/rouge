package engine

import (
  "github.com/labstack/echo"
  "html/template"
  "io"
)

type Renderer struct {
    Templates *template.Template
}

func (t *Renderer) Render(w io.Writer, name string, data interface{}, e echo.Context) error {
    return t.Templates.ExecuteTemplate(w, name, data)
}

var templates []string
func AddTemplate(path string) {
  templates = append(templates, path)
}
func AddPluginTemplate(plugin, relationPath string) {
  templates = append(templates, "plugins/" + plugin + "/" + relationPath)
}

func NewRenderer() *Renderer {
  return  &Renderer {
    Templates: template.Must(template.ParseFiles(templates...)),
  }
}

/*
 "templates/header.tmpl",
      "templates/footer.tmpl",      
      "templates/paper.tmpl",      
      "templates/papers.tmpl",
      "templates/tag.tmpl",
      "templates/top.tmpl",
      "templates/search.tmpl"*/