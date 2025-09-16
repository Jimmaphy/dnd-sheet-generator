package helpers

type CommandHandler interface {
	ParseArguments() *ExitError
	Handle(args []string) *ExitError
}

func ApplyHandler(handler CommandHandler, args []string) *ExitError {
	err := handler.ParseArguments()
	if err == nil {
		err = handler.Handle(args)
	}

	return err
}
