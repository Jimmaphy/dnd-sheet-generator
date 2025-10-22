package main

import (
	"fmt"
	"os"

	"github.com/jimmaphy/dnd-sheet-generator/bin"
)

// Main is the entry point of the application.
// It parses the first command-line argument to determine which command to execute.
// If no command is provided, it will continue with a default command.
// It handles errors by printing them to the console and exiting with a non-zero status code.
func main() {
	var args []string
	requestedCommand := "default"

	if len(os.Args) > 1 {
		requestedCommand = os.Args[1]
		args = os.Args[2:]
	}

	command, err := bin.GetCommander(requestedCommand)
	if err != nil {
		fmt.Println(err)
	}

	err = bin.ExecuteCommand(command, args)
	if err != nil {
		fmt.Println(err)
	}
}
