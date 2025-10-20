package domain

import (
	"errors"
	"fmt"
	"math"
	"strings"
)

type Character struct {
	Name        string
	Race        *Race
	Class       *Class
	Background  *Background
	Level       int
	BaseSkills  *SkillSet
	TotalSkills *SkillSet
}

// NewCharacter creates a new Character instance with the given name.
func NewCharacter(name string, level int) *Character {
	return &Character{
		Name:  name,
		Level: level,
	}
}

// The SetRace method assigns a race to the character.
func (character *Character) SetRace(race *Race) {
	character.Race = race
}

// The SetClass method assigns a class to the character.
func (character *Character) SetClass(class *Class) {
	character.Class = class
}

// The SetBackground method assigns a background to the character.
func (character *Character) SetBackground(background *Background) {
	character.Background = background
}

// The SetSkillSet method assigns a skill set to the character.
func (character *Character) SetSkillSet(skillSet *SkillSet) {
	character.BaseSkills = skillSet
}

// GetTotalStrength returns the total strength score of the character,
// combining base skills with racial modifiers.
func (character *Character) CalculateTotalSkills() error {
	if character.BaseSkills == nil || character.Race == nil {
		return errors.New("cannot calculate total skill point without base skills or race")
	}

	character.TotalSkills = NewSkillSet(0, 0, 0, 0, 0, 0)
	character.TotalSkills.Add(character.BaseSkills)
	character.TotalSkills.Add(character.Race.SkillModifiers)

	return nil
}

// GetProficiencyBonus calculates the proficiency bonus based on the character's level.
// The returned string includes a '+' sign.
func (character *Character) GetProficiencyBonus() string {
	bonus := 1 + int(math.Ceil(float64(character.Level)/4))
	return fmt.Sprintf("+%d", bonus)
}

// GetSkillProficiencyString returns the skill proficiency string.
// The skills are represented as a comma-separated list.
// The first two skills come from the class, and the next two from the background.
func (character *Character) GetSkillProficiencyString() (string, error) {
	if character.Class == nil || character.Background == nil {
		return "", errors.New("cannot get skill proficiencies without class or background")
	}

	proficiencies := []string{}
	proficiencies = append(proficiencies, character.Class.Skills[:2]...)
	proficiencies = append(proficiencies, character.Background.Skills[:2]...)

	return strings.Join(proficiencies, ", "), nil
}
