package dto

import "github.com/google/uuid"

type bookAttrs struct {
	Picture     string `json:"picture"`
	Description string `json:"description"`
	Rating      int    `json:"rating" validate:"min=1,max=10"`
}

type BookCreateRequest struct {
	Title     string    `json:"title" validate:"required,lte=255"`
	Author    string    `json:"author" validate:"required,lte=255"`
	BookAttrs bookAttrs `json:"book_attrs" validate:"required,dive"`
}

type BookUpdateRequest struct {
	Title      string    `json:"title" validate:"lte=255"`
	Author     string    `json:"author" validate:"lte=255"`
	UserID     uuid.UUID `json:"user_id" validate:"uuid"`
	BookStatus int       `json:"book_status"`
	BookAttrs  bookAttrs `json:"book_attrs" validate:"dive"`
}

type BookDeleteRequest struct {
	ID uuid.UUID `json:"id" validate:"required,uuid"`
}
