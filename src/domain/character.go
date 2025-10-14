package domain

type Character struct {
	Name string
	Race *Race
}

// NewCharacter creates a new Character instance with the given name.
func NewCharacter(name string) *Character {
	return &Character{Name: name}
}

// The SetRace method assigns a race to the character.
func (character *Character) SetRace(race *Race) {
	character.Race = race
}
