package user

import (
	_ "embed"

	"github.com/alextanhongpin/go-errors/domain/errors"
	"golang.org/x/text/language"
	"gopkg.in/yaml.v2"
)

//go:embed errors.json
// go:embed errors.toml
// go:embed errors.yaml
var errorCodes []byte

var (
	EN = language.English
	MS = language.Malay
)

var eb = errors.NewBundle(&errors.Options{
	DefaultLanguage:  EN,
	AllowedLanguages: []language.Tag{MS},
	AllowedKinds: []errors.Kind{
		"unknown",
		"internal",
		"bad_input",
		"not_found",
		"already_exists",
		"failed_preconditions",
		"unauthorized",
		"forbidden",
	},
	UnmarshalFn: yaml.Unmarshal,
	//UnmarshalFn: toml.Unmarshal,
	//UnmarshalFn: json.Unmarshal,
})

// User errors.
const (
	MinAge = 13
	MaxAge = 150
)

var (
	_              = eb.MustLoad(errorCodes)
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
