package user

import (
	_ "embed"

	"github.com/BurntSushi/toml"
	"github.com/alextanhongpin/go-errors/domain/errors"
	"golang.org/x/text/language"
)

//go:embed errors.toml
var errorCodes []byte

// User errors.
const (
	MinAge = 13
	MaxAge = 150
)

var (
	_             = errors.MustRegister(language.English, errorCodes, toml.Unmarshal)
	ErrNotFound   = errors.New("user.notFound")
	ErrInvalidAge = errors.NewParams[InvalidAgeParams]("user.invalidAge").SetParams(InvalidAgeParams{MaxAge: MaxAge})
	ErrUnderAge   = errors.NewParams[UnderAgeParams]("user.underAge").SetParams(UnderAgeParams{MinAge: MinAge})
)

type InvalidAgeParams struct {
	MaxAge int64 `json:"maxAge"`
}

type UnderAgeParams struct {
	MinAge int64 `json:"minAge"`
}
