package dto

import "github.com/oggnimodd/golang-clerk-practice/entities"

type BookResponse struct {
	ID          string  `json:"id"`
	Title       string  `json:"title"`
	Description *string `json:"description"`
	Author      *string `json:"author"`
	Cover       *string `json:"cover"`
	ShelfID     string  `json:"shelf_id"`
}

func FormatCreateBookResponse(b entities.Book) *BookResponse {
	return &BookResponse{
		ID:          b.ID,
		Title:       b.Title,
		Description: b.Description,
		Author:      b.Author,
		Cover:       b.Cover,
		ShelfID:     b.ShelfID,
	}
}
