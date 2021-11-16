package main

import (
	"fmt"
)

type People struct {
}

func (p *People) say() {
	fmt.Println("people say")
}

type Pp People

var x = 0x10

const y = 0x20

func main() {
	fmt.Println(x, &x)
	//fmt.Println(y, &y) // occur error
	p := new(People)
	p.say()
	//xp := new(Pp)  // occur error
	//xp.say()
}
