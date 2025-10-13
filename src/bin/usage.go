package bin

import (
	"fmt"
	"os"

	"github.com/jimmaphy/dnd-sheet-generator/services"
)

type usageCommand struct{}

// Create a new instance of usageCommand.
func NewUsageCommand() Command {
	return &usageCommand{}
}

// ParseArguments for usageCommand does not require any arguments and always returns nil.
func (u *usageCommand) ParseArguments(args []string) error {
	return nil
}

// Execute reads and prints the usage instructions from the usage.txt template file.
// If the file cannot be read, an error message is printed and the program exits with status 1.
func (u *usageCommand) Execute() error {
	templateService, err := services.NewTemplateService("usage.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	content, err := templateService.GetTemplateContent()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Print(content)
	return nil
}
