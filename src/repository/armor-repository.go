package repository

import "github.com/jimmaphy/dnd-sheet-generator/domain"

type ArmorRepository interface {
	Add(armor *domain.Armor) error
	Get(name string) (*domain.Armor, error)
	List() ([]string, error)
}
