package bin

import (
	"fmt"

	"github.com/jimmaphy/dnd-sheet-generator/services"
)

type usageCommand struct{}

// Create a new instance of usageCommand.
func NewUsageCommand() Command {
	return &usageCommand{}
}

// ParseArguments for usageCommand does not require any arguments and always returns nil.
func (command *usageCommand) ParseArguments(args []string) error {
	return nil
}

// Execute reads and prints the usage instructions from the usage.txt template file.
// If the file cannot be read, an error message is printed and the program exits with status 1.
func (command *usageCommand) Execute() error {
	templateService, err := services.NewTemplateService("usage.go")
	if err != nil {
		return err
	}

	content, err := templateService.GetTemplateContent()
	if err != nil {
		return err
	}

	fmt.Print(content)
	return nil
}
