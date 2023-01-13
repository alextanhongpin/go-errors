package usecase

import (
	"fmt"

	_ "embed"

	"github.com/BurntSushi/toml"
	"github.com/alextanhongpin/errors"
)

// go:embed *.toml
//var errorFiles embed.FS

//go:embed errors.toml
var errorBytes []byte
var (
	//_               = errors.MustLoadFS(errorFiles, toml.Unmarshal)
	_               = errors.MustLoad(errorBytes, toml.Unmarshal)
	ErrUserNotFound = errors.Get("user.not_found")
)

func UserNotFoundError(name string) error {
	type userNotFoundErrorParams struct {
		Name string
	}

	return errors.ToPartial[userNotFoundErrorParams](ErrUserNotFound).
		WithParams(userNotFoundErrorParams{
			Name: name,
		})
}

func init() {
	fmt.Println("usecase init")
}

func GetUser() error {
	return fmt.Errorf("not found: %w", UserNotFoundError("john"))
}
