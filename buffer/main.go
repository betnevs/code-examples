package main

import (
	"bytes"
	"fmt"
)

func main() {
	bb := []byte("abcdefg")
	buf1 := bytes.NewBuffer(bb)
	fmt.Println(bb, buf1.Bytes())
}
