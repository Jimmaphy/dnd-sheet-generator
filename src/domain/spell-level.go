package domain

type SpellLevel struct {
	Level       int
	SpellsKnown int
	Slots       map[string]int
}

// NewSpellLevel creates a new SpellLevel instance.
func NewSpellLevel(level int, spellsKnown int, slots map[string]int) *SpellLevel {
	return &SpellLevel{
		Level:       level,
		SpellsKnown: spellsKnown,
		Slots:       slots,
	}
}
