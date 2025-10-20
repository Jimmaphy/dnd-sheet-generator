package repository

import (
	"github.com/jimmaphy/dnd-sheet-generator/domain"
	"github.com/jimmaphy/dnd-sheet-generator/services"
)

type RaceJSONRepository struct {
	folder string
}

// Create a new instance of the JSON repository for races
// The folder for storage will be automatically set to "races"
func NewRaceJSONRepository() *RaceJSONRepository {
	return &RaceJSONRepository{
		folder: "races",
	}
}

// Get retrieves a race from the repository by name
// It returns the race if found, or an error if not found
func (repository *RaceJSONRepository) Get(name string) (*domain.Race, error) {
	jsonService, err := services.NewJSONService(repository.folder)
	if err != nil {
		return nil, err
	}

	var race domain.Race
	err = jsonService.ReadCaseInsensitive(name, &race)
	if err != nil {
		return nil, err
	}

	return &race, nil
}
