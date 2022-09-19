package controllers

import (
	"goapi/internal/modules/books/dto"
	"goapi/internal/modules/books/services"
	"goapi/pkg/errors"
	"goapi/pkg/jwt"
	"goapi/pkg/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// GetBooks func gets all exists books.
// @Description Get all exists books.
// @Summary get all exists books
// @Tags Books
// @Accept json
// @Produce json
// @Param limit query int false "limit" minimum(1)
// @Param page query int false "page" minimum(1)
// @Success 200 {object} dto.BookListAllResponse
// @Router /v1/books [get]
func GetBooks(c *fiber.Ctx) error {
	service, err := services.NewBookService(c)
	if err != nil {
		return errors.ErrorResponse(c, err)
	}

	// Get all books.
	books, err := service.GetBooks(c)
	if err != nil {
		return errors.ErrorResponse(c, err)
	}

	// Return status 200 OK.
	return c.Status(fiber.StatusOK).JSON(*books)
}

// GetBook func gets book by given ID or 404 error.
// @Description Get book by given ID.
// @Summary get book by given ID
// @Tags Books
// @Accept json
// @Produce json
// @Param id path string true "Book ID"
// @Success 200 {object} dto.BookResponse
// @Router /v1/books/{id} [get]
func GetBook(c *fiber.Ctx) error {
	service, err := services.NewBookService(c)
	if err != nil {
		return errors.ErrorResponse(c, err)
	}

	// Catch book ID from URL.
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	// Get book by ID.
	book, err := service.GetBook(c, id)
	if err != nil {
		// Return, if book not found.
		return errors.ErrorResponse(c, err)
	}

	return c.Status(fiber.StatusOK).JSON(book)
}

// CreateBook func for creates a new book.
// @Description Create a new book.
// @Summary create a new book
// @Tags Books
// @Accept json
// @Produce json
// @Param request body dto.BookCreateRequest true "Request Body"
// @Success 200 {object} dto.BookResponse
// @Security ApiKeyAuth
// @Router /v1/books [post]
func CreateBook(c *fiber.Ctx) error {
	service, err := services.NewBookService(c)
	if err != nil {
		return errors.ErrorResponse(c, err)
	}

	// Get claims from JWT.
	claims, err := jwt.ExtractTokenMetadata(c)
	if err != nil {
		// Return status 500 and JWT parse error.
		return errors.ErrorResponse(c, err)
	}

	// Parse body provided to entities
	book := dto.BookCreateRequest{}
	c.BodyParser(&book)

	// Create a new validator for a Book model.
	validate := utils.NewValidator()

	// Validate book fields.
	if err := validate.Struct(book); err != nil {
		// Return, if some fields are not valid.
		return c.Status(fiber.StatusUnprocessableEntity).JSON(utils.ValidatorErrors(err))
	}

	// Create book by given model.
	response, err := service.CreateBook(c, claims.UserID, &book)
	if err != nil {
		return errors.ErrorResponse(c, err)
	}

	// Return status 200 OK.
	return c.Status(fiber.StatusCreated).JSON(response)
}

// UpdateBook func for updates book by given ID.
// @Description Update book.
// @Summary update book
// @Tags Books
// @Accept json
// @Produce json
// @Param request body dto.BookUpdateRequest true "Request Body"
// @Success 202 {string} status "ok"
// @Security ApiKeyAuth
// @Router /v1/books [put]
func UpdateBook(c *fiber.Ctx) error {
	service, err := services.NewBookService(c)

	if err != nil {
		return errors.ErrorResponse(c, err)
	}

	// Catch book ID from URL.
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	// // Create new Book struct
	body := &dto.BookUpdateRequest{}
	c.BodyParser(body)

	// Create a new validator for a Book model.
	validate := utils.NewValidator()

	// Validate book fields.
	if err := validate.Struct(body); err != nil {
		// Return, if some fields are not valid.
		return c.Status(fiber.StatusUnprocessableEntity).JSON(utils.ValidatorErrors(err))
	}

	// Update book by given ID.
	res, err := service.UpdateBook(c, &id, body)

	if err != nil {
		return errors.ErrorResponse(c, err)
	}

	// Return status 201.
	return c.Status(fiber.StatusCreated).JSON(res)
}

// DeleteBook func for deletes book by given ID.
// @Description Delete book by given ID.
// @Summary delete book by given ID
// @Tags Books
// @Accept json
// @Produce json
// @Param request body dto.BookDeleteRequest true "Book ID"
// @Success 204 {string} status "ok"
// @Security ApiKeyAuth
// @Router /v1/books [delete]
func DeleteBook(c *fiber.Ctx) error {
	// // Get now time.
	// now := time.Now().Unix()

	// // Get claims from JWT.
	// claims, err := utils.ExtractTokenMetadata(c)
	// if err != nil {
	// 	// Return status 500 and JWT parse error.
	// 	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
	// 		"error": true,
	// 		"msg":   err.Error(),
	// 	})
	// }

	// // Set expiration time from JWT data of current book.
	// expires := claims.Expires

	// // Checking, if now time greather than expiration from JWT.
	// if now > expires {
	// 	// Return status 401 and unauthorized error message.
	// 	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
	// 		"error": true,
	// 		"msg":   "unauthorized, check expiration time of your token",
	// 	})
	// }

	// // Set credential `book:delete` from JWT data of current book.
	// credential := claims.Credentials[permissions.BookDeleteCredential]

	// // Only book creator with `book:delete` credential can delete his book.
	// if !credential {
	// 	// Return status 403 and permission denied error message.
	// 	return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
	// 		"error": true,
	// 		"msg":   "permission denied, check credentials of your token", }) }

	// // Create new Book struct
	// book := &dto.BookListAllResponse

	// // Check, if received JSON data is valid.
	// if err := c.BodyParser(book); err != nil {
	// 	// Return status 400 and error message.
	// 	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
	// 		"error": true,
	// 		"msg":   err.Error(),
	// 	})
	// }

	// // Create a new validator for a Book model.
	// validate := utils.NewValidator()

	// // Validate book fields.
	// if err := validate.StructPartial(book, "id"); err != nil {
	// 	// Return, if some fields are not valid.
	// 	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
	// 		"error": true,
	// 		"msg":   utils.ValidatorErrors(err),
	// 	})
	// }

	// // Create database connection.
	// db, err := database.OpenDBConnection()
	// if err != nil {
	// 	// Return status 500 and database connection error.
	// 	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
	// 		"error": true,
	// 		"msg":   err.Error(),
	// 	})
	// }

	// // Checking, if book with given ID is exists.
	// foundedBook, err := db.GetBook(book.ID)
	// if err != nil {
	// 	// Return status 404 and book not found error.
	// 	return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
	// 		"error": true,
	// 		"msg":   "book with this ID not found",
	// 	})
	// }

	// // Set user ID from JWT data of current user.
	// userID := claims.UserID

	// // Only the creator can delete his book.
	// if foundedBook.UserID == userID {
	// 	// Delete book by given ID.
	// 	if err := db.DeleteBook(foundedBook.ID); err != nil {
	// 		// Return status 500 and error message.
	// 		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
	// 			"error": true,
	// 			"msg":   err.Error(),
	// 		})
	// 	}

	// 	// Return status 204 no content.
	// 	return c.SendStatus(fiber.StatusNoContent)
	// } else {
	// 	// Return status 403 and permission denied error message.
	// 	return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
	// 		"error": true,
	// 		"msg":   "permission denied, only the creator can delete his book",
	// 	})
	// }

	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
		"error": true,
	})
}
