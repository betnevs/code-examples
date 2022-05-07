package main

import "fmt"

type myInt int

const n = 1
const n1 = 1

const (
	a = iota
	_
	b
	c = iota
)

func main() {
	a1 := [3]int{1, 2, 3}
	a2 := a1
	a1[0] = 99
	fmt.Println(a1, a2)
}
