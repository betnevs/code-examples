package main

import (
	"fmt"
)

func main() {
	var a []int
	fmt.Println(a == nil)
	a = []int{1, 2, 3}
	fmt.Println(a == nil)
}
