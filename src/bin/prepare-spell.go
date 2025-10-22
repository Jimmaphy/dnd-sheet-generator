package bin

import (
	"errors"
	"flag"
	"fmt"

	"github.com/jimmaphy/dnd-sheet-generator/repository"
)

type PrepareSpellCommand struct {
	name string
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

	err := flagSet.Parse(args)
	if err != nil {
		return err
	}

	if command.name == "" {
		return errors.New("name is required")
	}

	return nil
}

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

	fmt.Println("PrepareSpellCommand executed for character:", command.name)
	return nil
}
