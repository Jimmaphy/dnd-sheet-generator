package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/jimmaphy/dnd-sheet-generator/character"
	"github.com/jimmaphy/dnd-sheet-generator/helpers"
	"github.com/jimmaphy/dnd-sheet-generator/io"
)

type CreateHandler struct {
	name         string
	race         string
	class        string
	level        int
	strength     int
	dexterity    int
	constitution int
	intelligence int
	wisdom       int
	charisma     int
}

func (handler *CreateHandler) toCharacter() character.Character {
	return character.Character{
		Name:         handler.name,
		Race:         handler.race,
		Class:        handler.class,
		Level:        handler.level,
		Strength:     handler.strength,
		Dexterity:    handler.dexterity,
		Constitution: handler.constitution,
		Intelligence: handler.intelligence,
		Wisdom:       handler.wisdom,
		Charisma:     handler.charisma,
	}
}

func (handler *CreateHandler) ParseArguments() *helpers.ExitError {
	createCmd := flag.NewFlagSet("create", flag.ContinueOnError)

	createCmd.StringVar(&handler.name, "name", "", "character name (required)")
	createCmd.StringVar(&handler.race, "race", "", "character race (required)")
	createCmd.StringVar(&handler.class, "class", "", "character class (required)")
	createCmd.IntVar(&handler.level, "level", 1, "character level")
	createCmd.IntVar(&handler.strength, "str", 0, "strength score (required)")
	createCmd.IntVar(&handler.dexterity, "dex", 0, "dexterity score (required)")
	createCmd.IntVar(&handler.constitution, "con", 0, "constitution score (required)")
	createCmd.IntVar(&handler.intelligence, "int", 0, "intelligence score (required)")
	createCmd.IntVar(&handler.wisdom, "wis", 0, "wisdom score (required)")
	createCmd.IntVar(&handler.charisma, "cha", 0, "charisma score (required)")

	err := createCmd.Parse(os.Args[2:])
	if err != nil {
		return helpers.NewExitError(err, 2)
	}

	return nil
}

func (handler *CreateHandler) Handle(args []string) *helpers.ExitError {
	repository := io.NewJsonCharacterRepository()
	character := handler.toCharacter()

	err := repository.Add(character)
	if err != nil {
		return helpers.NewExitError(err, 4)
	}

	fmt.Println("saved character", character.Name)
	return nil
}
