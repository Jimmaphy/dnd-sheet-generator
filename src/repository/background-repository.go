package repository

import "github.com/jimmaphy/dnd-sheet-generator/domain"

type BackgroundRepository interface {
	Get(name string) (*domain.Background, error)
}
