package entities

import (
	"gorm.io/gorm"
)

type Notes struct {
	gorm.Model
	ID          string  `gorm:"type:varchar(36);primary_key"`
	Description string  `gorm:"type:text"`
	UserID      string  `gorm:"type:varchar(50);index"`
	BookID      *string `gorm:"type:varchar(50);index"`
}
