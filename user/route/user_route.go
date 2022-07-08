package route

import (
	"gomq/user/controller"
	"html/template"
	"io"

	"github.com/labstack/echo/v4"
)

type TemplateRenderer struct {
	templates *template.Template
}

// Render renders a template document
func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {

	return t.templates.ExecuteTemplate(w, name, data)
}

func UserRoute(e *echo.Echo) *echo.Echo {
	renderer := &TemplateRenderer{
		templates: template.Must(template.ParseGlob("user/template/*.html")),
	}

	e.Renderer = renderer

	e.GET("/", func(c echo.Context) error {
		return controller.HomePageVoteController(c)
	})

	e.POST("/uservote", func(c echo.Context) error {
		return controller.UserVoteController(c)
	})

	return e

}
