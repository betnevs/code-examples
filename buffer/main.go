package main

import (
	"bytes"
	"fmt"
)

func main() {
	bb := []byte("abcdef1g")
	buf1 := bytes.NewBuffer(bb)
	fmt.Println(bb, buf1.Bytes())
}
