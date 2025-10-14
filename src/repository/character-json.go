package repository

import (
	"strings"

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
// The character will be saved as a JSON file named after the character's name in lowercase
// The character's name should be unique to avoid overwriting existing files
func (repository *CharacterJSONRepository) Add(character *domain.Character) error {
	jsonService, err := services.NewJSONService(repository.folder, strings.ToLower(character.Name))
	if err != nil {
		return err
	}

	return jsonService.Save(character)
}
