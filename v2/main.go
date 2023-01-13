package main

import (
	"encoding/json"
	"fmt"

	stderrors "errors"

	"github.com/alextanhongpin/errors"
	"github.com/alextanhongpin/go-errors/v2/usecase"
)

func init() {
	fmt.Println("main init")
	_ = errors.MustAddKinds("not_found", "unknown")
}

func main() {
	err := usecase.GetUser()
	fmt.Println(stderrors.Is(err, usecase.ErrUserNotFound))
	fmt.Println(err)
	var appError *errors.Error
	if stderrors.As(err, &appError) {
		fmt.Println(appError.Tags)
		fmt.Println(appError.Params)
		b, err := json.Marshal(appError)
		if err != nil {
			panic(err)
		}
		fmt.Println(string(b))
	}
}
