package bin

import (
	"fmt"

	"github.com/jimmaphy/dnd-sheet-generator/services"
)

type UsageCommand struct{}

// Create a new instance of UsageCommand.
func NewUsageCommand() Command {
	return &UsageCommand{}
}

// ParseArguments for UsageCommand does not require any arguments and always returns nil.
func (command *UsageCommand) ParseArguments(args []string) error {
	return nil
}

// Execute reads and prints the usage instructions from the usage.txt template file.
// If the file cannot be read, an error message is printed and the program exits with status 1.
func (command *UsageCommand) Execute() error {
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
