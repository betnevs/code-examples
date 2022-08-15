package main

import (
	"fmt"
	"reflect"
)

type Foo struct {
	A int    `myTag:"value1" json:"a"`
	B string `myTag:"value2"`
}

func main() {
	var i int
	iType := reflect.TypeOf(i)
	fmt.Println(iType.Name())
	f := Foo{}
	ft1 := reflect.TypeOf(f)
	fmt.Println(ft1.Name())
	xpt := reflect.TypeOf(&i)
	fmt.Println(xpt.Name())
	fmt.Println(xpt.Kind())
	fmt.Println(xpt.Elem().Name())
	fmt.Println(xpt.Elem().Kind())

	var ff Foo
	ft := reflect.TypeOf(ff)
	for i := 0; i < ft.NumField(); i++ {
		curField := ft.Field(i)
		fmt.Println(curField.Name, curField.Type.Name(), curField.Tag.Get("myTag"), ",,,,", curField.Tag.Get("json"))
	}
}
