package adapters

import (
	"strings"

	"github.com/jimmaphy/dnd-sheet-generator/domain"
)

type ApiArmor struct {
	Name       string         `json:"name"`
	Category   string         `json:"armor_category"`
	ArmorClass *ApiArmorClass `json:"armor_class"`
}

// ToDomainModel takes the ApiArmor and converts it to the domain model Armor
// The ApiArmor fields are mapped to the corresponding Armor fields received from the API.
func (apiArmor *ApiArmor) ToDomainModel() *domain.Armor {
	armor := &domain.Armor{
		Name:              strings.ToLower(apiArmor.Name),
		Type:              strings.ToLower(apiArmor.Category + " armor"),
		ArmorClass:        apiArmor.ArmorClass.Base,
		DexterityModifier: apiArmor.ArmorClass.DexterityBonus,
		ModifierLimit:     apiArmor.ArmorClass.MaximumBonus,
	}

	return armor
}
