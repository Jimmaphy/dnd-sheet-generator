package domain

type Shield struct {
	Name               string
	Type               string
	ArmorClassAddition int
}

// Create a new instance of Shield.
func NewShield(name string) *Shield {
	return &Shield{
		Name:               name,
		Type:               "shield",
		ArmorClassAddition: 2,
	}
}

// GetArmorClassAddition returns the armor class addition provided by the shield.
func (shield *Shield) GetArmorClassAddition() int {
	return shield.ArmorClassAddition
}
