package review

import (
	"github.com/gin-gonic/gin"
	"github.com/oggnimodd/golang-clerk-practice/entities"
)

type RepositoryReviewInterface interface {
	GetReviewByBookId(userId string, bookId string) (entities.Review, error)
	CreateReview(review *entities.Review, userId string, bookId string) (*entities.Review, error)
	UpdateReview(id string, review *entities.Review, userId string) (*entities.Review, error)
	DeleteReview(id string, userId string) error
}

type ServiceReviewInterface interface {
	GetReviewByBookId(userId string, bookId string) (entities.Review, error)
	CreateReview(review *entities.Review, userId string, bookId string) (*entities.Review, error)
	UpdateReview(id string, review *entities.Review, userId string) (*entities.Review, error)
	DeleteReview(id string, userId string) error
}

type HandlerReviewInterface interface {
	GetReviewByBookId(c *gin.Context)
	CreateReview(c *gin.Context)
	UpdateReview(c *gin.Context)
	DeleteReview(c *gin.Context)
}
