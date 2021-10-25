package main

import "fmt"

const (
	CommandConn = iota + 0x01
	CommandSubmit
)

type AA struct {
}

func main() {
	count := 100000001
	id := fmt.Sprintf("%08d", count)
	fmt.Println(id)
}
