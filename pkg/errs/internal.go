package errs

// InternalError is a structure used by the application to pass errors internally.
type InternalError struct {
	Message string
	Err     error
}

// Error implements the Error interface.
func (i InternalError) Error() string {
	return i.Message
}

// Unwrap enables error chaining
func (i *InternalError) Unwrap() error {
	return i.Err
}

var (

	// NoDBResults is the error for a database query with no results.
	NoDBResults = &InternalError{Message: "The query returned no results."}
)
