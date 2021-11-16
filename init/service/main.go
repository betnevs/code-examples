package main

import (
	"fmt"

	"github.com/betNevS/code-examples/init/third"
)

func main() {
	third.TestB()
	third.TestA()
}

func init() {
	fmt.Println("this is main")
}
