package domain

import "strings"

type SkillSet struct {
	Strength     *Skill
	Dexterity    *Skill
	Constitution *Skill
	Intelligence *Skill
	Wisdom       *Skill
	Charisma     *Skill
}

// NewSkillSet creates a new SkillSet instance with the given values.
// A new instance of Skill is created for each attribute.
func NewSkillSet(str, dex, con, intl, wis, cha int) *SkillSet {
	return &SkillSet{
		Strength:     NewSkill(str),
		Dexterity:    NewSkill(dex),
		Constitution: NewSkill(con),
		Intelligence: NewSkill(intl),
		Wisdom:       NewSkill(wis),
		Charisma:     NewSkill(cha),
	}
}

// Add the values of another SkillSet to the current SkillSet.
func (set *SkillSet) Add(other *SkillSet) {
	set.Strength.Value += other.Strength.Value
	set.Dexterity.Value += other.Dexterity.Value
	set.Constitution.Value += other.Constitution.Value
	set.Intelligence.Value += other.Intelligence.Value
	set.Wisdom.Value += other.Wisdom.Value
	set.Charisma.Value += other.Charisma.Value
}

// GetModifierByName returns the modifier value for the given skill name.
// If the name does not match any skill, it returns 0.
// Capitalization is ignored.
func (set *SkillSet) GetModifierByName(name string) int {
	switch strings.ToLower(name) {
	case "strength":
		return set.Strength.GetModifier()
	case "dexterity":
		return set.Dexterity.GetModifier()
	case "constitution":
		return set.Constitution.GetModifier()
	case "intelligence":
		return set.Intelligence.GetModifier()
	case "wisdom":
		return set.Wisdom.GetModifier()
	case "charisma":
		return set.Charisma.GetModifier()
	default:
		return 0
	}
}
