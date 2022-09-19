package services

import (
	"goapi/internal/modules/books/dto"
	"goapi/internal/modules/books/entities"
	"goapi/pkg/convert"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func (repo *BookServicesImpl) UpdateBook(ctx *fiber.Ctx, id *uuid.UUID, body *dto.BookUpdateRequest) (*dto.BookResponse, error) {
	book := entities.Book{}
	// Parse from dto to entities struct
	convert.ToStruct(*body, &book)

	now := time.Now()
	book.UpdatedAt = &now

	// Get all books.
	books, err := repo.BookRepository.UpdateBook(ctx, id, &book)
	if err != nil {
		// Return, if books not found.
		return nil, err
	}

	res := dto.BookResponse{}
	convert.ToStruct(*books, &res)

	// Return status 200 OK.
	return &res, nil
}
