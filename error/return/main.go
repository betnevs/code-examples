package main

import (
	"fmt"
)

type doError struct {
}

func (d *doError) Error() string {
	return "aaaa"
}

func do() error {
	return nil
}

func wrapDo() error {
	return do()
}

func main() {
	err := wrapDo()
	switch err.(type) {
	case nil:
		fmt.Println("nil error")
	case error:
		fmt.Println("real error")
	}
	fmt.Println(err == nil)
}
