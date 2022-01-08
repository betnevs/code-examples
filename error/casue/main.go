package main

import (
	"errors"
	"fmt"
	xerrors "github.com/pkg/errors"
)

func main() {
	err1 := errors.New("one")
	err2 := fmt.Errorf("aa %w", err1)
	fmt.Printf("%+v\n", err2)

	err3 := errors.New("three")
	err4 := xerrors.Wrap(err3, "four")

	fmt.Printf("%+v\n", err4)
	fmt.Println(xerrors.Cause(err4))
	fmt.Println(xerrors.Cause(err2))
}
