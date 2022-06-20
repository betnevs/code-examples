package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type User struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Address string `json:"address"`
}

func main() {
	fmt.Println("marshal......")
	u := User{
		Name:    "yangjie",
		Age:     18,
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
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%#v\n", u1)
}
