package bin

import (
	"errors"
	"flag"
	"os"

	"github.com/jimmaphy/dnd-sheet-generator/infrastructure"
	"github.com/jimmaphy/dnd-sheet-generator/repository"
)

type ExportCommand struct {
	name string
}

// Create a new ExportCommand instance
func NewExportCommand() Command {
	return &ExportCommand{}
}

// ParseArguments for ExportCommand parses the command-line arguments specific to the export command.
// The name argument is required and specifies the name of the exported file without extension.
func (command *ExportCommand) ParseArguments(args []string) error {
	exportFlags := flag.NewFlagSet("exportGroup", flag.ContinueOnError)
	exportFlags.StringVar(&command.name, "name", "character_sheet", "The name of the exported file without extension")

	err := exportFlags.Parse(args)
	if err != nil {
		return errors.New("failed to parse export command arguments: " + err.Error())
	}

	if command.name == "" {
		return errors.New("name is required")
	}

	return nil
}

// Execute runs the export command logic.
// The result shall be exported to an HTML file with the specified name.
func (command *ExportCommand) Execute() error {
	characterRepository := repository.NewCharacterJSONRepository()
	character, err := characterRepository.Get(command.name)
	if err != nil {
		return errors.New("character \"" + command.name + "\" not found")
	}

	templateService, err := infrastructure.NewTemplateService("charactersheet.html")
	if err != nil {
		return err
	}

	content, err := templateService.GetParsable()
	if err != nil {
		return err
	}

	character.CalculateTotalSkills()
	file, err := os.Create("./export/" + character.Name + ".html")
	if err != nil {
		return errors.New("failed to create export file: " + err.Error())
	}

	defer file.Close()
	content.Execute(file, character)

	return nil
}
