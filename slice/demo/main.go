package main

import "fmt"

func main() {
	a := []int{1, 2, 3}
	b := a
	fmt.Printf("%p, %p\n", &a, &b)

	d := map[string]string{"aa": "aaa", "bb": "sxxx"}
	x := d
	fmt.Printf("%p, %p\n", &d, &x)
}
