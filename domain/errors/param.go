package errors

import (
	"golang.org/x/text/language"
)

type PartialError[T any] interface {
	SetParams(t T) *Error
	Self() *Error
}

type ErrorParams[T any] struct {
	err *Error
}

func NewErrorParams[T any](err *Error) *ErrorParams[T] {
	return &ErrorParams[T]{
		err: err.Clone(),
	}
}

func (e *ErrorParams[T]) SetParams(params T) *Error {
	err := e.err.Clone()
	err.translations = make(map[language.Tag]string)

	for lang, msg := range e.err.translations {
		tmsg, terr := makeTemplate(msg, params)
		if terr != nil {
			err.translations[lang] = msg
		} else {
			err.translations[lang] = tmsg
		}
	}

	err.Message = err.translations[err.lang]
	err.Params = params

	return err
}

// Self returns the base error. Useful when checking
// errors.Is without setting the params.
func (e *ErrorParams[T]) Self() *Error {
	return e.err.Clone()
}

// Partial returns a partial error that needs params to be
// set.
func Partial[T any](err *Error) PartialError[T] {
	return NewErrorParams[T](err)
}
