package main

import "fmt"

type T1 int
type T2 T1
type T3 string

func main() {
	var n1 T1
	var n2 T2 = 5
	fmt.Println(n1, n2)

	//var s T3 = "hello"
	//n1 = T1(s) // 错误：cannot convert s (type T3) to type T1
}
