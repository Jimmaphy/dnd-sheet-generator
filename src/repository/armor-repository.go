package repository

import "github.com/jimmaphy/dnd-sheet-generator/domain"

type ArmorRepository interface {
	Get(name string) (*domain.Armor, error)
}
