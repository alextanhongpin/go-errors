package errors

import (
	"errors"
	"fmt"

	"golang.org/x/text/language"
)

// Error represents an error.
type Error struct {
	Kind         string `json:"kind"`
	Code         string `json:"code"`
	Message      string `json:"message"`
	lang         language.Tag
	translations map[language.Tag]string
}

func (e *Error) SetLanguage(lang language.Tag) *Error {
	msg, ok := e.translations[lang]
	if !ok {
		panic(fmt.Errorf("language %q not supported", lang))
	}
	e.lang = lang
	e.Message = msg

	return e
}

// Error fulfills the error interface.
func (e Error) Error() string {
	return e.Message
}

// Is satisfies the error interface.
func (e Error) Is(target error) bool {
	var err *Error
	if errors.As(target, err) {
		return err.Kind == e.Kind && err.Code == e.Code
	}

	return false
}
