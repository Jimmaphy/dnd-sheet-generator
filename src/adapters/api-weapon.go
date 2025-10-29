package adapters

import (
	"strings"

	"github.com/jimmaphy/dnd-sheet-generator/domain"
)

type ApiWeapon struct {
	Name     string           `json:"name"`
	Category string           `json:"weapon_category"`
	Range    string           `json:"weapon_range"`
	Damage   *ApiWeaponDamage `json:"damage"`
}

// ToDomainModel takes the ApiWeapon and converts it to the domain model Weapon
// The ApiWeapon fields are mapped to the corresponding Weapon fields received from the API.
func (apiWeapon *ApiWeapon) ToDomainModel() *domain.Weapon {
	return &domain.Weapon{
		Name:       strings.ToLower(apiWeapon.Name),
		Category:   strings.ToLower(apiWeapon.Category),
		Type:       strings.ToLower(apiWeapon.Range),
		DamageDice: strings.ToLower(apiWeapon.Damage.DamageDice),
		DamageType: strings.ToLower(apiWeapon.Damage.DamageType.Name),
	}
}
