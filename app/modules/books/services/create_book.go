package services

import (
	"goapi/app/modules/books/dto"
	"goapi/app/modules/books/entities"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// GetBooks(ctx *fiber.Ctx) (*dto.BookListAllResponse, errors.RespError)
// GetBook(ctx *fiber.Ctx, bookId int) (*dto.BookResponse, errors.RespError)
// CreateBook(ctx *fiber.Ctx, book *dto.BookCreateRequest) (dto.BookResponse, errors.RespError)
// UpdateBook(ctx *fiber.Ctx, book *dto.BookUpdateRequest) (dto.BookResponse, errors.RespError)
// DeleteBook(ctx *fiber.Ctx, bookId int) (dto.BookDeleteResponse, errors.RespError)

func (repo BookServicesImpl) CreateBook(ctx *fiber.Ctx, userId uuid.UUID, body *dto.BookCreateRequest) (*dto.BookResponse, error) {

	// // Set initialized default data for book:
	book := entities.Book{
		ID:         uuid.New(),
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
		BookStatus: 1,
		UserID:     userId,
		Title:      body.Title,
		Author:     body.Author,
		BookAttrs:  entities.BookAttrs(body.BookAttrs),
	}

	// Insert Book provided into database.
	newBook, err := repo.BookRepository.CreateBook(ctx, &book)
	if err != nil {
		// Propagate errors implemented inside repository
		return nil, err

	}

	response := dto.BookResponse{
		ID:         newBook.ID,
		CreatedAt:  newBook.CreatedAt,
		UpdatedAt:  newBook.UpdatedAt,
		UserID:     newBook.UserID,
		Title:      newBook.Title,
		Author:     newBook.Author,
		BookStatus: newBook.BookStatus,
		BookAttrs:  dto.BookAttrs(newBook.BookAttrs),
	}

	// Return status 200 OK.
	return &response, nil
}
