package domain

type Shield struct {
	Name string
}

// Create a new instance of Shield.
func NewShield(name string) *Shield {
	return &Shield{Name: name}
}
