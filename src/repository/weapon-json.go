package repository

import (
	"strings"

	"github.com/jimmaphy/dnd-sheet-generator/domain"
	"github.com/jimmaphy/dnd-sheet-generator/infrastructure"
)

type WeaponJSONRepository struct {
	folder string
}

// Create a new instance of the JSON repository for armores
// The folder for storage will be automatically set to "armores"
func NewWeaponJSONRepository() *WeaponJSONRepository {
	return &WeaponJSONRepository{
		folder: "weapons",
	}
}

// Add a new armor to the repository
// The armor will be saved as a JSON file named after the armor's name
// The armor's name should be unique to avoid overwriting existing files
func (repository *WeaponJSONRepository) Add(weapon *domain.Weapon) error {
	jsonService, err := infrastructure.NewJSONService(repository.folder)
	if err != nil {
		return err
	}

	return jsonService.Save(weapon.Name, weapon)
}

// Get retrieves a armor from the repository by name
// It returns the armor if found, or an error if not found
func (repository *WeaponJSONRepository) Get(name string) (*domain.Weapon, error) {
	jsonService, err := infrastructure.NewJSONService(repository.folder)
	if err != nil {
		return nil, err
	}

	var armor domain.Weapon
	err = jsonService.ReadCaseInsensitive(name, &armor)
	if err != nil {
		return nil, err
	}

	return &armor, nil
}

// List retrieves all armors from the repository
// They wil be returned in lower case format
func (repository *WeaponJSONRepository) List() ([]string, error) {
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
