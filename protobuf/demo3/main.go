package main

import (
	"fmt"

	"github.com/betNevS/code-examples/protobuf/demo3/tutorialpb"
	"github.com/golang/protobuf/proto"
)

func main() {
	p := &tutorialpb.Person{
		Id:    1234,
		Name:  "John",
		Email: "111@37.com",
		Phones: []*tutorialpb.Person_PhoneNumber{
			{Number: "333-222", Type: tutorialpb.Person_HOME},
		},
	}
	b, err := proto.Marshal(p)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(b)
}
