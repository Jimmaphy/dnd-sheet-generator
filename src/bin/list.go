package bin

import (
	"fmt"

	"github.com/jimmaphy/dnd-sheet-generator/repository"
)

type ListCommand struct{}

// NewListCommand creates a new instance of ListCommand.
func NewListCommand() Command {
	return &ListCommand{}
}

// ParseArguments will parse the command-line arguments for the list command.
// The list command does not require any arguments, so this function does nothing.
func (command *ListCommand) ParseArguments(args []string) error {
	return nil
}

// Execute lists all character sheets saved in the JSON repository.
func (command *ListCommand) Execute() error {
	repo := repository.NewCharacterJSONRepository()
	characters, err := repo.List()
	if err != nil {
		return err
	}

	if len(characters) == 0 {
		fmt.Println("no characters found.")
		return nil
	}

	for _, character := range characters {
		fmt.Printf("- %s\n", character)
	}

	return nil
}
