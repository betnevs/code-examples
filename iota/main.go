package main

import (
	"fmt"
	"log"
	"os"
)

const (
	x = iota
	y
	z
)

const (
	_  = iota
	KB = 1 << (10 * iota)
	MB
	GB
	xx = iota
)

const (
	bf   = iota
	_, _ = iota, iota * 10
	a, b
	c, d
	xd1, xd2, xd3 = iota, iota, iota
)

const (
	a1 = iota
	b1
	c1 = 100
	d1
	e1 = iota
	f1
	g1 = iota
)

const (
	a2         = iota
	b2 float32 = iota
	c2         = iota
)

type color byte

const (
	black color = iota
	red
	blue
)

func test(c color) {
	println(c)
}

func main() {
	log.Println("aaa: ", "aaaa")
	fmt.Println("aaa:", "aaaa")
	os.Exit(1)
	fmt.Println(x, y, z)
	fmt.Println(KB, MB, GB, xx)
	fmt.Println(a, b, c, d, xd1, xd2, xd3, bf)
	fmt.Println(a1, b1, c1, d1, e1, f1, g1)
	fmt.Println(a2, b2, c2)
	fmt.Printf("%T, %T, %T", a2, b2, c2)

	test(red)
	test(100)
	jj := 2
	fmt.Printf("%T\n", jj)
	//test(jj)  // occur error
}
