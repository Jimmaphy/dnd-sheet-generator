package bin

import (
	"errors"
	"flag"
	"fmt"

	"github.com/jimmaphy/dnd-sheet-generator/repository"
)

type PrepareSpellCommand struct {
	name string
	spell string
}

// Create a new instance of PrepareSpellCommand.
func NewPrepareSpellCommand() Command {
	return &PrepareSpellCommand{}
}

// ParseArguments for PrepareSpellCommand will look for the name argument.
// If the name argument is missing, an error is returned.
func (command *PrepareSpellCommand) ParseArguments(args []string) error {
	flagSet := flag.NewFlagSet("viewFlags", flag.ContinueOnError)
	flagSet.StringVar(&command.name, "name", "", "Name of the character to view")
	flagSet.StringVar(&command.spell, "spell", "", "Name of the spell to prepare")

	err := flagSet.Parse(args)
	if err != nil {
		return err
	}

	if command.name == "" {
		return errors.New("name is required")
	}

	if command.spell == "" {
		return errors.New("spell is required")
	}

	return nil
}

// Execute will prepare the specified spell for the character.
// It will check if the character exists and if their class allows spell preparation.
// Then it will add the spell to the character's spell list.
// An error will be thrown if the character cannot learn the spell.
func (command *PrepareSpellCommand) Execute() error {
	characterRepository := repository.NewCharacterJSONRepository()
	character, err := characterRepository.Get(command.name)
	if err != nil {
		return errors.New("character \"" + command.name + "\" not found")
	}

	if character.Class.CasterType == "none" {
		return errors.New("this class can't cast spells")
	}

	if character.Class.CasterType == "learned" {
		return errors.New("this class learns spells and can't prepare them")
	}

	err = character.AddSpell(command.spell)
	if err != nil {
		return err
	}

	fmt.Println("Prepared spell", command.spell)
	return nil
}
