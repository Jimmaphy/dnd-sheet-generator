package repository

import (
	"github.com/jimmaphy/dnd-sheet-generator/domain"
	"github.com/jimmaphy/dnd-sheet-generator/services"
)

type CharacterJSONRepository struct {
	folder string
}

// Create a new instance of the JSON repository for characters
// The folder for storage will be automatically set to "characters"
func NewCharacterJSONRepository() *CharacterJSONRepository {
	return &CharacterJSONRepository{
		folder: "characters",
	}
}

// Add a new character to the repository
// The character will be saved as a JSON file named after the character's name
// The character's name should be unique to avoid overwriting existing files
func (repository *CharacterJSONRepository) Add(character *domain.Character) error {
	jsonService, err := services.NewJSONService(repository.folder)
	if err != nil {
		return err
	}

	return jsonService.Save(character.Name, character)
}

// List retrieves the names of all characters stored in the repository
func (repository *CharacterJSONRepository) List() ([]string, error) {
	jsonService, err := services.NewJSONService(repository.folder)
	if err != nil {
		return nil, err
	}

	var characterNames []string
	characterNames, err = jsonService.List()
	if err != nil {
		return nil, err
	}

	return characterNames, nil
}
