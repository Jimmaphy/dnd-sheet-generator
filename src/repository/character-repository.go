package repository

import "github.com/jimmaphy/dnd-sheet-generator/domain"

type CharacterRepository interface {
	Add(character *domain.Character) error
	List() ([]string, error)
	Delete(name string) error
	Get(name string) (*domain.Character, error)
}
