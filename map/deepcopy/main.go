package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"log"
)

func main() {
	var err error
	fmt.Println(err)
	src := map[string]interface{}{
		"aa": 1,
		"bb": "ccc",
		"cc": 0.1,
		"dd": map[string]string{"xxx": "xx1"},
		"ff": []interface{}{1, 2, 3, 4},
	}
	des, err := cloneMap(src)
	des["aa"] = 12111
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(src, des)
}

func cloneMap(src map[string]interface{}) (map[string]interface{}, error) {
	var buf bytes.Buffer
	gob.Register(map[string]interface{}{})
	gob.Register(map[string]string{})
	gob.Register([]interface{}{})

	enc := gob.NewEncoder(&buf)
	dec := gob.NewDecoder(&buf)

	err := enc.Encode(src)
	if err != nil {
		return nil, err
	}
	var copy map[string]interface{}
	err = dec.Decode(&copy)
	if err != nil {
		return nil, err
	}
	return copy, nil
}
