package main

import (
	"html/template"
	"io"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Templates struct {
	templates *template.Template
}

func (t *Templates) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func NewTemplates() *Templates {
	return &Templates{
		templates: template.Must(template.ParseGlob("views/index.html")),
	}
}

type Count struct {
	Count int
}

func main() {

	e := echo.New()
	e.Use(middleware.Logger())
	e.File("css/style.css", "css/style.css")

	e.Renderer = NewTemplates()

	e.GET("/",  func(c echo.Context) error {
		return c.Render(200, "index", 0)
	})

	e.GET("/left",  func(c echo.Context) error {
		return c.Render(200, "left", 0)
	})

	e.GET("/center",  func(c echo.Context) error {
		return c.Render(200, "center", 0)
	})

	e.GET("/right",  func(c echo.Context) error {
		return c.Render(200, "right", 0)
	})

	e.Logger.Fatal(e.Start(":42069"))
}
