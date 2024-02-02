package dto

import (
	"github.com/oggnimodd/golang-clerk-practice/entities"
	"github.com/oggnimodd/golang-clerk-practice/utils"
)

type CreateShelfRequest struct {
	Type      string `json:"type" binding:"required"`
	IsDefault *bool  `json:"is_default"`
}

type UpdateShelfRequest struct {
	Type      *string `json:"type"`
	IsDefault *bool   `json:"is_default"`
}

func FormatCreateShelfRequest(s *CreateShelfRequest) *entities.Shelf {
	return &entities.Shelf{
		Type:      s.Type,
		IsDefault: s.IsDefault,
	}
}

func FormatUpdateShelfRequest(s *UpdateShelfRequest) *entities.Shelf {
	return &entities.Shelf{
		Type:      utils.SafeString(s.Type),
		IsDefault: s.IsDefault,
	}
}
