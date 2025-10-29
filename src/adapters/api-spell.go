package adapters

import "github.com/jimmaphy/dnd-sheet-generator/domain"

type ApiSpell struct {
	Index string `json:"index"`
	Name  string `json:"name"`
	Level int    `json:"level"`
	Url   string `json:"url"`
}

// ToDomainModel converts the ApiSpell to the domain model Spell.
func (apiSpell *ApiSpell) ToDomainModel() *domain.Spell {
	return &domain.Spell{
		Name:  apiSpell.Name,
		Level: apiSpell.Level,
	}
}
