package main

import (
	"fmt"
	"reflect"
)

func IsNil(i interface{}) {
	if i != nil {
		if reflect.ValueOf(i).IsNil() {
			fmt.Println("i is nil")
			return
		}
		fmt.Println("i isn't nil")
	}
	fmt.Println("i is nil")
}

func main() {
	var sl []string
	IsNil(sl)  // i is nil
	IsNil(nil) // i is nil
	IsNil(1)   // panic
}
