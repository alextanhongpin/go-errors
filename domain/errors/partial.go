package errors

// PartialError creates partial errors that has to be built with params.
// Building the error will replace the message with the params.
type PartialError interface {
	Build(M) *Error
}

// Partial transform an error to a partial error.
func Partial(err *Error) PartialError {
	return err
}
