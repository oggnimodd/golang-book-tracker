package entities

import (
	"gorm.io/gorm"
)

type Shelf struct {
	gorm.Model
	ID        string `gorm:"type:varchar(36);primary_key"`
	UserID    string `gorm:"type:varchar(50);index"`
	Type      string `gorm:"type:varchar(50)"`
	IsDefault *bool  `gorm:"type:boolean"`
	Books     []Book `gorm:"foreignKey:ShelfID"`
}
