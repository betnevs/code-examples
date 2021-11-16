package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func String2Bytes(s string) []byte {
	sh := (*reflect.StringHeader)(unsafe.Pointer(&s))
	bh := reflect.SliceHeader{
		Data: sh.Data,
		Len:  sh.Len,
		Cap:  sh.Len,
	}
	return *(*[]byte)(unsafe.Pointer(&bh))
}

func Bytes2String(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

func main() {
	a := []byte{1, 2, 3}
	b := a
	fmt.Println(a, b)
	fmt.Printf("a = %p, &a = %p\n", a, &a)
	fmt.Printf("b = %p, &b = %p\n", b, &b)
	a[0] = 99
	fmt.Println(a, b)
	fmt.Printf("a = %p, &a = %p\n", a, &a)
	fmt.Printf("b = %p, &b = %p\n", b, &b)
}
