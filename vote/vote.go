package vote

import (
	"fmt"
	"html/template"
	"io"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Vote struct {
	Vote string `json:"vote"`
}

type TemplateRenderer struct {
	templates *template.Template
}

// Render renders a template document
func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {

	// Add global methods if data is a map
	if viewContext, isMap := data.(map[string]interface{}); isMap {
		viewContext["reverse"] = c.Echo().Reverse
	}

	return t.templates.ExecuteTemplate(w, name, data)
}

func VoteService(e *echo.Echo) *echo.Echo {
	renderer := &TemplateRenderer{
		templates: template.Must(template.ParseGlob("vote/template/*.html")),
	}

	e.Renderer = renderer

	e.GET("/", func(c echo.Context) error {
		return c.Render(http.StatusOK, "index.html", map[string]interface{}{})
	}).Name = "home"

	e.POST("/vote", func(c echo.Context) error {

		name := c.FormValue("name")
		socialnetwork := c.FormValue("socialnetwork")
		fmt.Println(name)
		fmt.Println(socialnetwork)

		return c.Render(http.StatusOK, "index.html", map[string]interface{}{})
	})

	return e
}
