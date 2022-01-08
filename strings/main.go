package main

import (
	"fmt"
	"github.com/pkg/errors"
)

func main() {
	e := bb()
	fmt.Println(e)
	e1 := bc()
	fmt.Println(e1)
	fmt.Printf("%v, %v\n", e, e1)
	fmt.Println("------")
	fmt.Printf("%+v, %+v", e, e1)
	e = errors.Wrap(e, "new e")
	fmt.Printf("%+v", e)

}

func bb() error {
	return fmt.Errorf("aaa %s", "bb")
}

func bc() error {
	return errors.New("aaaa")
}
