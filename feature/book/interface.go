package book

import (
	"github.com/gin-gonic/gin"
	"github.com/oggnimodd/golang-clerk-practice/entities"
	"github.com/oggnimodd/golang-clerk-practice/utils"
)

// In the RepositoryBookInterface
type RepositoryBookInterface interface {
	GetAllBooks(p *utils.Pagination, userId string) ([]entities.Book, error)
	GetBookById(id string, userId string) (entities.Book, error)
	GetBooksByShelfId(p *utils.Pagination, shelfId string, userId string) ([]entities.Book, error)
	CreateBook(book *entities.Book, userId string) (*entities.Book, error)
	UpdateBook(id string, book *entities.Book, userId string) (*entities.Book, error)
	DeleteBook(id string, userId string) error
}

// In the BookService
type ServiceBookInterface interface {
	GetAllBooks(p *utils.Pagination, userId string) ([]entities.Book, error)
	GetBookById(id string, userId string) (entities.Book, error)
	GetBooksByShelfId(p *utils.Pagination, shelfId string, userId string) ([]entities.Book, error)
	CreateBook(book *entities.Book, userId string) (*entities.Book, error)
	UpdateBook(id string, book *entities.Book, userId string) (*entities.Book, error)
	DeleteBook(id string, userId string) error
}

// In the BookHandler
type HandlerBookInterface interface {
	GetAllBooks(c *gin.Context)
	GetBookById(c *gin.Context)
	GetBooksByShelfId(c *gin.Context)
	CreateBook(c *gin.Context)
	UpdateBook(c *gin.Context)
	DeleteBook(c *gin.Context)
}
