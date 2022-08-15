package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	PhoneNum string
	School   string
	City     string
	Name     string
	Sex      string
	Age      int
}

func main() {
	p1 := Person{
		Name:     "tom",
		Sex:      "male",
		Age:      10,
		PhoneNum: "1000000",
		School:   "spb-kindergarden",
		City:     "cq",
	}

	rv := reflect.ValueOf(p1)
	rt := reflect.TypeOf(p1)
	if rv.Kind() == reflect.Struct {
		for i := 0; i < rt.NumField(); i++ {
			//按顺序遍历
			fmt.Printf("field:%+v,value:%+v\n", rt.Field(i).Name, rv.Field(i))
		}
	}
}
