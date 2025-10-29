package adapters

type ApiSpellCasting struct {
	Ability Listing   `json:"spellcasting_ability"`
	Info    []ApiInfo `json:"info"`
}
