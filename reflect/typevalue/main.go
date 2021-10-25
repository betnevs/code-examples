package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	Name     string
	Sex      string
	Age      int
	PhoneNum string
}

func main() {
	p1 := Person{
		Name:     "tom",
		Sex:      "male",
		Age:      10,
		PhoneNum: "122",
	}
	rv := reflect.ValueOf(p1)
	rt := reflect.TypeOf(p1)
	if rv.Kind() == reflect.Struct {
		for i := 0; i < rt.NumField(); i++ {
			fmt.Printf("field: %+v, value: %+v\n", rt.Field(i).Name, rv.Field(i))
		}
	}
	rv = reflect.ValueOf(0)
	rt = reflect.TypeOf(0)
	fmt.Println(rv.Kind(), rt.Kind())
}
