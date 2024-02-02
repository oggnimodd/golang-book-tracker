package note

import (
	"github.com/google/uuid"
	"github.com/oggnimodd/golang-clerk-practice/entities"
	"github.com/oggnimodd/golang-clerk-practice/utils"
	"gorm.io/gorm"
)

type NoteRepository struct {
	db *gorm.DB
}

func NewNoteRepository(db *gorm.DB) RepositoryNoteInterface {
	return &NoteRepository{db: db}
}

func (r *NoteRepository) GetAllNotes(p *utils.Pagination, userId string, bookId string) ([]entities.Notes, error) {
	offset := p.GetOffset()
	limit := p.GetLimit()
	order := p.GetSort()

	var notes []entities.Notes
	result := r.db.Where("user_id = ? AND book_id = ?", userId, bookId).Order(order).Limit(limit).Offset(offset).Find(&notes)
	if result.Error != nil {
		return nil, result.Error
	}
	return notes, nil
}

func (r *NoteRepository) GetNoteById(id string, userId string) (entities.Notes, error) {
	var note entities.Notes
	result := r.db.Where("id = ? AND user_id = ?", id, userId).First(&note)
	if result.Error != nil {
		return entities.Notes{}, result.Error
	}
	return note, nil
}

func (r *NoteRepository) CreateNote(note *entities.Notes, userId string, bookId string) (*entities.Notes, error) {
	note.UserID = userId
	note.BookID = &bookId
	note.ID = uuid.New().String()
	result := r.db.Create(note)
	if result.Error != nil {
		return nil, result.Error
	}
	return note, nil
}

func (r *NoteRepository) UpdateNote(id string, note *entities.Notes, userId string) (*entities.Notes, error) {
	result := r.db.Model(&entities.Notes{}).Where("id = ? AND user_id = ?", id, userId).Updates(note)
	if result.Error != nil {
		return nil, result.Error
	}
	return note, nil
}

func (r *NoteRepository) DeleteNote(id string, userId string) error {
	result := r.db.Where("id = ? AND user_id = ?", id, userId).Delete(&entities.Notes{})
	if result.Error != nil {
		return result.Error
	}
	return nil
}
