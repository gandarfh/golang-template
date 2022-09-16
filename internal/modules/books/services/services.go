package services

import (
	"goapi/internal/modules/books/dto"
	"goapi/internal/modules/books/repositories"

	"github.com/gofiber/fiber/v2"
)

type BookServices interface {
	GetBooks(ctx *fiber.Ctx) (*dto.BookListAllResponse, error)
	GetBook(ctx *fiber.Ctx, bookId int) (*dto.BookResponse, error)
	CreateBook(ctx *fiber.Ctx, body *dto.BookCreateRequest) (*dto.BookResponse, error)
	UpdateBook(ctx *fiber.Ctx, book *dto.BookUpdateRequest) (*dto.BookResponse, error)
	DeleteBook(ctx *fiber.Ctx, bookId int) (*dto.BookDeleteResponse, error)
}

type BookServicesImpl struct {
	BookRepository *repositories.BookRepo
}

func NewBookService(ctx *fiber.Ctx) (*BookServicesImpl, error) {
	repo, err := repositories.NewBookRepository(ctx)

	if err != nil {
		return nil, err
	}

	return &BookServicesImpl{
		BookRepository: repo,
	}, nil
}
