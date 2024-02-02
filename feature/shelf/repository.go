package shelf

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/oggnimodd/golang-clerk-practice/entities"
	"gorm.io/gorm"
)

type RepositoryShelf struct {
	db *gorm.DB
}

func NewRepositoryShelf(db *gorm.DB) RepositoryShelfInterface {
	return &RepositoryShelf{db: db}
}

func (r *RepositoryShelf) GetAllShelves(userId string) ([]entities.Shelf, error) {
	var shelves []entities.Shelf
	result := r.db.Where("user_id = ?", userId).Find(&shelves)
	if result.Error != nil {
		return nil, result.Error
	}
	return shelves, nil
}

func (r *RepositoryShelf) GetShelfById(id string, userId string) (entities.Shelf, error) {
	var shelf entities.Shelf
	result := r.db.Where("id = ? AND user_id = ?", id, userId).First(&shelf)
	if result.Error != nil {
		return entities.Shelf{}, result.Error
	}
	return shelf, nil
}

func (r *RepositoryShelf) CreateShelf(shelf *entities.Shelf, userId string) (*entities.Shelf, error) {
	fmt.Println("shelf", shelf)
	shelf.UserID = userId
	shelf.ID = uuid.New().String()
	result := r.db.Create(shelf)
	if result.Error != nil {
		return nil, result.Error
	}
	return shelf, nil
}

func (r *RepositoryShelf) UpdateShelf(id string, shelf *entities.Shelf, userId string) (*entities.Shelf, error) {
	result := r.db.Model(&entities.Shelf{}).Where("id = ? AND user_id = ?", id, userId).Updates(shelf)
	if result.Error != nil {
		return nil, result.Error
	}
	return shelf, nil
}

func (r *RepositoryShelf) DeleteShelf(id string, userId string) error {
	result := r.db.Where("id = ? AND user_id = ?", id, userId).Delete(&entities.Shelf{})
	if result.Error != nil {
		return result.Error
	}
	return nil
}
