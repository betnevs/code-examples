package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	err := errors.New("aaaa")
	fmt.Fprint(os.Stderr, err)
}
