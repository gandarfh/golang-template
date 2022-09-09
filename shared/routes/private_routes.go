package routes

import (
	"goapi/app/middleware"
	"goapi/app/modules/books"

	"github.com/gofiber/fiber/v2"
)

// PrivateRoutes func for describe group of private routes.
func PrivateRoutes(app *fiber.App) {
	// Create routes group.
	api := app.Group("/api/v1", middleware.JWTProtected())

	books.PrivateRoutes(api)
}
