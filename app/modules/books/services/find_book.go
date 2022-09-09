package services

import (
	"goapi/app/modules/books/dto"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func (repo BookServicesImpl) GetBook(ctx *fiber.Ctx, id uuid.UUID) (*dto.BookResponse, error) {
	// Get all books.
	books, err := repo.BookRepository.GetBook(ctx, id)
	if err != nil {
		// Return, if books not found.
		return nil, err
	}

	res := dto.BookResponse{
		ID:         books.ID,
		CreatedAt:  books.CreatedAt,
		UpdatedAt:  books.UpdatedAt,
		UserID:     books.UserID,
		Title:      books.Title,
		Author:     books.Author,
		BookStatus: books.BookStatus,
		BookAttrs:  dto.BookAttrs(books.BookAttrs),
	}

	// Return status 200 OK.
	return &res, nil

}
