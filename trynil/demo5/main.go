package main

import "fmt"

type P struct {
	name string
}

func returnError() error {
	var p *MyError = nil
	return p
}

type MyError struct {
}

func (m *MyError) Error() string {
	return "test"
}

func main() {
	b := returnError()
	fmt.Printf("%T\n", b)
	fmt.Printf("%v\n", b)

}
