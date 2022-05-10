package main

import (
	"fmt"
	"reflect"
)

type T struct {
}

func (t T) M1() {}

func (T) M2() {}

type S T

func main() {
	var t T
	dumpMethodSet(t)
	t.M1()

	var s S
	dumpMethodSet(s)

	var ss *S
	dumpMethodSet(ss)

}

func dumpMethodSet(i interface{}) {
	dynTyp := reflect.TypeOf(i)

	if dynTyp == nil {
		fmt.Printf("there is no dynamic type\n")
		return
	}

	n := dynTyp.NumMethod()
	if n == 0 {
		fmt.Printf("%s's method set is empty!\n", dynTyp)
		return
	}

	fmt.Printf("%s's method set:\n", dynTyp)
	for j := 0; j < n; j++ {
		fmt.Println("-", dynTyp.Method(j).Name)
	}
	fmt.Printf("\n")
}
