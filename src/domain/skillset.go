package domain

import (
	"fmt"
	"math"
)

type SkillSet struct {
	Strength     int
	Dexterity    int
	Constitution int
	Intelligence int
	Wisdom       int
	Charisma     int
}

// NewSkillSet creates a new SkillSet instance with the given values.
func NewSkillSet(str, dex, con, intl, wis, cha int) *SkillSet {
	return &SkillSet{
		Strength:     str,
		Dexterity:    dex,
		Constitution: con,
		Intelligence: intl,
		Wisdom:       wis,
		Charisma:     cha,
	}
}

// Add the values of another SkillSet to this SkillSet.
func (set *SkillSet) Add(other *SkillSet) {
	set.Strength += other.Strength
	set.Dexterity += other.Dexterity
	set.Constitution += other.Constitution
	set.Intelligence += other.Intelligence
	set.Wisdom += other.Wisdom
	set.Charisma += other.Charisma
}

// Internal function for formatting modifier strings.
func getModifierString(modifier int) string {
	if modifier >= 0 {
		return fmt.Sprintf("+%d", modifier)
	}

	return fmt.Sprintf("%d", modifier)
}

// GetStrengthModifierString returns the strength modifier as a string.
func (set *SkillSet) GetStrengthModifierString() string {
	modifier := int(math.Floor(float64(set.Strength-10) / 2))
	return getModifierString(modifier)
}

// GetDexterityModifierString returns the dexterity modifier as a string.
func (set *SkillSet) GetDexterityModifierString() string {
	modifier := int(math.Floor(float64(set.Dexterity-10) / 2))
	return getModifierString(modifier)
}

// GetConstitutionModifierString returns the constitution modifier as a string.
func (set *SkillSet) GetConstitutionModifierString() string {
	modifier := int(math.Floor(float64(set.Constitution-10) / 2))
	return getModifierString(modifier)
}

// GetIntelligenceModifierString returns the intelligence modifier as a string.
func (set *SkillSet) GetIntelligenceModifierString() string {
	modifier := int(math.Floor(float64(set.Intelligence-10) / 2))
	return getModifierString(modifier)
}

// GetWisdomModifierString returns the wisdom modifier as a string.
func (set *SkillSet) GetWisdomModifierString() string {
	modifier := int(math.Floor(float64(set.Wisdom-10) / 2))
	return getModifierString(modifier)
}

// GetCharismaModifierString returns the charisma modifier as a string.
func (set *SkillSet) GetCharismaModifierString() string {
	modifier := int(math.Floor(float64(set.Charisma-10) / 2))
	return getModifierString(modifier)
}
