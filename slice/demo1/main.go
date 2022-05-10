package main

import (
	"fmt"
)

func myAppend(sl []int, elems ...int) []int {
	fmt.Printf("%T\n", elems) // []int
	fmt.Println(elems == nil)
	if len(elems) == 0 {
		println("no elems to append")
		return sl
	}

	sl = append(sl, elems...)
	return sl
}

func main() {
	var a interface{}
	fmt.Printf("a T = %T\n", a)
	sl := []int{1, 2, 3}
	sl = myAppend(sl) // no elems to append
	fmt.Println(sl)   // [1 2 3]
	sl = myAppend(sl, 4, 5, 6)
	fmt.Println(sl) // [1 2 3 4 5 6]
}
