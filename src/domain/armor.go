package domain

type Armor struct {
	Name              string
	Type              string
	ArmorClass        int
	DexterityModifier bool
	ModifierLimit     int
}

// Create a new instance of Armor.
func NewArmor(name string) *Armor {
	return &Armor{
		Name:              name,
		Type:              "Light",
		ArmorClass:        11,
		DexterityModifier: true,
		ModifierLimit:     0,
	}
}

// The GetArmorClass method returns the armor class of the armor.
// It takes the dexterity modifier as an argument and applies it if allowed.
func (armor *Armor) GetArmorClass(dexModifier int) int {
	armorClass := armor.ArmorClass
	dexterityBonus := 0

	if armor.DexterityModifier {
		dexterityBonus = dexModifier
	}

	if armor.ModifierLimit > 0 && dexterityBonus > armor.ModifierLimit {
		dexterityBonus = armor.ModifierLimit
	}

	return armorClass + dexterityBonus
}
