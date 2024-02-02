package shelf

import (
	"errors"

	"github.com/google/uuid"
	"github.com/oggnimodd/golang-clerk-practice/entities"
	"github.com/oggnimodd/golang-clerk-practice/utils"
)

type ServiceShelf struct {
	repo RepositoryShelfInterface
}

func NewServiceShelf(repo RepositoryShelfInterface) ServiceShelfInterface {
	return &ServiceShelf{repo: repo}
}

func (s *ServiceShelf) GetAllShelves(userId string) ([]entities.Shelf, error) {
	return s.repo.GetAllShelves(userId)
}

func (s *ServiceShelf) GetShelfById(id string, userId string) (entities.Shelf, error) {
	return s.repo.GetShelfById(id, userId)
}

func (s *ServiceShelf) CreateShelf(shelf *entities.Shelf, userId string) (*entities.Shelf, error) {
	return s.repo.CreateShelf(shelf, userId)
}

func (s *ServiceShelf) UpdateShelf(id string, shelf *entities.Shelf, userId string) (*entities.Shelf, error) {
	return s.repo.UpdateShelf(id, shelf, userId)
}

func (s *ServiceShelf) DeleteShelf(id string, userId string) error {
	return s.repo.DeleteShelf(id, userId)
}

var INITIAL_SHELVES = []entities.Shelf{
	{
		Type:      "To Be Read",
		IsDefault: utils.BoolPtr(true),
	},
	{
		Type:      "Currently Reading",
		IsDefault: utils.BoolPtr(true),
	},
	{
		Type:      "Read",
		IsDefault: utils.BoolPtr(true),
	},
}

func (s *ServiceShelf) GenerateInitialShelves(userId string) ([]entities.Shelf, error) {
	existingShelves, err := s.repo.GetAllShelves(userId)
	if err != nil {
		return nil, err
	}

	if len(existingShelves) > 0 {
		return nil, errors.New("shelves are already initialized")
	}

	var initialShelves []entities.Shelf
	for _, shelf := range INITIAL_SHELVES {
		shelf.UserID = userId
		shelf.ID = uuid.New().String() // Assign a new UUID to the shelf
		newShelf, err := s.repo.CreateShelf(&shelf, userId)
		if err != nil {
			return nil, err
		}
		initialShelves = append(initialShelves, *newShelf)
	}

	return initialShelves, nil
}
