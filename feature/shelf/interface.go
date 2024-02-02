package shelf

import (
	"github.com/gin-gonic/gin"
	"github.com/oggnimodd/golang-clerk-practice/entities"
)

type RepositoryShelfInterface interface {
	GetAllShelves(userId string) ([]entities.Shelf, error)
	GetShelfById(id string, userId string) (entities.Shelf, error)
	CreateShelf(shelf *entities.Shelf, userId string) (*entities.Shelf, error)
	UpdateShelf(id string, shelf *entities.Shelf, userId string) (*entities.Shelf, error)
	DeleteShelf(id string, userId string) error
}

type ServiceShelfInterface interface {
	GetAllShelves(userId string) ([]entities.Shelf, error)
	GetShelfById(id string, userId string) (entities.Shelf, error)
	CreateShelf(shelf *entities.Shelf, userId string) (*entities.Shelf, error)
	UpdateShelf(id string, shelf *entities.Shelf, userId string) (*entities.Shelf, error)
	DeleteShelf(id string, userId string) error
	GenerateInitialShelves(userId string) ([]entities.Shelf, error)
}

type HandlerShelfInterface interface {
	GetAllShelves(c *gin.Context)
	GetShelfById(c *gin.Context)
	CreateShelf(c *gin.Context)
	UpdateShelf(c *gin.Context)
	DeleteShelf(c *gin.Context)
	GenerateInitialShelves(c *gin.Context)
}
