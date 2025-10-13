package main

import (
	"fmt"
	"os"

	"github.com/jimmaphy/dnd-sheet-generator/bin"
)

func main() {
	command, err := bin.GetCommander("usage")
	if err != nil {
		bin.ExecuteCommand(command, os.Args[1:])
		fmt.Println(err)
		os.Exit(1)
	}

	err = bin.ExecuteCommand(command, os.Args[1:])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
