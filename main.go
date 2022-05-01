package main

import (
	"encoding/json"
	stderrors "errors"
	"fmt"

	"github.com/alextanhongpin/go-errors/domain/entity/user"
	"github.com/alextanhongpin/go-errors/domain/errors"
	"golang.org/x/text/language"
)

var (
	_ = errors.AddLanguage(language.MustParse("ms"))
)

func main() {
	var err error

	// This will fail, because it is a Partial error and does not fulfill the
	// error interface.
	//err = user.ErrInvalidAge
	err = user.ErrInvalidAge
	fmt.Println(err)
	fmt.Printf("%#v\n", err)
	fmt.Println(stderrors.Is(err, user.ErrInvalidAge))

	err = user.ErrNotFound
	fmt.Println(err)
	fmt.Println(stderrors.Is(err, user.ErrNotFound))

	err = fmt.Errorf("failed to find user: %w", err)
	fmt.Println(err)

	var userNotFoundErr *errors.Error
	if stderrors.As(err, &userNotFoundErr) {
		fmt.Println(userNotFoundErr)
	}
	b, err := json.Marshal(userNotFoundErr)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(b))

	b, err = json.Marshal(user.ErrInvalidAge.SetLanguage(language.MustParse("ms")))
	if err != nil {
		panic(err)
	}
	fmt.Println(string(b))
}

//Unknown            Kind = iota // unknown
//Internal                       // internal
//BadInput                       // bad_input
//NotFound                       // not_found
//AlreadyExists                  // already_exists
//FailedPrecondition             // failed_precondition
//Unauthorized                   // unauthorized
//Forbidden                      // forbidden
