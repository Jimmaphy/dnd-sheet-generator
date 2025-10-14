package domain

type Character struct {
	Name string
}

// NewCharacter creates a new Character instance with the given name.
func NewCharacter(name string) *Character {
	return &Character{Name: name}
}
