package adapters

type ApiSpellLevel struct {
	Cantrips        int `json:"cantrips_known"`
	SpellsAvailable int `json:"spells_known"`
	Level1          int `json:"spell_slots_level_1"`
	Level2          int `json:"spell_slots_level_2"`
	Level3          int `json:"spell_slots_level_3"`
	Level4          int `json:"spell_slots_level_4"`
	Level5          int `json:"spell_slots_level_5"`
	Level6          int `json:"spell_slots_level_6"`
	Level7          int `json:"spell_slots_level_7"`
	Level8          int `json:"spell_slots_level_8"`
	Level9          int `json:"spell_slots_level_9"`
}
