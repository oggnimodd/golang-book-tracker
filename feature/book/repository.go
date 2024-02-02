package book

import (
	"github.com/google/uuid"
	"github.com/oggnimodd/golang-clerk-practice/entities"
	"github.com/oggnimodd/golang-clerk-practice/utils"
	"gorm.io/gorm"
)

type BookRepository struct {
	db *gorm.DB
}

func NewBookRepository(db *gorm.DB) RepositoryBookInterface {
	return &BookRepository{db: db}
}

func (r *BookRepository) GetAllBooks(p *utils.Pagination, userId string) ([]entities.Book, error) {
	offset := p.GetOffset()
	limit := p.GetLimit()
	order := p.GetSort()

	var books []entities.Book
	result := r.db.Where("user_id = ?", userId).Order(order).Limit(limit).Offset(offset).Find(&books)
	if result.Error != nil {
		return nil, result.Error
	}
	return books, nil
}

func (r *BookRepository) GetBookById(id string, userId string) (entities.Book, error) {
	var book entities.Book
	result := r.db.Where("id = ? AND user_id = ?", id, userId).First(&book)
	if result.Error != nil {
		return entities.Book{}, result.Error
	}
	return book, nil
}

func (r *BookRepository) GetBooksByShelfId(p *utils.Pagination, shelfId string, userId string) ([]entities.Book, error) {
	offset := p.GetOffset()
	limit := p.GetLimit()
	order := p.GetSort()

	var books []entities.Book
	result := r.db.Where("shelf_id = ? AND user_id = ?", shelfId, userId).Order(order).Limit(limit).Offset(offset).Find(&books)
	if result.Error != nil {
		return nil, result.Error
	}
	return books, nil
}

func (r *BookRepository) CreateBook(book *entities.Book, userId string) (*entities.Book, error) {
	book.UserID = userId
	book.ID = uuid.New().String()
	result := r.db.Create(book)
	if result.Error != nil {
		return nil, result.Error
	}
	return book, nil
}

func (r *BookRepository) UpdateBook(id string, book *entities.Book, userId string) (*entities.Book, error) {
	result := r.db.Model(&entities.Book{}).Where("id = ? AND user_id = ?", id, userId).Updates(book)
	if result.Error != nil {
		return nil, result.Error
	}
	return book, nil
}

func (r *BookRepository) DeleteBook(id string, userId string) error {
	result := r.db.Where("id = ? AND user_id = ?", id, userId).Delete(&entities.Book{})
	if result.Error != nil {
		return result.Error
	}
	return nil
}
