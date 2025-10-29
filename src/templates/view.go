Name: {{.Name}}
Class: {{.Class.Name}}
Race: {{.Race.Name}}
Background: {{.Background.Name}}
Level: {{.Level}}
Ability scores:
  STR: {{.TotalSkills.Strength.Value}} ({{.TotalSkills.Strength.GetModifierString}})
  DEX: {{.TotalSkills.Dexterity.Value}} ({{.TotalSkills.Dexterity.GetModifierString}})
  CON: {{.TotalSkills.Constitution.Value}} ({{.TotalSkills.Constitution.GetModifierString}})
  INT: {{.TotalSkills.Intelligence.Value}} ({{.TotalSkills.Intelligence.GetModifierString}})
  WIS: {{.TotalSkills.Wisdom.Value}} ({{.TotalSkills.Wisdom.GetModifierString}})
  CHA: {{.TotalSkills.Charisma.Value}} ({{.TotalSkills.Charisma.GetModifierString}})
Proficiency bonus: {{.GetProficiencyBonusString}}
Skill proficiencies: {{.GetSkillProficiencyString}}{{ if .Class.CanCastSpells .Level }}
Spell slots:
{{ .Class.SpellSlotsString .Level }}
Spellcasting ability: {{.Class.CastAbility}}
Spell save DC: {{.GetSpellSaveDC}}
Spell attack bonus: {{.GetSpellAttackBonusString}}{{ end }}{{ if .MainHand }}
Main hand: {{.MainHand.Name}}{{ end }}{{ if .OffHand }}
Off hand: {{.OffHand.Name}}{{ end }}{{ if .Armor }}
Armor: {{.Armor.Name}}{{ end }}{{ if .Shield }}
Shield: {{.Shield.Name}}{{ end }}
Armor class: {{.GetArmorClass}}
Initiative bonus: {{.GetInitiativeBonus}}
Passive perception: {{.GetPassivePerception}}
