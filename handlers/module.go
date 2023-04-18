package handlers

import (
	"html/template"
	"net/http"

	"github.com/labstack/echo"

	"github.com/devnandito/echoweb/models"
)

var mod models.Module

func HandlerShowModule(c echo.Context) error {
	renderer := &TemplateRenderer{
		templates: template.Must(template.ParseGlob("views/modules/*.html")),
  }
  c.Echo().Renderer = renderer
	headers := [3]string{"ID", "Description", "Action"}
	response, err := mod.ShowModuleGorm()

	if err != nil {
		panic(err)
	}

	return c.Render(http.StatusOK, "show.html", map[string]interface{}{
		"Objects": response,
		"Headers": headers,
	})

}