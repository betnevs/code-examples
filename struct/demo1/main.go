package main

import "fmt"

type Sayer interface {
	Say() string
}

type People struct {
	Sayer
	//xp *pp
}

type pp struct {
}

func (p *pp) Say() string {
	return "this is pp"
}

//func (p *People) Say() string {
//	return "sss"
//}

func (p *People) Walk() string {
	return "walk"
}

func main() {

	s := &People{}
	fmt.Println(s.Walk())
	fmt.Println(s.Say())
	//d := &People{}
	//d.Sayer = &pp{}
	//fmt.Println(d.Say())
	//fmt.Println(d.Sayer.Say())

}
