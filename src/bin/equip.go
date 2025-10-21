package bin

import (
	"errors"
	"flag"
	"fmt"

	"github.com/jimmaphy/dnd-sheet-generator/domain"
	"github.com/jimmaphy/dnd-sheet-generator/repository"
)

type EquipCommand struct {
	name   string
	weapon string
	slot   string
	armor  string
	shield string
	state  string
}

// Create a new instance of EquipCommand.
func NewEquipCommand() Command {
	return &EquipCommand{}
}

// ParseArguments will parse the command-line arguments for the equip command.
// There needs to be a valid combination of arguments provided.
// These are either item and slot; armor; or shield.
// They will be checked in order, and the first valid combination will be used.
// Don't forget the name of the character to equip the item to.
func (command *EquipCommand) ParseArguments(args []string) error {
	flagSet := flag.NewFlagSet("equipFlags", flag.ContinueOnError)
	flagSet.StringVar(&command.name, "name", "", "Name of the character")
	flagSet.StringVar(&command.weapon, "weapon", "", "Weapon to equip")
	flagSet.StringVar(&command.slot, "slot", "", "Slot to equip the item")
	flagSet.StringVar(&command.armor, "armor", "", "Armor to equip")
	flagSet.StringVar(&command.shield, "shield", "", "Shield to equip")

	err := flagSet.Parse(args)
	if err != nil {
		return err
	}

	if command.name == "" {
		return errors.New("name is required")
	}

	if command.weapon != "" && command.slot != "" {
		command.state = "weapon"
	} else if command.armor != "" {
		command.state = "armor"
	} else if command.shield != "" {
		command.state = "shield"
	} else {
		return errors.New("either item and slot, armor, or shield is required")
	}

	return nil
}

// Execute equips the specified item to the character based on the parsed arguments.
// First, the character is retrieved from the repository.
// Then, depending on the state, the appropriate equip function is called.
func (command *EquipCommand) Execute() error {
	characterRepository := repository.NewCharacterJSONRepository()
	character, err := characterRepository.Get(command.name)
	if err != nil {
		return errors.New("character \"" + command.name + "\" not found")
	}

	switch command.state {
	case "weapon":
		return command.equipWeapon(character)
	case "armor":
		return command.equipArmor(character)
	case "shield":
		return command.equipShield(character)
	default:
		return errors.New("invalid equip state")
	}
}

func (command *EquipCommand) equipWeapon(character *domain.Character) error {
	fmt.Println("Executing state", command.state, "for", character.Name)
	return nil
}

func (command *EquipCommand) equipArmor(character *domain.Character) error {
	fmt.Println("Executing state", command.state, "for", character.Name)
	return nil
}

func (command *EquipCommand) equipShield(character *domain.Character) error {
	fmt.Println("Executing state", command.state, "for", character.Name)
	return nil
}
