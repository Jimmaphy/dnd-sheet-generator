package domain

type Class struct {
	Name       string
	SkillCount int
	Skills     []string
	CasterType string
}

// Create a new class based on the provided name.
// The skills are not initialized here and can be set later.
func NewClass(name string) *Class {
	return &Class{
		Name: name,
	}
}
