package review

import (
	"github.com/google/uuid"
	"github.com/oggnimodd/golang-clerk-practice/entities"
	"gorm.io/gorm"
)

type ReviewRepository struct {
	db *gorm.DB
}

func NewReviewRepository(db *gorm.DB) RepositoryReviewInterface {
	return &ReviewRepository{db: db}
}

func (r *ReviewRepository) GetReviewByBookId(userId string, bookId string) (entities.Review, error) {
	var review entities.Review
	result := r.db.Where("user_id = ? AND book_id = ?", userId, bookId).First(&review)
	if result.Error != nil {
		return entities.Review{}, result.Error
	}
	return review, nil
}

func (r *ReviewRepository) CreateReview(review *entities.Review, userId string, bookId string) (*entities.Review, error) {
	review.UserID = userId
	review.BookID = &bookId
	review.ID = uuid.New().String()
	result := r.db.Create(review)
	if result.Error != nil {
		return nil, result.Error
	}
	return review, nil
}

func (r *ReviewRepository) UpdateReview(id string, review *entities.Review, userId string) (*entities.Review, error) {
	result := r.db.Model(&entities.Review{}).Where("id = ? AND user_id = ?", id, userId).Updates(review)
	if result.Error != nil {
		return nil, result.Error
	}
	return review, nil
}

func (r *ReviewRepository) DeleteReview(id string, userId string) error {
	result := r.db.Where("id = ? AND user_id = ?", id, userId).Delete(&entities.Review{})
	if result.Error != nil {
		return result.Error
	}
	return nil
}
