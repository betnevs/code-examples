package main

import "fmt"

type SS struct {
	name string
}

type Child struct {
	cs SS
}

func (c Child) Cry2() {
	fmt.Println("struct cry2")
}

func (c *Child) Cry() {
	fmt.Println("point cry")
}

type Person struct {
	*Child
}

func main() {
	c := new(Child)
	c.cs.name = "aaa"
	fmt.Println(c.cs)
}
