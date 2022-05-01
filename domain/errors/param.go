package errors

import (
	"golang.org/x/text/language"
)

type partialError[T any] interface {
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

func Partial[T any](err *Error) partialError[T] {
	return NewErrorParams[T](err)
}

func (e *ErrorParams[T]) SetParams(t T) *Error {
	err := e.err.Clone()
	newTranslations := make(map[language.Tag]string)

	for lang, msg := range err.translations {
		tmsg, err := makeTemplate(msg, t)
		if err != nil {
			newTranslations[lang] = msg
		} else {
			newTranslations[lang] = tmsg
		}
	}

	err.translations = newTranslations
	err.Message = newTranslations[err.lang]
	return err
}

func (e *ErrorParams[T]) Self() *Error {
	return e.err.Clone()
}
