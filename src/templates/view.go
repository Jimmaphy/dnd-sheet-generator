Name: {{.Name}}
Class: {{.Class.Name}}
Race: {{.Race.Name}}
Background: {{.Background.Name}}
Level: {{.Level}}
Ability scores:
  STR: {{.TotalSkills.Strength}} ({{.TotalSkills.GetStrengthModifierString}})
  DEX: {{.TotalSkills.Dexterity}} ({{.TotalSkills.GetDexterityModifierString}})
  CON: {{.TotalSkills.Constitution}} ({{.TotalSkills.GetConstitutionModifierString}})
  INT: {{.TotalSkills.Intelligence}} ({{.TotalSkills.GetIntelligenceModifierString}})
  WIS: {{.TotalSkills.Wisdom}} ({{.TotalSkills.GetWisdomModifierString}})
  CHA: {{.TotalSkills.Charisma}} ({{.TotalSkills.GetCharismaModifierString}})
Proficiency bonus: {{.GetProficiencyBonus}}
Skill proficiencies: {{.GetSkillProficiencyString}}