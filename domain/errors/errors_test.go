package errors_test

import (
	stderrors "errors"
	"testing"

	_ "embed"

	"github.com/alextanhongpin/go-errors/domain/errors"
)

//go:embed errors_test.toml
var errorCodes []byte

// Sample error
var (
	_              = errors.Register(errorCodes)
	UserNotFound   = errors.C("user.notFound")
	UserIDNotFound = errors.P("user.idNotFound")
)

func TestError(t *testing.T) {
	t.Parallel()

	var err error
	err = UserNotFound
	if err.Error() != "User not found" {
		t.Error("error.Error() should return 'User not found'")
	}

	if !stderrors.Is(err, UserNotFound) {
		t.Error("errors.Is should be UserNotFound")
	}

	err = UserIDNotFound.Build(map[string]interface{}{
		"ID": 10,
	})
	if want := "User ID 10 not found"; want != err.Error() {
		t.Errorf("expected err.Error() to be %q, got %q", want, err.Error())
	}

	err = UserIDNotFound.Build(errors.M{
		"ID": "random-uuid",
	})
	if want := "User ID random-uuid not found"; want != err.Error() {
		t.Errorf("expected err.Error() to be %q, got %q", want, err.Error())
	}

	if !stderrors.Is(err, UserIDNotFound.Build(nil)) {
		t.Error("stderrors.Is should be UserIDNotFoundError")
	}

	if !errors.Is(err, UserIDNotFound) {
		t.Error("errors.Is should be UserIDNotFoundError")
	}

	var userIDNotFoundError *errors.Error
	if !stderrors.As(err, &userIDNotFoundError) {
		t.Error("errors.As should be UserIDNotFoundError")
	}

	if errors.NotFound != userIDNotFoundError.Kind {
		t.Error("userIDNotFoundError.Kind should be errors.NotFound")
	}

	if errors.Code("user.idNotFound") != userIDNotFoundError.Code {
		t.Error("userIDNotFoundError.Code should be errors.UserIDNotFound")
	}
}
