// This file is to do	DB initialization using sqlite and gorm

package db

import (
	"fmt"

	// Models
	"github.com/oggnimodd/golang-clerk-practice/entities"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	err := db.AutoMigrate(
		&entities.Book{},
		&entities.Shelf{},
		&entities.Review{},
		&entities.Notes{},
	)

	if err != nil {
		fmt.Println(err)
		panic("failed to migrate database")
	}
}

func Init() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("bookstore.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	Migrate(db)

	fmt.Println("Connection Opened to Database")
	return db
}
