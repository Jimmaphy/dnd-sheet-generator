package repository

import (
	"github.com/jimmaphy/dnd-sheet-generator/domain"
	"github.com/jimmaphy/dnd-sheet-generator/services"
)

type ClassJSONRepository struct {
	folder string
}

// Create a new instance of the JSON repository for classes
// The folder for storage will be automatically set to "classes"
func NewClassJSONRepository() *ClassJSONRepository {
	return &ClassJSONRepository{
		folder: "classes",
	}
}

// Get retrieves a class from the repository by name
// It returns the class if found, or an error if not found
func (repository *ClassJSONRepository) Get(name string) (*domain.Class, error) {
	jsonService, err := services.NewJSONService(repository.folder)
	if err != nil {
		return nil, err
	}

	var class domain.Class
	err = jsonService.ReadCaseInsensitive(name, &class)
	if err != nil {
		return nil, err
	}

	return &class, nil
}
