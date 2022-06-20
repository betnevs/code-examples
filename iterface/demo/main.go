package main

import "fmt"

func cal() interface{} {
	return "yyy"
}

func main() {
	var i interface{}
	i = 1
	fmt.Println(i)
	i = cal
	fmt.Println(i)
	i = cal()
	fmt.Println(i)
	c := func() interface{} {
		return "bb"
	}

	fmt.Println(c())
	i = func() interface{} {
		return "dd"
	}
	fmt.Println(i.(func() interface{})())

}
