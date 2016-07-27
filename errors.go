package msgraph

// GraphAPIError is an implementation of error.
type GraphAPIError struct {
	// Message is a string representation of this error.
	Message string
	// InnerError is the error that triggered this error, if any.
	InnerError error
}

// Error implements the Error interface.
func (e *GraphAPIError) Error() string {
	return e.Message
}
