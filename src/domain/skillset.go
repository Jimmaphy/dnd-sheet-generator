package domain

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
