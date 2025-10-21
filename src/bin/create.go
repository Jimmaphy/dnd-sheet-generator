package bin

import (
	"flag"
	"fmt"
	"os"

	"github.com/jimmaphy/dnd-sheet-generator/domain"
	"github.com/jimmaphy/dnd-sheet-generator/repository"
)

type CreateCommand struct {
	name       string
	race       string
	class      string
	background string
	level      int
	str        int
	dex        int
	con        int
	intl       int
	wis        int
	cha        int
}

// NewCreateCommand creates a new instance of CreateCommand.
func NewCreateCommand() Command {
	return &CreateCommand{}
}

// ParseArguments will parse the command-line arguments for the create command.
// The name and race flags are required.
// If the name argument is missing, an error is returned.
func (command *CreateCommand) ParseArguments(args []string) error {
	flagSet := flag.NewFlagSet("createFlags", flag.ContinueOnError)
	flagSet.StringVar(&command.name, "name", "", "Name of the character")
	flagSet.StringVar(&command.race, "race", "human", "Race of the character")
	flagSet.StringVar(&command.class, "class", "barbarian", "Class of the character")
	flagSet.StringVar(&command.background, "background", "acolyte", "Background of the character")
	flagSet.IntVar(&command.level, "level", 1, "Level of the character")
	flagSet.IntVar(&command.str, "str", 10, "Strength score")
	flagSet.IntVar(&command.dex, "dex", 10, "Dexterity score")
	flagSet.IntVar(&command.con, "con", 10, "Constitution score")
	flagSet.IntVar(&command.intl, "int", 10, "Intelligence score")
	flagSet.IntVar(&command.wis, "wis", 10, "Wisdom score")
	flagSet.IntVar(&command.cha, "cha", 10, "Charisma score")

	err := flagSet.Parse(args)
	if err != nil {
		return err
	}

	if command.name == "" {
		fmt.Println("name is required")
		os.Exit(2)
	}

	return nil
}

// Execute creates a new character sheet using the provided name.
// Then it attaches the race and class to the character, based on the name.
// It creates a character and saves it to a JSON file using the repository package.
func (command *CreateCommand) Execute() error {
	race, err := repository.NewRaceJSONRepository().Get(command.race)
	if err != nil {
		return err
	}

	class, err := repository.NewClassJSONRepository().Get(command.class)
	if err != nil {
		return err
	}

	background, err := repository.NewBackgroundJSONRepository().Get(command.background)
	if err != nil {
		return err
	}

	character := domain.NewCharacter(command.name, command.level)
	command.assignAllProperties(character, race, class, background)
	err = repository.NewCharacterJSONRepository().Add(character)
	if err != nil {
		return err
	}

	fmt.Println("saved character", character.Name)
	return nil
}

// Internal function for assigning all properties to a character
// There are no checks here, as this is only used internally within the package.
func (command *CreateCommand) assignAllProperties(character *domain.Character, race *domain.Race, class *domain.Class, background *domain.Background) {
	skills := domain.NewSkillSet(command.str, command.dex, command.con, command.intl, command.wis, command.cha)

	character.SetRace(race)
	character.SetClass(class)
	character.SetBackground(background)
	character.SetSkillSet(skills)

	character.CalculateTotalSkills()
}
