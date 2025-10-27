package repository

import (
	"github.com/jimmaphy/dnd-sheet-generator/domain"
	"github.com/jimmaphy/dnd-sheet-generator/services"
)

type ArmorJSONRepository struct {
	folder string
}

// Create a new instance of the JSON repository for armores
// The folder for storage will be automatically set to "armores"
func NewArmorJSONRepository() *ArmorJSONRepository {
	return &ArmorJSONRepository{
		folder: "armors",
	}
}

// Get retrieves a armor from the repository by name
// It returns the armor if found, or an error if not found
func (repository *ArmorJSONRepository) Get(name string) (*domain.Armor, error) {
	jsonService, err := services.NewJSONService(repository.folder)
	if err != nil {
		return nil, err
	}

	var armor domain.Armor
	err = jsonService.ReadCaseInsensitive(name, &armor)
	if err != nil {
		return nil, err
	}

	return &armor, nil
}
