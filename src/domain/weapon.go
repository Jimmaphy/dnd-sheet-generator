package domain

type Weapon struct {
	Name       string
	Category   string
	Type       string
	DamageDice string
	DamageType string
}

// Create a new instance of Weapon.
func NewWeapon(name string) *Weapon {
	return &Weapon{Name: name}
}

// GetName returns the name of the weapon.
func (weapon *Weapon) GetName() string {
	return weapon.Name
}
