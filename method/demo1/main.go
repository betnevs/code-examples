package main

import "fmt"

type T struct {
	a int
}

func (t T) Get() int {
	return t.a
}

func (t *T) Set(a int) int {
	t.a = a
	return t.a
}

func main() {

	var t T
	t.Get()
	fmt.Printf("%#v\n", t)
	(&t).Set(2)
	fmt.Printf("%#v\n", t)
	fmt.Printf("%T, %T, %T, %T\n", t.Get, T.Get, (&t).Set, (*T).Set)
	T.Get(t)
	fmt.Printf("%#v\n", t)
	(*T).Set(&t, 10)
	fmt.Printf("%#v\n", t)
	f1 := (*T).Set
	f2 := T.Get
	fmt.Printf("%T, %T\n", f1, f2)

	f3 := func() {
	}
	f4 := func(a int, b int) int {
		return 0
	}
	fmt.Printf("%T, %T\n", f3, f4)
	f5 := t.Get
	fmt.Printf("%T, %T\n", f5, f4)
}
