package route

import (
	"gomq/vote/controller"
	"html/template"
	"io"

	"github.com/labstack/echo/v4"
)

type TemplateRenderer struct {
	templates *template.Template
}

// Render renders a template document
func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {

	if viewContext, isMap := data.(map[string]interface{}); isMap {
		viewContext["reverse"] = c.Echo().Reverse
	}

	return t.templates.ExecuteTemplate(w, name, data)
}

func VoteRoute(e *echo.Echo, IvoteController controller.IVoteController) *echo.Echo {
	renderer := &TemplateRenderer{
		templates: template.Must(template.ParseGlob("vote/template/*.html")),
	}

	e.Renderer = renderer

	e.GET("/vote", func(c echo.Context) error {
		return IvoteController.ShowTheVotesController(c)
	})

	return e
}
