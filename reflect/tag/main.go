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
	t := TagTest{Name: "tom", Age: 10}
	rtt := reflect.TypeOf(t)
	for i := 0; i < rtt.NumField(); i++ {
		field := rtt.Field(i)
		if json, ok := field.Tag.Lookup("json"); ok {
			fmt.Printf("tag is %+v, value is %+v\n", json, field.Tag.Get("json"))
		}
	}
}
