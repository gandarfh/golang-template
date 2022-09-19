package services

import (
	"goapi/internal/modules/books/dto"

	"github.com/gofiber/fiber/v2"
)

func (repo *BookServicesImpl) GetBooks(ctx *fiber.Ctx) (*dto.BookListAllResponse, error) {
	// Get all allBooks.
	res, err := repo.BookRepository.GetBooks(ctx)

	if err != nil {
		// Return, if books not found.
		return nil, err
	}

	// Return status 200 OK.
	return res, nil
}
