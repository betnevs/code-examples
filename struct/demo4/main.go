package main

import "fmt"

//Person 这里抽象人的特性为属性
type Person struct {
	Name string
	Age  int32
}

type LittlePerson struct {
	Person
	Name string
}

//Mover  定义行为
type Mover interface {
	Move()
	Eat()
}

func (p *Person) Move() {
	fmt.Println("person move.")
}

func (l *LittlePerson) Move() {
	fmt.Println("littlePerson move.")
}

func (l *LittlePerson) Eat() {
	fmt.Println("littlePerson eat.")
}

func main() {
	person := Person{"bobo", 18}
	//person.Move()
	little := LittlePerson{
		person, "littlePerson",
	}
	var m Mover = &little
	m.Move()
	//todo 后续完善
}
