package main

import (
	"fmt"
	"path"
)

type myint int

func foo(arr [5]int) {}
func main() {
	var arr1 [5]int
	//var arr2 [6]int
	//var arr3 [5]string
	//var arr4 [5]myint

	foo(arr1) // ok
	//foo(arr2) // 错误：[6]int与函数foo参数的类型[5]int不是同一数组类型
	//foo(arr3) // 错误：[5]string与函数foo参数的类型[5]int不是同一数组类型
	//foo(arr4) // 错误：[5]myint与函数foo参数的类型[5]int不是同一数组类型

	arr := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	sl := arr[3:7:10]
	fmt.Println(sl)

	absolutePath := "cc"
	relativePath := "aa"
	finalPath := path.Join(absolutePath, relativePath)
	fmt.Println(absolutePath, relativePath, finalPath)
}
