package adapters

import (
	"strings"

	"github.com/jimmaphy/dnd-sheet-generator/domain"
)

type ApiClass struct {
	Name               string                  `json:"name"`
	HitDie             int                     `json:"hit_die"`
	ProficiencyChoices []ApiProficiencyChoices `json:"proficiency_choices"`
	SpellCasting       ApiSpellCasting         `json:"spellcasting"`
	SpellUrl           string                  `json:"spells"`
}

// ToDomainModel converts the ApiClass to a domain model Class
func (response *ApiClass) ToDomainModel(spells []*domain.Spell, spellLevels []*domain.SpellLevel) *domain.Class {
	return &domain.Class{
		Name:          strings.ToLower(response.Name),
		HitDie:        response.HitDie,
		SkillCount:    response.ProficiencyChoices[0].Choose,
		Skills:        response.getSkillOptions(),
		CasterType:    response.getCasterType(),
		CastAbility:   response.getCasterAbility(),
		SpellSaveBase: 8,
		Spells:        spells,
		SpellLevels:   spellLevels,
	}
}

// GetSkillOptions returns the skill options for the class
// The response is formatted as a slice of strings.
func (response *ApiClass) getSkillOptions() []string {
	skillOptions := []string{}

	for _, option := range response.ProficiencyChoices[0].From.Options {
		name, _ := strings.CutPrefix(option.Item.Name, "Skill: ")
		name = strings.ToLower(name)
		skillOptions = append(skillOptions, name)
	}

	return skillOptions
}

// GetCasterType determines the caster type of the class
// It returns "none", "prepared", "learned", or "pact"
func (response *ApiClass) getCasterType() string {
	for _, info := range response.SpellCasting.Info {
		if strings.Contains(strings.ToLower(info.Name), "preparing") {
			return "prepared"
		}
	}

	if response.Name == "Warlock" {
		return "pact"
	}

	if len(response.SpellCasting.Info) == 0 {
		return "none"
	}

	return "learned"
}

// GetCasterAbility returns the caster ability of the class
func (response *ApiClass) getCasterAbility() string {
	switch strings.ToLower(response.SpellCasting.Ability.Name) {
	case "str":
		return "strength"
	case "dex":
		return "dexterity"
	case "con":
		return "constitution"
	case "wis":
		return "wisdom"
	case "int":
		return "intelligence"
	case "cha":
		return "charisma"
	default:
		return ""
	}
}
