package services

import (
	"goapi/internal/modules/books/dto"
	"goapi/internal/modules/books/entities"
	"goapi/pkg/convert"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func (repo BookServicesImpl) CreateBook(ctx *fiber.Ctx, userId uuid.UUID, body *dto.BookCreateRequest) (*dto.BookResponse, error) {
	book := entities.Book{}
	convert.ToStruct(body, &book)

	id := uuid.New()
	status := 1
	now := time.Now()

	// Set initialized default data for book:
	book.ID = &id
	book.CreatedAt = &now
	book.BookStatus = &status
	book.UserID = &userId

	// Insert Book provided into database.
	newBook, err := repo.BookRepository.CreateBook(ctx, &book)
	if err != nil {
		// Propagate errors implemented inside repository
		return nil, err

	}

	response := dto.BookResponse{}
	convert.ToStruct(newBook, &response)

	// Return status 200 OK.
	return &response, nil
}
