package route

import (
  "github.com/labstack/echo"
  "html/template"
  "io"
)

type Renderer struct {
    templates *template.Template
}

func (t *Renderer) Render(w io.Writer, name string, data interface{}, e echo.Context) error {
    return t.templates.ExecuteTemplate(w, name, data)
}

func CreateRenderer() *Renderer {
  return  &Renderer {
    templates: template.Must(template.ParseFiles(
      "templates/header.tmpl",
      "templates/footer.tmpl",      
      "templates/paper.tmpl",      
      "templates/papers.tmpl",
      "templates/tag.tmpl",
      "templates/top.tmpl",
      "templates/search.tmpl")),
  }
}