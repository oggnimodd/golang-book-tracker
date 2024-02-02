package book

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/oggnimodd/golang-clerk-practice/feature/book/dto"
	"github.com/oggnimodd/golang-clerk-practice/utils"
)

type BookHandler struct {
	bookService ServiceBookInterface
}

func NewBookHandler(bookService ServiceBookInterface) HandlerBookInterface {
	return &BookHandler{
		bookService: bookService,
	}
}

// @Summary Get All Books
// @Description Retrieve all books
// @Produce json
// @Success 200 {array} entities.Book
// @Router /books [get]
func (h *BookHandler) GetAllBooks(c *gin.Context) {
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

	books, err := h.bookService.GetAllBooks(p, user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, books)
}

// @Summary Get Book By ID
// @Description Retrieve a book by ID
// @Produce json
// @Param id path string true "Book ID"
// @Success 200 {object} entities.Book
// @Failure 404 {object} map[string]interface{}
// @Router /books/{id} [get]
func (h *BookHandler) GetBookById(c *gin.Context) {
	user, err := utils.GetUserFromContext(c)
	if err != nil {
		utils.Unauthorized(c)
		return
	}

	id := c.Param("id")
	book, err := h.bookService.GetBookById(id, user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, book)
}

// @Summary Get Books By Shelf ID
// @Description Retrieve all books on a particular shelf
// @Produce json
// @Param shelf_id path string true "Shelf ID"
// @Success 200 {array} entities.Book
// @Router /shelves/{shelf_id}/books [get]
func (h *BookHandler) GetBooksByShelfId(c *gin.Context) {
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

	shelfId := c.Param("shelf_id")
	books, err := h.bookService.GetBooksByShelfId(p, shelfId, user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, books)
}

// @Summary Create Book
// @Description Create a new book
// @Accept json
// @Produce json
// @Param book body dto.CreateBookRequest true "Create book"
// @Success 201 {object} entities.Book
// @Failure 400 {object} map[string]interface{}
// @Router /books [post]
func (h *BookHandler) CreateBook(c *gin.Context) {
	user, err := utils.GetUserFromContext(c)
	if err != nil {
		utils.Unauthorized(c)
		return
	}
	var req dto.CreateBookRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Convert the CreateBookRequest to an entities.Book
	book := dto.FormatCreateBookRequest(&req)
	newBook, err := h.bookService.CreateBook(book, user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, dto.FormatCreateBookResponse(*newBook))
}

// @Summary Update Book
// @Description Update an existing book
// @Accept json
// @Produce json
// @Param id path string true "Book ID"
// @Param book body dto.UpdateBookRequest true "Update book"
// @Success 200 {object} entities.Book
// @Failure 400 {object} map[string]interface{}
// @Router /books/{id} [put]
func (h *BookHandler) UpdateBook(c *gin.Context) {
	user, err := utils.GetUserFromContext(c)
	if err != nil {
		utils.Unauthorized(c)
		return
	}

	id := c.Param("id")
	var req dto.UpdateBookRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedBook, err := h.bookService.UpdateBook(id, dto.FormatUpdateBookRequest(&req), user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, updatedBook)
}

// @Summary Delete Book
// @Description Delete an existing book
// @Accept json
// @Produce json
// @Param id path string true "Book ID"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Router /books/{id} [delete]
func (h *BookHandler) DeleteBook(c *gin.Context) {
	user, err := utils.GetUserFromContext(c)
	if err != nil {
		utils.Unauthorized(c)
		return
	}

	id := c.Param("id")
	err = h.bookService.DeleteBook(id, user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, nil)
}
