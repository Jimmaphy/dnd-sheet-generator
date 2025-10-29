package adapters

type ApiSpellList struct {
	Count   int         `json:"count"`
	Results []*ApiSpell `json:"results"`
}
