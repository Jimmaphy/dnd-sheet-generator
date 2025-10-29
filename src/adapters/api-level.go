package adapters

import "github.com/jimmaphy/dnd-sheet-generator/domain"

type ApiLevel struct {
	Level        int           `json:"level"`
	SpellCasting ApiSpellLevel `json:"spellcasting"`
}

// ToDomainModel converts the ApiLevel to the domain model SpellLevel.
func (apiLevel *ApiLevel) ToDomainModel() *domain.SpellLevel {
	slots := map[string]int{
		"0": apiLevel.SpellCasting.Cantrips,
		"1": apiLevel.SpellCasting.Level1,
		"2": apiLevel.SpellCasting.Level2,
		"3": apiLevel.SpellCasting.Level3,
		"4": apiLevel.SpellCasting.Level4,
		"5": apiLevel.SpellCasting.Level5,
		"6": apiLevel.SpellCasting.Level6,
		"7": apiLevel.SpellCasting.Level7,
		"8": apiLevel.SpellCasting.Level8,
		"9": apiLevel.SpellCasting.Level9,
	}

	for level, count := range slots {
		if count == 0 {
			delete(slots, level)
		}
	}

	return &domain.SpellLevel{
		Level:       apiLevel.Level,
		SpellsKnown: apiLevel.SpellCasting.SpellsAvailable,
		Slots:       slots,
	}
}
