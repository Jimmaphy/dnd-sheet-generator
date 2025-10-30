package domain

import (
	"errors"
	"fmt"
	"math"
	"slices"
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
	Spells      []Spell
}

// NewCharacter creates a new Character instance with the given name.
// An empty spell list is initialized.
func NewCharacter(name string, level int) *Character {
	return &Character{
		Name:   name,
		Level:  level,
		Spells: []Spell{},
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

	character.Spells = append(character.Spells, *spell)
	return nil
}

// GetHitPoints calculates the total hit points of the character.
// Base + (die-average * (level - 1)) + (constitution modifier * level)
func (character *Character) GetHitPoints() int {
	base := character.Class.HitDie
	dice := character.Class.GetHitDieAverage() * (character.Level - 1)
	constitution := character.TotalSkills.Constitution.GetModifier() * character.Level

	return base + dice + constitution
}

// GetSpellStrings combines alle spell names into a single string,
// These are the spells the character has learned or prepared.
func (character *Character) GetSpellString() string {
	spellNames := []string{}

	for _, spell := range character.Spells {
		spellNames = append(spellNames, spell.Name)
	}

	return strings.Join(spellNames, ", ")
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
func (character *Character) GetProficiencyBonus() int {
	return 1 + int(math.Ceil(float64(character.Level)/4))
}

// GetProficiencyBonus calculates the proficiency bonus based on the character's level.
// The returned string includes a '+' sign.
func (character *Character) GetProficiencyBonusString() string {
	return fmt.Sprintf("+%d", character.GetProficiencyBonus())
}

// GetInitiativeBonus returns the initiative bonus of the character.
// This is equal to the dexterity modifier.
func (character *Character) GetInitiativeBonus() int {
	return character.TotalSkills.Dexterity.GetModifier()
}

// GetPassivePerception calculates the passive perception of the character.
// This is equal to 10 + wisdom modifier
func (character *Character) GetPassivePerception() int {
	wisdomModifier := character.TotalSkills.Wisdom.GetModifier()
	return 10 + wisdomModifier
}

// IsProficient checks if the character is proficient in the given skill.
// Returns true if proficient, false otherwise.
func (character *Character) IsProficient(skillName string) bool {
	return slices.Contains(character.GetSkillProficiency(), skillName)
}

// GetSkillBonusString returns the skill bonus string for a given skill name.
// The function that the main skill name, like dexterity, and returns the modifier as a string with a '+' sign.
// As second parameter, it takes the name of the skill to check for proficiency.
func (character *Character) GetSkillBonusString(mainSkill string, skillName string) string {
	modifier := character.TotalSkills.GetModifierByName(mainSkill)

	if character.IsProficient(skillName) {
		modifier += character.GetProficiencyBonus()
	}

	return fmt.Sprintf("+%d", modifier)
}

// GetSpellSaveDC calculates the spell save DC for the character.
// The formula is 8 + proficiency bonus + spellcasting ability modifier.
// If the character has no class, 0 is returned.
func (character *Character) GetSpellSaveDC() int {
	castAbility := character.Class.CastAbility
	castModifier := character.TotalSkills.GetModifierByName(castAbility)
	spellSaveDC := 8 + character.GetProficiencyBonus() + castModifier

	return spellSaveDC
}

// GetSpellAttackBonusString calculates the spell attack bonus for the character.
// The formula is proficiency bonus + spellcasting ability modifier.
// The returned string includes a '+' sign.
// If the character has no class, an empty string is returned.
func (character *Character) GetSpellAttackBonusString() string {
	castAbility := character.Class.CastAbility
	castModifier := character.TotalSkills.GetModifierByName(castAbility)
	spellAttackBonus := character.GetProficiencyBonus() + castModifier

	return fmt.Sprintf("+%d", spellAttackBonus)
}

// GetSkillProficiency returns a slice of skill names that the character is proficient in.
// The skills are taken from the character's class, based on the skill count.
// Then two additional skills are taken from the character's background.
func (character *Character) GetSkillProficiency() []string {
	classSkillCount := character.Class.SkillCount
	proficiencies := []string{}
	proficiencies = append(proficiencies, character.Class.Skills[:classSkillCount]...)
	proficiencies = append(proficiencies, character.Background.Skills[:2]...)

	return proficiencies
}

// GetSkillProficiencyString returns the skill proficiency string.
// The skills are represented as a comma-separated list.
// First, the skills are taken from the character's class, based on the skill count.
// Then two additional skills are taken from the character's background.
func (character *Character) GetSkillProficiencyString() (string, error) {
	classSkillCount := character.Class.SkillCount
	proficiencies := []string{}
	proficiencies = append(proficiencies, character.Class.Skills[:classSkillCount]...)
	proficiencies = append(proficiencies, character.Background.Skills[:2]...)

	return strings.Join(proficiencies, ", "), nil
}

// GetArmorClass calculates the total armor class of the character.
// It considers the armor, shield, and dexterity modifier.
func (character *Character) GetArmorClass() int {
	dexterityModifier := character.TotalSkills.Dexterity.GetModifier()
	unarmoredModifiers := character.Class.UnarmoredDefenseModifiers
	armorClass := 10

	if character.Armor != nil {
		armorClass = character.Armor.GetArmorClass(dexterityModifier)
	} else if len(unarmoredModifiers) > 0 {
		for _, modifier := range unarmoredModifiers {
			armorClass += character.TotalSkills.GetModifierByName(modifier)
		}
	} else {
		armorClass += dexterityModifier
	}

	if character.Shield != nil {
		armorClass += character.Shield.GetArmorClassAddition()
	}

	return armorClass
}
