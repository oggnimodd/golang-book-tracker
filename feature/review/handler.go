package review

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/oggnimodd/golang-clerk-practice/entities"
	"github.com/oggnimodd/golang-clerk-practice/utils"
)

type ReviewHandler struct {
	reviewService ServiceReviewInterface
}

func NewReviewHandler(reviewService ServiceReviewInterface) HandlerReviewInterface {
	return &ReviewHandler{
		reviewService: reviewService,
	}
}

// @Summary Get Review By Book Id
// @Description Retrieve a review by book ID
// @Produce json
// @Param book_id path string true "Book ID"
// @Success 200 {object} entities.Review
// @Failure 500 {object} map[string]interface{}
// @Router /reviews/{book_id} [get]
func (h *ReviewHandler) GetReviewByBookId(c *gin.Context) {
	user, err := utils.GetUserFromContext(c)
	if err != nil {
		utils.Unauthorized(c)
		return
	}

	bookId := c.Param("book_id")
	review, err := h.reviewService.GetReviewByBookId(user.ID, bookId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, review)
}

// @Summary Create Review
// @Description Create a new review
// @Accept json
// @Produce json
// @Param book_id path string true "Book ID"
// @Param review body entities.Review true "Create review"
// @Success 201 {object} entities.Review
// @Failure 400 {object} map[string]interface{}
// @Router /reviews/{book_id} [post]
func (h *ReviewHandler) CreateReview(c *gin.Context) {
	user, err := utils.GetUserFromContext(c)
	if err != nil {
		utils.Unauthorized(c)
		return
	}

	var review entities.Review
	if err := c.ShouldBindJSON(&review); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	bookId := c.Param("book_id")
	newReview, err := h.reviewService.CreateReview(&review, user.ID, bookId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, newReview)
}

// @Summary Update Review
// @Description Update an existing review
// @Accept json
// @Produce json
// @Param id path string true "Review ID"
// @Param review body entities.Review true "Update review"
// @Success 200 {object} entities.Review
// @Failure 400 {object} map[string]interface{}
// @Router /reviews/{id} [put]
func (h *ReviewHandler) UpdateReview(c *gin.Context) {
	user, err := utils.GetUserFromContext(c)
	if err != nil {
		utils.Unauthorized(c)
		return
	}

	id := c.Param("id")
	var updateReview entities.Review
	if err := c.ShouldBindJSON(&updateReview); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedReview, err := h.reviewService.UpdateReview(id, &updateReview, user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, updatedReview)
}

// @Summary Delete Review
// @Description Delete an existing review
// @Accept json
// @Produce json
// @Param id path string true "Review ID"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Router /reviews/{id} [delete]
func (h *ReviewHandler) DeleteReview(c *gin.Context) {
	user, err := utils.GetUserFromContext(c)
	if err != nil {
		utils.Unauthorized(c)
		return
	}

	id := c.Param("id")
	err = h.reviewService.DeleteReview(id, user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, nil)
}
