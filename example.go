package main

import "fmt"

const (
	Comment = 0x01
)

type f func() error

func main() {
	a := make([]int, 3)
	b := []int{1, 2, 23, 3, 3}
	num := copy(a, b)
	fmt.Println(a, num)
}
