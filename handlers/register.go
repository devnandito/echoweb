package handlers

import (
	"html/template"
	"net/http"

	"github.com/labstack/echo"
)

func HandlerSignIn(c echo.Context) error {
	renderer := &TemplateRenderer{
		templates: template.Must(template.ParseGlob("views/users/*.html")),
  }
  c.Echo().Renderer = renderer
	title := "Register"

	return c.Render(http.StatusOK, "signin.html", map[string]interface{}{
		"Title": title,
	})

}