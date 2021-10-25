package main

import (
	"fmt"
	"reflect"
)

type TagTest struct {
	Name string `json:"name_json"`
	Age  int    `json:"age_json"`
}

func main() {
	ta := 10
	vta := reflect.ValueOf(&ta)
	if vta.Elem().CanSet() {
		vta.Elem().Set(reflect.ValueOf(11))
	}
	fmt.Println("vta", vta, vta.Elem(), ta)

	ts := []int{1, 2, 3}
	tsV := reflect.ValueOf(ts)
	if tsV.Index(0).CanSet() {
		tsV.Index(0).Set(reflect.ValueOf(100))
	}
	fmt.Println(ts)

	t1 := TagTest{}
	tV := reflect.ValueOf(t1)

	//结构体指针
	t1V := reflect.ValueOf(&t1)

	fmt.Printf("tV:%+v\n", tV)
	for i := 0; i < tV.NumField(); i++ {
		val := tV.Field(i)
		t1V.Elem().Field(i).Set(val)
	}
	//a := 1
	//t1V.Field(0).Set(reflect.ValueOf(a))
	fmt.Printf("t1 is %+v\n", t1)
}
