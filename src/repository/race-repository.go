package repository

import "github.com/jimmaphy/dnd-sheet-generator/domain"

type RaceRepository interface {
	Get(name string) (*domain.Race, error)
}
