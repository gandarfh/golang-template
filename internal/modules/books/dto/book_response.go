package dto

import (
	"goapi/internal/modules/books/entities"
	"goapi/pkg/pagination"
	"time"

	"github.com/google/uuid"
)

// BookAttrs struct to describe book attributes.
type BookAttrs struct {
	Picture     string `json:"picture"`
	Description string `json:"description"`
	Rating      *int   `json:"rating"`
}

type BookResponse struct {
	ID         *uuid.UUID `json:"id"`
	CreatedAt  *time.Time `json:"created_at"`
	UpdatedAt  *time.Time `json:"updated_at"`
	UserID     *uuid.UUID `json:"user_id"`
	Title      string     `json:"title"`
	Author     string     `json:"author"`
	BookStatus *int       `json:"book_status"`
	BookAttrs  *BookAttrs `json:"book_attrs"`
}

// Book struct to describe book object.
type BookListAllResponse struct {
	*pagination.Paginate[entities.Book]
}

type BookDeleteResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}
