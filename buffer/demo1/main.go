package main

import (
	"bytes"
	"compress/gzip"
	"fmt"
)

func main() {
	a := bytes.NewBuffer(nil)
	a.WriteString("abc")
	fmt.Println("a:", string(a.Bytes()))

	b := bytes.NewBuffer(nil)
	b.Write([]byte("xyz"))
	fmt.Println("b:", string(b.Bytes()))

	a.Write([]byte("123"))
	fmt.Println("a:", string(a.Bytes()))

	a.WriteTo(b)
	fmt.Println("b:", string(b.Bytes()))
	fmt.Println("a:", string(a.Bytes()))
	b.WriteTo(a)
	fmt.Println("b:", string(b.Bytes()))
	fmt.Println("a:", string(a.Bytes()))

	a.Reset()
	gz := gzip.NewWriter(a)
	gz.Write([]byte("yangjie"))
	fmt.Println("a:", string(a.Bytes()))
	gz.Close()
	fmt.Println("a:", string(a.Bytes()))

}
