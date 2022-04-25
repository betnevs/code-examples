package main

import (
	"fmt"
)

func main() {
	a := make(map[string]interface{})
	b := map[string]string{"aa": "bb"}
	a["one"] = b
	fmt.Println(a["one"])
	//fmt.Println(a["one"]["aa"])

	var p *int
	fmt.Println(p)
	fmt.Println(*p)
}
