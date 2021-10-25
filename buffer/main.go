package main

import (
	"bytes"
	"fmt"
)

func main() {
	buf1 := bytes.NewBufferString("abcdefg")
	p := make([]byte, 2)
	buf1.Read(p)
	fmt.Println(p)
	fmt.Println(buf1)
	buf1.Truncate(2)
	fmt.Println(buf1)
}
