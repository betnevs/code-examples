package main

import (
	"fmt"
	"unicode/utf8"
)

const (
	a = iota
	b = 2
	c = 1
	d = iota
)

func main() {
	fmt.Println(a, b, c, d)
	var r rune = 0x4e2d
	fmt.Printf("%c\n", r)
	buf := make([]byte, 3)
	utf8.EncodeRune(buf, r)
	fmt.Println(buf)
	fmt.Printf("%#X\n", buf)

	//var s = "中国人"
	var s = "abc"
	fmt.Println(utf8.RuneCountInString(s))
	fmt.Printf("%#x\n", s[0]) // 0xe4：字符“中” utf-8编码的第一个字节

	//const m = 1222222
	//var k int8 = 3
	//j := m + k
	//fmt.Println(j)
}
