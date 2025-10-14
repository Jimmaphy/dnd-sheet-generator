package domain

type Race struct {
	name string
}

// Create a new race based on the provided name.
func NewRace(name string) *Race {
	return &Race{name: name}
}

// AddRace will add the provided race to the character.
// By doing so, it will modify attributes where necessary.
func (character *Character) AddRace(race *Race) error {
	character.Race = race
	return nil
}
