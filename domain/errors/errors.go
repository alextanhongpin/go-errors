package errors

import (
	"bytes"
	_ "embed"
	"fmt"
	"html/template"
	"log"

	"github.com/BurntSushi/toml"
)

var errors = map[Code]*Error{}
var uniqueErrors = map[Code]struct{}{}

// Code represents the breakdown on an error.
type Code string

// M is an alias for map[string]interface{}.
type M map[string]interface{}

// Error represents an error.
type Error struct {
	Kind    string `json:"kind"`
	Code    Code   `json:"code"`
	Message string `json:"message"`
	Params  M      `json:"params,omitempty"`
}

// newError creates a new error. This should only be done by the error package.
func newError(kind Kind, code Code, msg string) *Error {
	return &Error{
		Kind:    kind.String(),
		Code:    code,
		Message: msg,
		Params:  make(map[string]interface{}),
	}
}

// Error fulfills the error interface.
func (e Error) Error() string {
	msg, err := build(e.Message, e.Params)
	if err != nil {
		log.Printf("failed to build error message: %s\n", err)
	}
	return msg
}

// Build builds the error with the given params.
func (e Error) Build(params M) *Error {
	// Instead of overriding, we merge the params.
	for k, v := range params {
		e.Params[k] = v
	}
	return &e
}

// Is satisfies the error interface.
func (e Error) Is(target error) bool {
	t, ok := target.(*Error)
	if !ok {
		return false
	}
	return e.Kind == t.Kind &&
		e.Code == t.Code
}

// Is a shortcut for Is.
func Is(err error, tgt interface{}) bool {
	switch t := tgt.(type) {
	case PartialError:
		return t.Build(nil).Is(err)
	case *Error:
		return t.Is(err)
	default:
		return false
	}
}

func build(msg string, data map[string]interface{}) (string, error) {
	t := template.Must(template.New("").Parse(msg))
	var buf bytes.Buffer
	if err := t.Execute(&buf, data); err != nil {
		return msg, err
	}
	return buf.String(), nil
}

func checkUnique(code Code) {
	_, exists := uniqueErrors[code]
	if exists {
		panic(fmt.Sprintf("error code already exists: %s", code))
	}
	uniqueErrors[code] = struct{}{}
}

func C(code Code) *Error {
	checkUnique(code)

	err, ok := errors[code]
	if !ok {
		panic(fmt.Sprintf("error code not found: %s", code))
	}
	return err
}

func P(code Code) PartialError {
	return Partial(C(code))
}

func Register(raw []byte) bool {
	var data map[string]map[string]string
	if err := toml.Unmarshal(raw, &data); err != nil {
		panic(err)
	}
	for kind, codes := range data {
		for code, msg := range codes {
			c := Code(code)
			errors[c] = newError(KindFromStr(kind), c, msg)
		}
	}
	return true
}
