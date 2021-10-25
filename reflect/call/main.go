package main

import (
	"errors"
	"fmt"
	"reflect"
)

type T struct {
}

func (t *T) Add(a, b int) {
	fmt.Printf("a+b is %+v\n", a+b)
}

func (t *T) AddRetErr(a, b int) (int, error) {
	if a+b < 10 {
		return a + b, errors.New("total lt 10")
	}
	return a + b, nil
}

func main() {
	funcName := "Add"
	typeT := &T{}
	a := reflect.ValueOf(1)
	b := reflect.ValueOf(2)
	in := []reflect.Value{a, b}
	reflect.ValueOf(typeT).MethodByName(funcName).Call(in)
	fmt.Println(reflect.ValueOf(typeT).Kind())

	funcName2 := "AddRetErr"
	ret := reflect.ValueOf(typeT).MethodByName(funcName2).Call(in)
	for i := 0; i < len(ret); i++ {
		fmt.Printf("ret index: %+v, type: %+v, value:%+v\n", i, ret[i].Kind(), ret[i].Interface())
	}
}
