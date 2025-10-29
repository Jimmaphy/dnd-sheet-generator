package repository

import (
	"strings"

	"github.com/jimmaphy/dnd-sheet-generator/domain"
	"github.com/jimmaphy/dnd-sheet-generator/infrastructure"
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

// Add a new class to the repository
// The class will be saved as a JSON file named after the class's name
// The class's name should be unique to avoid overwriting existing files
func (repository *ClassJSONRepository) Add(class *domain.Class) error {
	jsonService, err := infrastructure.NewJSONService(repository.folder)
	if err != nil {
		return err
	}

	return jsonService.Save(class.Name, class)
}

// Get retrieves a class from the repository by name
// It returns the class if found, or an error if not found
func (repository *ClassJSONRepository) Get(name string) (*domain.Class, error) {
	jsonService, err := infrastructure.NewJSONService(repository.folder)
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

// List retrieves all classes from the repository
// They wil be returned in lower case format
func (repository *ClassJSONRepository) List() ([]string, error) {
	jsonService, err := infrastructure.NewJSONService(repository.folder)
	if err != nil {
		return nil, err
	}

	var classNames []string
	classNames, err = jsonService.List()
	if err != nil {
		return nil, err
	}

	for i, name := range classNames {
		classNames[i] = strings.ToLower(name)
	}

	return classNames, nil
}
