package books

import (
	"goapi/app/middleware"
	"goapi/app/permissions"

	"github.com/gofiber/fiber/v2"
)

// All public routes of book module
func PublicRoutes(route fiber.Router) {
	// get list of all books
	route.Get("/books", GetBooks)

	// get one book by ID
	route.Get("/books/:id", GetBook)
}

// All private routes of book module
func PrivateRoutes(route fiber.Router) {
	// create a new book
	route.Post("/books",
		middleware.Credentials(permissions.BookCreateCredential),
		CreateBook)

	// Update especific book
	route.Patch("/books/:id",
		middleware.Credentials(permissions.BookUpdateCredential),
		UpdateBook)
}
