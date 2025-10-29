package adapters

type ApiProficiencyChoices struct {
	Choose int               `json:"choose"`
	From   []*ApiProficiency `json:"from"`
}
