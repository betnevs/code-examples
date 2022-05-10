package main

import (
	"fmt"
)

type Adder interface {
	add(int, int) int
}

type Iadd func(int, int) int

func (i *Iadd) add(a int, b int) int {
	return (*i)(a, b)
}

func main() {
	f := func(a int, b int) int {
		return a + b
	}

	var x Adder
	x = (*Iadd)(&f)
	fmt.Println(x.add(1, 2))

}
