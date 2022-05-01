package errors

import (
	"encoding/json"
	"errors"
	"fmt"

	"golang.org/x/text/language"
)

type ErrorParams[T any] struct {
	err    *Error
	Params T
}

func NewErrorParams[T any](err *Error) *ErrorParams[T] {
	return &ErrorParams[T]{
		err: err,
	}
}

func (e *ErrorParams[T]) SetLanguage(lang language.Tag) *ErrorParams[T] {
	msg, ok := e.err.translations[lang]
	if !ok {
		panic(fmt.Errorf("language %q not supported", lang))
	}
	e.err.lang = lang
	e.err.Message = msg

	return e
}

func (e *ErrorParams[T]) SetParams(t T) *ErrorParams[T] {
	e.Params = t

	return e
}

// Error fulfills the error interface.
func (e ErrorParams[T]) Error() string {
	msg, err := makeTemplate(e.err.Message, e.Params)
	if err != nil {
		return e.err.Message
	}
	return msg
}

// Is satisfies the error interface.
func (e ErrorParams[T]) Is(target error) bool {
	var err *ErrorParams[T]
	if errors.As(target, err) {
		return err.err.Is(e.err)
	}

	return false
}

func (e ErrorParams[T]) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Kind    string `json:"kind"`
		Code    string `json:"code"`
		Message string `json:"message"`
		Params  T      `json:"params"`
	}{
		Kind:    e.err.Kind,
		Code:    e.err.Code,
		Message: e.Error(),
		Params:  e.Params,
	})
}
