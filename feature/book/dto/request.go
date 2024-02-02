package dto

import (
	"github.com/oggnimodd/golang-clerk-practice/entities"
	"github.com/oggnimodd/golang-clerk-practice/utils"
)

type CreateBookRequest struct {
	Title          string  `json:"title" binding:"required"`
	Description    *string `json:"description"`
	Author         *string `json:"author"`
	Cover          *string `json:"cover"`
	GoogleBooksUrl *string `json:"google_books_url"`
	ShelfID        string  `json:"shelf_id" binding:"required"`
}

type UpdateBookRequest struct {
	Title       *string `json:"title"`
	Description *string `json:"description"`
	Author      *string `json:"author"`
	Cover       *string `json:"cover"`
	ShelfID     string  `json:"shelf_id"`
}

func FormatCreateBookRequest(b *CreateBookRequest) *entities.Book {
	return &entities.Book{
		Title:          b.Title,
		Description:    b.Description,
		Author:         b.Author,
		Cover:          b.Cover,
		GoogleBooksUrl: b.GoogleBooksUrl,
		ShelfID:        b.ShelfID,
	}
}

func FormatUpdateBookRequest(b *UpdateBookRequest) *entities.Book {
	return &entities.Book{
		Title:       utils.SafeString(b.Title),
		Description: b.Description,
		Author:      b.Author,
		Cover:       b.Cover,
		ShelfID:     b.ShelfID,
	}
}
