package usecase

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"

	"github.com/alextanhongpin/go-errors/domain/errors"
	"github.com/go-playground/validator/v10"
)

type Book struct {
	Author string `json:"author" validate:"required,min=1"`
}

type User struct {
	Name  string `json:"name" validate:"required,min=1"`
	Age   int64  `json:"age" validate:"required,gt=0"`
	Books []Book `json:"books" validate:"dive,required"`
}

func main() {
	v := validator.New()
	// Obtain the field name from the json struct tag.
	v.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})
	req := &User{Name: "a", Books: []Book{{}}}
	err := NewValidationErrors(v, req, "Validation failed for User")
	verr, _ := err.(*ValidationErrors)
	b, err := json.Marshal(verr)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(b))
	fmt.Println("done")
}

type FieldError struct {
	Name   string `json:"name"`
	Reason string `json:"reason"`
}

type ValidationErrors struct {
	Code    string       `json:"code"`
	Kind    string       `json:"kind"`
	Message string       `json:"message"`
	Fields  []FieldError `json:"reasons"`
}

func (err *ValidationErrors) Error() string {
	reasons := make([]string, len(err.Fields))
	for i, field := range err.Fields {
		reasons[i] = field.Reason
	}
	return fmt.Sprintf("%s: %s", err.Message, strings.Join(reasons, "\n"))
}

type Validator interface {
	Struct(in interface{}) error
}

func NewValidationErrors(v Validator, input interface{}, msg string) error {
	if err := v.Struct(input); err != nil {
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			return err
		}
		fields := make([]FieldError, len(errs))
		for i, e := range errs {
			fields[i] = FieldError{
				Name:   e.Namespace(),
				Reason: e.Error(),
			}
		}
		return &ValidationErrors{
			Code:    "ValidationError",
			Kind:    errors.BadInput.String(),
			Message: msg,
			Fields:  fields,
		}
	}
	return nil
}
