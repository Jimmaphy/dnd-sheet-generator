package helpers

type ExitError struct {
	error
	Code int
}

func (exitError *ExitError) Error() string {
	return exitError.error.Error()
}

func (exitError *ExitError) ExitCode() int {
	return exitError.Code
}

func NewExitError(err error, code int) *ExitError {
	return &ExitError{
		error: err,
		Code:  code,
	}
}
