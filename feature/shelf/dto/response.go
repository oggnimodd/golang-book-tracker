package dto

import (
	"github.com/oggnimodd/golang-clerk-practice/entities"
)

type ShelfResponse struct {
	ID        string          `json:"id"`
	Type      string          `json:"type"`
	IsDefault *bool           `json:"is_default"`
	Books     []entities.Book `json:"books"`
}

type GenerateShelfResponse []entities.Shelf

func FormatCreateShelfResponse(s entities.Shelf) *ShelfResponse {
	return &ShelfResponse{
		ID:        s.ID,
		Type:      s.Type,
		IsDefault: s.IsDefault,
	}
}

func FormatGenerateShelfResponse(s []entities.Shelf) []ShelfResponse {
	var res []ShelfResponse
	for _, v := range s {
		res = append(res, ShelfResponse{
			ID:        v.ID,
			Type:      v.Type,
			IsDefault: v.IsDefault,
			Books:     v.Books,
		})
	}
	return res
}
