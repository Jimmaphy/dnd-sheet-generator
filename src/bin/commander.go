package bin

import (
	"errors"
)

var commands = map[string]func() Command{
	"usage": NewUsageCommand,
}

func GetCommander(name string) (Command, error) {
	if initializeCommand, exists := commands[name]; exists {
		return initializeCommand(), nil
	}

	return nil, errors.New("command '" + name + "' not found")
}
