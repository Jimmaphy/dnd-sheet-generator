package main

import (
	"fmt"
	"os"

	"github.com/jimmaphy/dnd-sheet-generator/bin"
)

func main() {
	command, err := bin.GetCommander("usage")
	if err != nil {
		fmt.Println(err)
		return
	}

	err = command.ParseArguments(os.Args[1:])
	if err != nil {
		fmt.Println(err)
		return
	}

	err = command.Execute()
	if err != nil {
		fmt.Println(err)
		return
	}
}
