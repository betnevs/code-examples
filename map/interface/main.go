package main

import (
	"fmt"
)

func main() {
	a := make(map[string]interface{})
	b := map[string]string{"aa": "bb"}
	a["one"] = b
	fmt.Println(a)
	fmt.Println(a["one"])
	delete(a, "one")
	fmt.Println(a)
	//fmt.Println(a["one"]["aa"])
	var x int = 1
	d := &x
	fmt.Printf("%p, %p\n", d, &x)
}
