package review

import "github.com/oggnimodd/golang-clerk-practice/entities"

type ReviewService struct {
	reviewRepo RepositoryReviewInterface
}

func NewReviewService(reviewRepo RepositoryReviewInterface) ServiceReviewInterface {
	return &ReviewService{
		reviewRepo: reviewRepo,
	}
}

func (s *ReviewService) GetReviewByBookId(userId string, bookId string) (entities.Review, error) {
	return s.reviewRepo.GetReviewByBookId(userId, bookId)
}

func (s *ReviewService) CreateReview(review *entities.Review, userId string, bookId string) (*entities.Review, error) {
	return s.reviewRepo.CreateReview(review, userId, bookId)
}

func (s *ReviewService) UpdateReview(id string, review *entities.Review, userId string) (*entities.Review, error) {
	return s.reviewRepo.UpdateReview(id, review, userId)
}

func (s *ReviewService) DeleteReview(id string, userId string) error {
	return s.reviewRepo.DeleteReview(id, userId)
}
