package entities

import (
	"gorm.io/gorm"
)

type Review struct {
	gorm.Model
	ID          string   `gorm:"type:varchar(36);primary_key"`
	Rating      *float64 `gorm:"type:real"`
	Description string   `gorm:"type:text"`
	UserID      string   `gorm:"type:varchar(50);index"`
	BookID      *string  `gorm:"type:varchar(50);index;unique"`
}
