package repository

import "github.com/jimmaphy/dnd-sheet-generator/domain"

type ClassRepository interface {
	Get(name string) (*domain.Race, error)
}
