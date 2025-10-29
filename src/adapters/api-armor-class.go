package adapters

type ApiArmorClass struct {
	Base           int  `json:"base"`
	DexterityBonus bool `json:"dex_bonus"`
	MaximumBonus   int  `json:"max_bonus"`
}
