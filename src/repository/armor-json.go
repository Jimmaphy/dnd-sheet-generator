package repository

import (
	"strings"

	"github.com/jimmaphy/dnd-sheet-generator/domain"
	"github.com/jimmaphy/dnd-sheet-generator/infrastructure"
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

// Add a new armor to the repository
// The armor will be saved as a JSON file named after the armor's name
// The armor's name should be unique to avoid overwriting existing files
func (repository *ArmorJSONRepository) Add(armor *domain.Armor) error {
	jsonService, err := infrastructure.NewJSONService(repository.folder)
	if err != nil {
		return err
	}

	return jsonService.Save(armor.Name, armor)
}

// Get retrieves a armor from the repository by name
// It returns the armor if found, or an error if not found
func (repository *ArmorJSONRepository) Get(name string) (*domain.Armor, error) {
	jsonService, err := infrastructure.NewJSONService(repository.folder)
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

// List retrieves all armors from the repository
// They wil be returned in lower case format
func (repository *ArmorJSONRepository) List() ([]string, error) {
	jsonService, err := infrastructure.NewJSONService(repository.folder)
	if err != nil {
		return nil, err
	}

	var armorNames []string
	armorNames, err = jsonService.List()
	if err != nil {
		return nil, err
	}

	for i, name := range armorNames {
		armorNames[i] = strings.ToLower(name)
	}

	return armorNames, nil
}
