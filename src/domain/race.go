package domain

type Race struct {
	Name           string
	SkillModifiers *SkillSet
}

// Create a new race based on the provided name.
func NewRace(name string) *Race {
	return &Race{
		Name: name,
	}
}
