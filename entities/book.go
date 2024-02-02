package entities

import (
	"time"

	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	ID             string    `gorm:"column:id;type:varchar(36);primary_key"`
	Title          string    `gorm:"column:title;type:varchar(100)"`
	Description    *string   `gorm:"column:description;type:text"`
	Author         *string   `gorm:"column:author;type:varchar(100)"`
	Cover          *string   `gorm:"column:cover;type:text"`
	UserID         string    `gorm:"column:user_id;type:varchar(50)"`
	ShelfID        string    `gorm:"column:shelf_id;type:varchar(50);index"`
	Notes          []Notes   `gorm:"foreignKey:BookID"`
	Review         *Review   `gorm:"foreignKey:BookID"`
	Progress       *int      `gorm:"column:progress;type:integer"`
	GoogleBooksUrl *string   `gorm:"column:google_books_url;type:text;unique"`
	CreatedAt      time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt      time.Time `gorm:"column:updated_at;autoUpdateTime"`
}
