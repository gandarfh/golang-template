package services

import (
	"goapi/internal/modules/books/repositories"

	"github.com/gofiber/fiber/v2"
)

type BookServicesImpl struct {
	BookRepository *repositories.BookRepositoryImpl
}

func NewBookService(ctx *fiber.Ctx) (*BookServicesImpl, error) {
	bookRepository, err := repositories.NewBookRepository(ctx)

	if err != nil {
		return nil, err
	}

	return &BookServicesImpl{bookRepository}, nil
}
