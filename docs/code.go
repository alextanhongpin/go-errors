package docs

import (
	_ "embed"

	"github.com/alextanhongpin/go-errors/domain/errors"
)

//go:embed errors.toml
var errorCodes []byte

// User errors.
var (
	// Register error codes before using them.
	_ = errors.Register(errorCodes)

	// Sentinel error from a given error code.
	ErrNotFound = errors.C("user.notFound")

	// A partial error from a given error code. Partial must be called with
	// Build, in order to ensure the params are passed in for constructing the
	// error message.
	ErrInvalidAge = errors.P("user.invalidAge")

	// Uncomment to hit runtime error on duplicate key.
	//DuplicateError = errors.C("user.notFound")

	// Uncomment to hit runtime error on missing key.
	//ErrInvalidEmail = errors.C("user.invalidEmail")
)
