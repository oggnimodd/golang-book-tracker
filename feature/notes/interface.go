package note

import (
	"github.com/gin-gonic/gin"
	"github.com/oggnimodd/golang-clerk-practice/entities"
	"github.com/oggnimodd/golang-clerk-practice/utils"
)

type RepositoryNoteInterface interface {
	GetAllNotes(p *utils.Pagination, userId string, bookId string) ([]entities.Notes, error)
	GetNoteById(id string, userId string) (entities.Notes, error)
	CreateNote(note *entities.Notes, userId string, bookId string) (*entities.Notes, error)
	UpdateNote(id string, note *entities.Notes, userId string) (*entities.Notes, error)
	DeleteNote(id string, userId string) error
}

type ServiceNoteInterface interface {
	GetAllNotes(p *utils.Pagination, userId string, bookId string) ([]entities.Notes, error)
	GetNoteById(id string, userId string) (entities.Notes, error)
	CreateNote(note *entities.Notes, userId string, bookId string) (*entities.Notes, error)
	UpdateNote(id string, note *entities.Notes, userId string) (*entities.Notes, error)
	DeleteNote(id string, userId string) error
}

type HandlerNoteInterface interface {
	GetAllNotes(c *gin.Context)
	GetNoteById(c *gin.Context)
	CreateNote(c *gin.Context)
	UpdateNote(c *gin.Context)
	DeleteNote(c *gin.Context)
}
