package domain

import (
	"errors"
	"sort"
	"strconv"
	"strings"
)

type Class struct {
	Name        string
	SkillCount  int
	Skills      []string
	CasterType  string
	SpellLevels []SpellLevel
	Spells      []*Spell
}

// Create a new class based on the provided name.
// The skills are not initialized here and can be set later.
func NewClass(name string) *Class {
	return &Class{
		Name: name,
	}
}

// The IsLegalSpell method checks if a spell can be learned or prepared by the class.
// It returns nill if the spell is legal, or an error otherwise.
// First, it checks whether the spell is in the list of spells.
// Then, based on the provided level, it checks if the spell can be learned at that level.
func (class *Class) GetSpell(spellName string, level int) (*Spell, error) {
	for _, targetSpell := range class.Spells {
		if targetSpell.Name == spellName {
			if class.GetSpellSlots(level, targetSpell.Level) == 0 {
				return nil, errors.New("the spell has higher level than the available spell slots")
			}

			return targetSpell, nil
		}
	}

	return nil, errors.New("the spell is not available for this class")
}

// CanCastsSpells checks if the class can cast spells.
// The function returns false if the CasterType is "none", true otherwise.
func (class *Class) CanCastSpells() bool {
	return class.CasterType != "none"
}

// GetSpellSlots returns the spell slots available for a given spell level.
// For a pact caster, it returns the slots for the highest spell level available at the character's level.
// For other casters, it returns the slots for the specified spell level.
func (class *Class) GetSpellSlots(level int, spellLevel int) int {
	for _, spellLevelData := range class.SpellLevels {
		if spellLevelData.Level == level {
			for slotLevel, slots := range spellLevelData.Slots {
				convertedSlotLevel, _ := strconv.Atoi(slotLevel)
				if class.CasterType == "pact" && convertedSlotLevel >= spellLevel {
					return slots
				} else if convertedSlotLevel == spellLevel {
					return slots
				}
			}
		}
	}

	return 0
}

// The SpellSlotsString method returns a string representation of the spell slots for a given level.
// A prefix of two spaces is added for formatting purposes.
// The format is "Level X: Y" for each spell level.
// Where X is the spell level and Y is the number of slots available.
// The levels will be order in ascending order.
func (class *Class) SpellSlotsString(level int) string {
	var slotsStrings []string

	for _, spellLevelData := range class.SpellLevels {
		if spellLevelData.Level == level {
			for slotLevel, slots := range spellLevelData.Slots {
				slotsStrings = append(slotsStrings, "  Level "+slotLevel+": "+strconv.Itoa(slots))
			}
		}
	}

	sort.Strings(slotsStrings)
	return strings.Join(slotsStrings, "\n")
}
