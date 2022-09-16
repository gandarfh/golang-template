package routes

import (
	"html/template"

	"github.com/gofiber/fiber/v2"

	swagger "github.com/gofiber/swagger"
)

// SwaggerRoute func for describe group of API Docs routes.
func SwaggerRoute(a *fiber.App) {
	// Create routes group.
	route := a.Group("/docs")

	// Routes for GET method:
	route.Get("*", swagger.New(swagger.Config{
		Title:  "Api boilerplate",
		Layout: "BaseLayout",
		Presets: []template.JS{
			template.JS("SwaggerUIBundle.presets.apis"),
			template.JS("SwaggerUIStandalonePreset"),
		},
	}))
}
