package user

import (
	_ "embed"

	"github.com/BurntSushi/toml"
	"github.com/alextanhongpin/go-errors/domain/errors"
	"golang.org/x/text/language"
)

//go:embed errors.toml
var errorCodes []byte

var eb = errors.NewBundle(
	language.English,
	[]language.Tag{language.English, language.MustParse("ms")},
	[]errors.Kind{
		"unknown",
		"internal",
		"bad_input",
		"not_found",
		"already_exists",
		"failed_preconditions",
		"unauthorized",
		"forbidden",
	},
)

// User errors.
const (
	MinAge = 13
	MaxAge = 150
)

var (
	_              = eb.MustLoad(errorCodes, toml.Unmarshal)
	ErrNotFound    = eb.Code("user.notFound")
	ErrInvalidName = errors.Partial[InvalidNameParams](eb.Code("user.invalidName"))
	ErrInvalidAge  = errors.Partial[InvalidAgeParams](eb.Code("user.invalidAge")).
			SetParams(InvalidAgeParams{MaxAge: MaxAge})
	ErrUnderAge = errors.Partial[UnderAgeParams](eb.Code("user.underAge")).
			SetParams(UnderAgeParams{MinAge: MinAge})
	ErrValidationErrors = errors.Partial[ValidationErrors](eb.Code("user.validationErrors"))
)

type InvalidAgeParams struct {
	MaxAge int64 `json:"maxAge"`
}

type UnderAgeParams struct {
	MinAge int64 `json:"minAge"`
}

type InvalidNameParams struct {
	Name string `json:"name"`
}

type ValidationErrors []ValidationFieldError

type ValidationFieldError struct {
	Field  string `json:"field"`
	Reason string `json:"reason"`
}
