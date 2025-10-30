package adapters

type ApiWeaponDamage struct {
	DamageDice string  `json:"damage_dice"`
	DamageType Listing `json:"damage_type"`
}
