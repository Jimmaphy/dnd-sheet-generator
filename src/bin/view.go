package bin

import (
	"errors"
	"flag"
	"os"

	"github.com/jimmaphy/dnd-sheet-generator/repository"
	"github.com/jimmaphy/dnd-sheet-generator/services"
)

type viewCommand struct {
	name string
}

// Create a new instance of viewCommand.
func NewViewCommand() Command {
	return &viewCommand{}
}

// ParseArguments for viewCommand will look for the name argument.
// If the name argument is missing, an error is returned.
func (command *viewCommand) ParseArguments(args []string) error {
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

// Execute reads and prints the usage instructions from the usage.txt template file.
// If the file cannot be read, an error message is printed and the program exits with status 1.
func (command *viewCommand) Execute() error {
	characterRepository := repository.NewCharacterJSONRepository()
	character, err := characterRepository.Get(command.name)
	if err != nil {
		return errors.New("character \"" + command.name + "\" not found")
	}

	templateService, err := services.NewTemplateService("view.go")
	if err != nil {
		return err
	}

	content, err := templateService.GetParsable()
	if err != nil {
		return err
	}

	character.CalculateTotalSkills()
	content.Execute(os.Stdout, character)
	return nil
}
