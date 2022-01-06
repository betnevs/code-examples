package main

import (
	"fmt"
	"github.com/betNevS/code-examples/protobuf/example/studentpb"

	"github.com/golang/protobuf/proto"
	"log"
)

func main() {
	test := &studentpb.Student{
		Name:   "yangjie",
		Male:   true,
		Scores: []int32{99, 98, 97},
	}
	data, err := proto.Marshal(test)
	if err != nil {
		log.Fatal("marshaling error:", err)
	}
	newTest := &studentpb.Student{}
	err = proto.Unmarshal(data, newTest)
	if err != nil {
		log.Fatal("unmarshaling error:", err)
	}
	if test.GetName() != newTest.GetName() {
		log.Fatalf("data mismath %q != %q", test.GetName(), newTest.GetName())
	} else {
		fmt.Println(test.GetName(), newTest.GetName())
	}
}
