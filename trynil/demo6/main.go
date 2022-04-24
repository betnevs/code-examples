package main

import "fmt"

type Summer interface {
	Sum() int
}

type P struct {
	name string
}

func (p *P) Sum() int {
	return 1
}

type ints []int

func (i ints) Sum() int {
	s := 0
	for _, v := range i {
		s += v
	}
	return s
}

func main() {
	var i ints
	fmt.Println(i, i == nil)
	fmt.Println(doSum(i))

	var s []int
	fmt.Println(s, s == nil)

	var p *P
	fmt.Println(doSum(p))

	fmt.Println(doSum(nil))

}

func doSum(s Summer) int {
	if s == nil {
		fmt.Println("eq nil")
		return 0
	} else {
		fmt.Println("not eq nil")
		return s.Sum()
	}
}
