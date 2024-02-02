package shelf

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/oggnimodd/golang-clerk-practice/feature/shelf/dto"
	"github.com/oggnimodd/golang-clerk-practice/utils"
)

type HandlerShelf struct {
	service ServiceShelfInterface
}

func NewHandlerShelf(service ServiceShelfInterface) HandlerShelfInterface {
	return &HandlerShelf{service: service}
}

// @Summary Get All Shelves
// @Description Retrieve all shelves
// @Produce json
// @Success 200 {array} entities.Shelf
// @Router /shelves [get]
func (h *HandlerShelf) GetAllShelves(c *gin.Context) {
	user, err := utils.GetUserFromContext(c)
	if err != nil {
		utils.Unauthorized(c)
		return
	}

	shelves, err := h.service.GetAllShelves(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, shelves)
}

// @Summary Get Shelf By ID
// @Description Retrieve a shelf by ID
// @Produce json
// @Param id path string true "Shelf ID"
// @Success 200 {array} entities.Shelf
// @Failure 404 {object} map[string]interface{}
// @Router /shelves/{id} [get]
func (h *HandlerShelf) GetShelfById(c *gin.Context) {
	user, err := utils.GetUserFromContext(c)
	if err != nil {
		utils.Unauthorized(c)
		return
	}

	id := c.Param("id")
	shelf, err := h.service.GetShelfById(id, user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, shelf)
}

// @Summary Create Shelf
// @Description Create a new shelf
// @Accept json
// @Produce json
// @Param shelf body dto.CreateShelfRequest true "Create shelf"
// @Success 201 {object} entities.Shelf
// @Failure 400 {object} map[string]interface{}
// @Router /shelves [post]
func (h *HandlerShelf) CreateShelf(c *gin.Context) {
	user, err := utils.GetUserFromContext(c)
	if err != nil {
		utils.Unauthorized(c)
		return
	}

	var req dto.CreateShelfRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	shelf := dto.FormatCreateShelfRequest(&req)
	newShelf, err := h.service.CreateShelf(shelf, user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, dto.FormatCreateShelfResponse(*newShelf))
}

// @Summary Update Shelf
// @Description Update an existing shelf
// @Accept json
// @Produce json
// @Param id path string true "Shelf ID"
// @Param shelf body dto.UpdateShelfRequest true "Update shelf"
// @Success 200 {object} entities.Shelf
// @Failure 400 {object} map[string]interface{}
// @Router /shelves/{id} [put]
func (h *HandlerShelf) UpdateShelf(c *gin.Context) {
	user, err := utils.GetUserFromContext(c)
	if err != nil {
		utils.Unauthorized(c)
		return
	}

	id := c.Param("id")
	var req dto.UpdateShelfRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedShelf, err := h.service.UpdateShelf(id, dto.FormatUpdateShelfRequest(&req), user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, updatedShelf)
}

// @Summary Delete Shelf
// @Description Delete an existing shelf
// @Accept json
// @Produce json
// @Param id path string true "Shelf ID"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Router /shelves/{id} [delete]
func (h *HandlerShelf) DeleteShelf(c *gin.Context) {
	user, err := utils.GetUserFromContext(c)
	if err != nil {
		utils.Unauthorized(c)
		return
	}

	id := c.Param("id")
	err = h.service.DeleteShelf(id, user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, nil)
}

// @Summary Generate Initial Shelves
// @Description Generate initial shelves
// @Produce json
// @Success 200 {array} entities.Shelf
// @Failure 400 {object} map[string]interface{}
// @Router /shelves/initial [get]
func (h *HandlerShelf) GenerateInitialShelves(c *gin.Context) {
	user, err := utils.GetUserFromContext(c)
	if err != nil {
		utils.Unauthorized(c)
		return
	}

	initialShelves, err := h.service.GenerateInitialShelves(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, dto.FormatGenerateShelfResponse(initialShelves))
}
