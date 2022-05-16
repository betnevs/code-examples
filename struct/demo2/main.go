package main

import "fmt"

type Animal struct {
	name string
}

func (a *Animal) Name() string {
	return a.name
}

func (a *Animal) Speak() string {
	return fmt.Sprintf("my name is %s", a.name)
}

func (a *Animal) Play() {
	fmt.Println(a.Speak())
}

type Dog struct {
	Animal
	Gender string
}

func (d *Dog) Speak() string {
	return fmt.Sprintf("%v and my gender is %v", d.Animal.Speak(), d.Gender)
}

func main() {
	d := Dog{
		Gender: "male",
		Animal: Animal{
			name: "bbo",
		},
	}

	fmt.Println(d.Name())
	fmt.Println(d.Speak())
	d.Play()
}
