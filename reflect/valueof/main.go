package main

import (
	"fmt"
	"reflect"
)

func main() {
	i := 10
	iv := reflect.ValueOf(&i)
	ivv := iv.Elem()
	ivv.SetInt(99)
	fmt.Println(i)

	i2 := 11
	i2v := reflect.ValueOf(&i2)
	i2v.Elem().SetInt(100)
	fmt.Println(i2)
}
