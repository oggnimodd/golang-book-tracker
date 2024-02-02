package note

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/oggnimodd/golang-clerk-practice/entities"
	"github.com/oggnimodd/golang-clerk-practice/utils"
)

type NoteHandler struct {
	noteService ServiceNoteInterface
}

func NewNoteHandler(noteService ServiceNoteInterface) HandlerNoteInterface {
	return &NoteHandler{
		noteService: noteService,
	}
}

// @Summary Get All Notes
// @Description Retrieve all notes
// @Produce json
// @Param book_id path string true "Book ID"
// @Success 200 {array} entities.Notes
// @Failure 500 {object} map[string]interface{}
// @Router /notes/{book_id} [get]
func (h *NoteHandler) GetAllNotes(c *gin.Context) {
	user, err := utils.GetUserFromContext(c)
	if err != nil {
		utils.Unauthorized(c)
		return
	}

	p, err := utils.GetPagination(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	bookId := c.Param("book_id")
	notes, err := h.noteService.GetAllNotes(p, user.ID, bookId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, notes)
}

// @Summary Get Note By ID
// @Description Retrieve a note by ID
// @Produce json
// @Param id path string true "Note ID"
// @Success 200 {object} entities.Notes
// @Failure 500 {object} map[string]interface{}
// @Router /notes/{id} [get]
func (h *NoteHandler) GetNoteById(c *gin.Context) {
	user, err := utils.GetUserFromContext(c)
	if err != nil {
		utils.Unauthorized(c)
		return
	}

	id := c.Param("id")
	note, err := h.noteService.GetNoteById(id, user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, note)
}

// @Summary Create Note
// @Description Create a new note
// @Accept json
// @Produce json
// @Param book_id path string true "Book ID"
// @Param note body entities.Notes true "Create note"
// @Success 201 {object} entities.Notes
// @Failure 400 {object} map[string]interface{}
// @Router /notes/{book_id} [post]
func (h *NoteHandler) CreateNote(c *gin.Context) {
	user, err := utils.GetUserFromContext(c)
	if err != nil {
		utils.Unauthorized(c)
		return
	}

	var note entities.Notes
	if err := c.ShouldBindJSON(&note); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	bookId := c.Param("book_id")
	newNote, err := h.noteService.CreateNote(&note, user.ID, bookId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, newNote)
}

// @Summary Update Note
// @Description Update an existing note
// @Accept json
// @Produce json
// @Param id path string true "Note ID"
// @Param note body entities.Notes true "Update note"
// @Success 200 {object} entities.Notes
// @Failure 400 {object} map[string]interface{}
// @Router /notes/{id} [put]
func (h *NoteHandler) UpdateNote(c *gin.Context) {
	user, err := utils.GetUserFromContext(c)
	if err != nil {
		utils.Unauthorized(c)
		return
	}

	id := c.Param("id")
	var updateNote entities.Notes
	if err := c.ShouldBindJSON(&updateNote); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedNote, err := h.noteService.UpdateNote(id, &updateNote, user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, updatedNote)
}

// @Summary Delete Note
// @Description Delete an existing note
// @Accept json
// @Produce json
// @Param id path string true "Note ID"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Router /notes/{id} [delete]
func (h *NoteHandler) DeleteNote(c *gin.Context) {
	user, err := utils.GetUserFromContext(c)
	if err != nil {
		utils.Unauthorized(c)
		return
	}

	id := c.Param("id")
	err = h.noteService.DeleteNote(id, user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, nil)
}
