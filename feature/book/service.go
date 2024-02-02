package book

import (
	"github.com/oggnimodd/golang-clerk-practice/entities"
	"github.com/oggnimodd/golang-clerk-practice/utils"
)

type BookService struct {
	bookRepo RepositoryBookInterface
}

func NewBookService(bookRepo RepositoryBookInterface) ServiceBookInterface {
	return &BookService{
		bookRepo: bookRepo,
	}
}

func (s *BookService) GetAllBooks(p *utils.Pagination, userId string) ([]entities.Book, error) {
	return s.bookRepo.GetAllBooks(p, userId)
}

func (s *BookService) GetBookById(id string, userId string) (entities.Book, error) {
	return s.bookRepo.GetBookById(id, userId)
}

func (s *BookService) GetBooksByShelfId(p *utils.Pagination, shelfId string, userId string) ([]entities.Book, error) {
	return s.bookRepo.GetBooksByShelfId(p, shelfId, userId)
}

func (s *BookService) CreateBook(book *entities.Book, userId string) (*entities.Book, error) {
	return s.bookRepo.CreateBook(book, userId)
}

func (s *BookService) UpdateBook(id string, book *entities.Book, userId string) (*entities.Book, error) {
	return s.bookRepo.UpdateBook(id, book, userId)
}

func (s *BookService) DeleteBook(id string, userId string) error {
	return s.bookRepo.DeleteBook(id, userId)
}
