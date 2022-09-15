package services

import (
	"goapi/app/modules/books/dto"
	"goapi/shared/convert"

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

	res := dto.BookResponse{}

	err = convert.ToStruct(*books, &res)
	if err != nil {
		// Return, if books not found.
		return nil, err
	}

	// Return status 200 OK.
	return &res, nil

}
