package domain

type Armor struct {
	Name string
}

// Create a new instance of Armor.
func NewArmor(name string) *Armor {
	return &Armor{Name: name}
}
