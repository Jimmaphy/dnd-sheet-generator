package domain

type Background struct {
	Name   string
	Skills []string
}

// Create a new background based on the provided name.
// The skills are not initialized here and can be set later.
func NewBackground(name string) *Background {
	return &Background{
		Name: name,
	}
}
