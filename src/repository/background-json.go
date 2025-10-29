package repository

import (
	"github.com/jimmaphy/dnd-sheet-generator/domain"
	"github.com/jimmaphy/dnd-sheet-generator/infrastructure"
)

type BackgroundJSONRepository struct {
	folder string
}

// Create a new instance of the JSON repository for backgrounds
// The folder for storage will be automatically set to "backgrounds"
func NewBackgroundJSONRepository() *BackgroundJSONRepository {
	return &BackgroundJSONRepository{
		folder: "backgrounds",
	}
}

// Get retrieves a background from the repository by name
// It returns the background if found, or an error if not found
func (repository *BackgroundJSONRepository) Get(name string) (*domain.Background, error) {
	jsonService, err := infrastructure.NewJSONService(repository.folder)
	if err != nil {
		return nil, err
	}

	var background domain.Background
	err = jsonService.ReadCaseInsensitive(name, &background)
	if err != nil {
		return nil, err
	}

	return &background, nil
}
