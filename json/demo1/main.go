package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type User struct {
	Name    string  `json:"name"`
	Age     float64 `json:"age"`
	Address string  `json:"address"`
}

func main() {
	fmt.Println("marshal......")
	u := User{
		Name:    "yangjie",
		Age:     11,
		Address: "hahahah",
	}
	result, err := json.Marshal(u)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("get: ", string(result))

	raw := `{"NaMe":"yangjie111","age":19,"addresS":"hahahah"}`
	u1 := User{}
	err = json.Unmarshal([]byte(raw), &u1)
	fmt.Printf("%T\n", u1.Age)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%#v\n", u1)
}
