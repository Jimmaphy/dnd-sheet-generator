package bin

import (
	"errors"
	"flag"

	"github.com/jimmaphy/dnd-sheet-generator/repository"
)

type DeleteCommand struct {
	name string
}

func NewDeleteCommand() Command {
	return &DeleteCommand{}
}

// ParseArguments will parse the command-line arguments for the delete command.
// The name flag is required.
// If the name argument is missing, an error is returned.
func (command *DeleteCommand) ParseArguments(args []string) error {
	flagSet := flag.NewFlagSet("deleteFlags", flag.ExitOnError)
	flagSet.StringVar(&command.name, "name", "", "Name of the character to delete")

	err := flagSet.Parse(args)
	if err != nil {
		return err
	}

	if command.name == "" {
		return errors.New("name is required")
	}

	return nil
}

// Execute deletes a character sheet using the provided name.
// It removes the character from the JSON file using the repository package.
func (command *DeleteCommand) Execute() error {
	repository := repository.NewCharacterJSONRepository()
	err := repository.Delete(command.name)
	if err != nil {
		return err
	}

	println("deleted", command.name)
	return nil
}
