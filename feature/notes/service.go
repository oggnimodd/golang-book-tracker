package note

import (
	"github.com/oggnimodd/golang-clerk-practice/entities"
	"github.com/oggnimodd/golang-clerk-practice/utils"
)

type NoteService struct {
	noteRepo RepositoryNoteInterface
}

func NewNoteService(noteRepo RepositoryNoteInterface) ServiceNoteInterface {
	return &NoteService{
		noteRepo: noteRepo,
	}
}

func (s *NoteService) GetAllNotes(p *utils.Pagination, userId string, bookId string) ([]entities.Notes, error) {
	return s.noteRepo.GetAllNotes(p, userId, bookId)
}

func (s *NoteService) GetNoteById(id string, userId string) (entities.Notes, error) {
	return s.noteRepo.GetNoteById(id, userId)
}

func (s *NoteService) CreateNote(note *entities.Notes, userId string, bookId string) (*entities.Notes, error) {
	return s.noteRepo.CreateNote(note, userId, bookId)
}

func (s *NoteService) UpdateNote(id string, note *entities.Notes, userId string) (*entities.Notes, error) {
	return s.noteRepo.UpdateNote(id, note, userId)
}

func (s *NoteService) DeleteNote(id string, userId string) error {
	return s.noteRepo.DeleteNote(id, userId)
}
