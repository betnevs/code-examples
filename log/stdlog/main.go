package main

import (
	"fmt"
	"runtime"
)

func main() {
	Foo()
	buf := make([]byte, 4096)
	n := runtime.Stack(buf[:], false)
	fmt.Println(string(buf[:n]))
	n1 := runtime.Stack(buf[:], false)
	fmt.Println(string(buf[:n1]))
}
func Foo() {
	fmt.Printf("我是 %s, %s 在调用我!\n", printMyName(), printCallerName())
	Bar()
}
func Bar() {
	fmt.Printf("我是 %s, %s 又在调用我!\n", printMyName(), printCallerName())
}
func printMyName() string {
	pc, _, _, _ := runtime.Caller(1)
	return runtime.FuncForPC(pc).Name()
}
func printCallerName() string {
	pc, _, _, _ := runtime.Caller(2)
	return runtime.FuncForPC(pc).Name()
}
