package main

import (
	"encoding/json"
	stderrors "errors"
	"fmt"

	"github.com/alextanhongpin/go-errors/domain/entity/user"
	"github.com/alextanhongpin/go-errors/domain/errors"
)

func main() {
	var err error
	// This will fail, because it is a Partial error and does not fulfill the
	// error interface.
	//err = user.ErrInvalidAge
	err = user.ErrInvalidAge.Build(errors.M{"MaxAge": 200})
	fmt.Println(err)
	fmt.Printf("%#v\n", err)
	fmt.Println(errors.Is(err, user.ErrInvalidAge))

	err = user.ErrNotFound
	fmt.Println(err)
	fmt.Println(errors.Is(err, user.ErrNotFound))

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
}
