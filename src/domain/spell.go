package domain

type Spell struct {
	Name string
	Level int
}

// Create a new spell instance with a name.
func NewSpell(name string) *Spell {
	return &Spell{Name: name}
}
