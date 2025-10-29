package repository

import "github.com/jimmaphy/dnd-sheet-generator/domain"

type ClassRepository interface {
	Add(class *domain.Class) error
	Get(name string) (*domain.Class, error)
	List() ([]string, error)
}
