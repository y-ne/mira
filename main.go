package main

import (
	"io"
	"text/template"

	"github.com/labstack/echo/v4"
	"github.com/y-ne/mira/routes"
)

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func main() {
	t := &Template{
		templates: template.Must(template.ParseGlob("public/views/*.html")),
	}

	e := echo.New()
	e.Renderer = t

	routes.Router(e)

	e.Logger.Fatal(e.Start(":8123"))
}
