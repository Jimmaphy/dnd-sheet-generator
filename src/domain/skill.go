package domain

import (
	"math"
	"strconv"
)

type Skill struct {
	Value int
}

// Initialize a skill based on a name and a value.
func NewSkill(value int) *Skill {
	return &Skill{
		Value: value,
	}
}

// GetModifier calculates the skill modifier based on the skill value.
func (skill *Skill) GetModifier() int {
	return int(math.Floor(float64(skill.Value-10) / 2))
}

// GetModifierString returns the skill modifier as a formatted string.
// A '+' sign is included for positive modifiers.
func (skill *Skill) GetModifierString() string {
	modifier := skill.GetModifier()

	if modifier >= 0 {
		return "+" + strconv.Itoa(modifier)
	}

	return strconv.Itoa(modifier)
}
