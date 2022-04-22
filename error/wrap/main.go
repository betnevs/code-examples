package main

import (
	"fmt"

	"github.com/pkg/errors"
)

var sentinelErr = fmt.Errorf("i am error")

func wrapNewPointerError() error {
	return errors.Wrap(fmt.Errorf("i am error"), "wrap err")
}

func wrapConstantPointerError() error {
	return fmt.Errorf("wrap err: %w", sentinelErr)
}

func main() {
	fmt.Println(errors.Is(wrapNewPointerError(), fmt.Errorf("i am error")))
	fmt.Println(errors.Is(wrapConstantPointerError(), fmt.Errorf("i am error")))
	fmt.Println(errors.Is(wrapConstantPointerError(), sentinelErr))
	b := wrapConstantPointerError()
	c := errors.Wrap(b, "aaa")
	fmt.Println(b, errors.Is(b, sentinelErr))
	fmt.Printf("%+v", c)
}
