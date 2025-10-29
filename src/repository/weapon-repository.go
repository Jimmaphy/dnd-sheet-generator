package repository

import "github.com/jimmaphy/dnd-sheet-generator/domain"

type WeaponRepository interface {
	Add(weapon *domain.Weapon) error
	Get(name string) (*domain.Weapon, error)
	List() ([]string, error)
}
