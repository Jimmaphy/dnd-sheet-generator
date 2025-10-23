package bin

import (
	"errors"
	"flag"
	"fmt"

	"github.com/jimmaphy/dnd-sheet-generator/repository"
)

type LearnSpellCommand struct {
	name string
	spell string
}

// Create a new instance of LearnSpellCommand.
func NewLearnSpellCommand() Command {
	return &LearnSpellCommand{}
}

// ParseArguments for LearnSpellCommand will look for the name argument.
// If the name argument is missing, an error is returned.
func (command *LearnSpellCommand) ParseArguments(args []string) error {
	flagSet := flag.NewFlagSet("viewFlags", flag.ContinueOnError)
	flagSet.StringVar(&command.name, "name", "", "Name of the character to view")
	flagSet.StringVar(&command.spell, "spell", "", "Name of the spell to learn")

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

func (command *LearnSpellCommand) Execute() error {
	characterRepository := repository.NewCharacterJSONRepository()
	character, err := characterRepository.Get(command.name)
	if err != nil {
		return errors.New("character \"" + command.name + "\" not found")
	}

	if character.Class.CasterType == "none" {
		return errors.New("this class can't cast spells")
	}

	if character.Class.CasterType == "prepared" {
		return errors.New("this class prepares spells and can't learn them")
	}

	err = character.AddSpell(command.spell)
	if err != nil {
		return err
	}

	fmt.Println("Learned spell", command.spell)
	return nil
}
