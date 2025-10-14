package bin

import (
	"errors"
)

var commands = map[string]func() Command{
	"default": NewUsageCommand,
	"usage":   NewUsageCommand,
	"create":  NewCreateCommand,
}

// GetCommander returns a Command instance based on the provided name.
// If the command name does not exist, the default command is returned,
// along with an error indicating the command was not found.
func GetCommander(name string) (Command, error) {
	if initializeCommand, exists := commands[name]; exists {
		return initializeCommand(), nil
	}

	return commands["default"](), errors.New("command '" + name + "' not found")
}

// ExecuteCommand executes the command with the given arguments.
func ExecuteCommand(command Command, args []string) error {
	if err := command.ParseArguments(args); err != nil {
		return err
	}

	if err := command.Execute(); err != nil {
		return err
	}

	return nil
}
