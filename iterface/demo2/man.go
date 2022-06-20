package main

import "fmt"

type Sayer interface {
	Say()
}

type b struct {
}

func (a *b) Say() {
	fmt.Println("aaa")
}

type Peo struct {
	Sayer
}

func main() {
	p := Peo{
		Sayer: &b{},
	}
	p.Say()
}
