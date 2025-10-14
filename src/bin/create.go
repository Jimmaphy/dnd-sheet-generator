package bin

import (
	"errors"
	"flag"
	"fmt"

	"github.com/jimmaphy/dnd-sheet-generator/domain"
)

type CreateCommand struct {
	name string
}

// NewCreateCommand creates a new instance of CreateCommand.
func NewCreateCommand() Command {
	return &CreateCommand{}
}

// ParseArguments will parse the command-line arguments for the create command.
// The only required argument is --name.
// If the name argument is missing, an error is returned.
func (command *CreateCommand) ParseArguments(args []string) error {
	flagSet := flag.NewFlagSet("createFlags", flag.ExitOnError)
	flagSet.StringVar(&command.name, "name", "", "Name of the character")

	err := flagSet.Parse(args)
	if err != nil {
		return err
	}

	if command.name == "" {
		return errors.New("name is required")
	}

	return nil
}

// Execute creates a new character sheet using the services package.
// It creates a character and saves it
func (command *CreateCommand) Execute() error {
	character := domain.NewCharacter(command.name)
	fmt.Println("saved character", character.Name)
	return nil
}
