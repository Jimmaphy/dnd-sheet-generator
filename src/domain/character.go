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
	MainHand    *Weapon
	OffHand     *Weapon
	Armor       *Armor
	Shield      *Shield
	Spells	  	[]*Spell
}

// NewCharacter creates a new Character instance with the given name.
// An empty spell list is initialized.
func NewCharacter(name string, level int) *Character {
	return &Character{
		Name:  name,
		Level: level,
		Spells: []*Spell{},
	}
}

// The equip method equips an item to the character.
// The mainHand parameter indicates whether to equip to the main hand (true) or off hand (false).
func (character *Character) EquipWeapon(weapon *Weapon, mainHand bool) error {
	if mainHand {
		if character.MainHand != nil {
			return errors.New("main hand already occupied")
		}

		character.MainHand = weapon
	} else {
		if character.OffHand != nil {
			return errors.New("off hand already occupied")
		}

		character.OffHand = weapon
	}

	return nil
}

// AddSpell adds a spell to the character's spell list.
// First, a check is performed to ensure the spell can be learned by the class.
func (character *Character) AddSpell(spellName string) error {
	spell, err := character.Class.GetSpell(spellName, character.Level)
	if err != nil {
		return err
	}

	character.Spells = append(character.Spells, spell)
	return nil
}

// The EquipArmor method equips an armor to the character.
func (character *Character) EquipArmor(armor *Armor) error {
	character.Armor = armor
	return nil
}

// The EquipShield method equips a shield to the character.
func (character *Character) EquipShield(shield *Shield) error {
	character.Shield = shield
	return nil
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
// First, the skills are taken from the character's class, based on the skill count.
// Then two additional skills are taken from the character's background.
func (character *Character) GetSkillProficiencyString() (string, error) {
	if character.Class == nil || character.Background == nil {
		return "", errors.New("cannot get skill proficiencies without class or background")
	}

	classSkillCount := character.Class.SkillCount
	proficiencies := []string{}
	proficiencies = append(proficiencies, character.Class.Skills[:classSkillCount]...)
	proficiencies = append(proficiencies, character.Background.Skills[:2]...)

	return strings.Join(proficiencies, ", "), nil
}
