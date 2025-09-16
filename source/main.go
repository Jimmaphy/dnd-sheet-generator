package main

import (
	"errors"
	"os"

	"github.com/jimmaphy/dnd-sheet-generator/helpers"
)

func main() {
	var err *helpers.ExitError
	err = nil

	if len(os.Args) > 1 {
		switch os.Args[1] {

		case "create":
			err = helpers.ApplyHandler(&CreateHandler{}, os.Args[2:])

		case "view":

		case "list":

		case "delete":

		case "equip":

		case "learn-spell":

		case "prepare-spell":

		default:
			err = helpers.NewExitError(errors.New("unknown command: "+os.Args[1]), 127)
		}
	} else {
		err = helpers.NewExitError(errors.New("no command was provided"), 1)
	}

	if err != nil {
		println(err.Error())
		printUsage()
		os.Exit(err.ExitCode())
	}
}
