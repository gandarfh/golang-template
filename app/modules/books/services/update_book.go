package services

import (
	"goapi/app/modules/books/dto"
	"goapi/app/modules/books/entities"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func (repo BookServicesImpl) UpdateBook(ctx *fiber.Ctx, id *uuid.UUID, body *dto.BookUpdateRequest) (*dto.BookResponse, error) {
	// Parse from dto to entities struct
	book := entities.Book{
		ID:         *id,
		UpdatedAt:  time.Now(),
		UserID:     body.UserID,
		Title:      body.Title,
		Author:     body.Author,
		BookStatus: body.BookStatus,
		BookAttrs:  entities.BookAttrs(body.BookAttrs),
	}

	// Get all books.
	books, err := repo.BookRepository.UpdateBook(ctx, &book)
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
