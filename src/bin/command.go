package bin

type Command interface {
	// Parse command-line arguments
	// Return an error if when required arguments are missing or invalid
	ParseArguments(args []string) error

	// Execute the command with the parsed arguments
	// Return an error if execution fails
	Execute() error
}
