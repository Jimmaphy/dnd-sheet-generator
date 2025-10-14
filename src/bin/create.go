package bin

import (
	"errors"
	"flag"
	"fmt"

	"github.com/jimmaphy/dnd-sheet-generator/domain"
	"github.com/jimmaphy/dnd-sheet-generator/repository"
)

type CreateCommand struct {
	name string
	race string
}

// NewCreateCommand creates a new instance of CreateCommand.
func NewCreateCommand() Command {
	return &CreateCommand{}
}

// ParseArguments will parse the command-line arguments for the create command.
// The name and race flags are required.
// If the name argument is missing, an error is returned.
func (command *CreateCommand) ParseArguments(args []string) error {
	flagSet := flag.NewFlagSet("createFlags", flag.ExitOnError)
	flagSet.StringVar(&command.name, "name", "", "Name of the character")
	flagSet.StringVar(&command.race, "race", "", "Race of the character")

	err := flagSet.Parse(args)
	if err != nil {
		return err
	}

	if command.name == "" {
		return errors.New("name is required")
	}

	if command.race == "" {
		return errors.New("race is required")
	}

	return nil
}

// Execute creates a new character sheet using the provided name.
// Then it attaches the race and class to the character, based on the name.
// It creates a character and saves it to a JSON file using the repository package.
func (command *CreateCommand) Execute() error {
	character := domain.NewCharacter(command.name)

	race := domain.NewRace(command.race)
	character.SetRace(race)

	err := repository.NewCharacterJSONRepository().Add(character)
	if err != nil {
		return err
	}

	fmt.Println("saved character", character.Name)
	return nil
}
