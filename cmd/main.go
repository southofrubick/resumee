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
		templates: template.Must(template.ParseGlob("views/*.html")),
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

	e.GET("/contact-info", func(c echo.Context) error {
		return c.Render(200, "contact-info", 0)
	})

	e.GET("/summary", func(c echo.Context) error {
		return c.Render(200, "summary", 0)
	})

	e.GET("/education", func(c echo.Context) error {
		return c.Render(200, "education", 0)
	})

	e.GET("/work-experience", func(c echo.Context) error {
		return c.Render(200, "work-experience", 0)
	})

	e.GET("/skills", func(c echo.Context) error {
		return c.Render(200, "skills", 0)
	})

	e.GET("/projects", func(c echo.Context) error {
		return c.Render(200, "projects", 0)
	})

	e.Logger.Fatal(e.Start(":42069"))
}
