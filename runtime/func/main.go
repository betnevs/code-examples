package main

import (
	"fmt"
	"runtime"
	"strings"
	"time"
)

func main() {
	a := "aa bb cc"
	fmt.Println(strings.Fields(a))
	go func() {
		time.Sleep(time.Second)
	}()
	Foo()

	time.Sleep(2 * time.Second)
}

func Foo() {
	fmt.Printf("I am %s, %s call me?\n", printMyName(), printCallerName())
	Bar()
}

func Bar() {
	fmt.Printf("I am %s, %s call me?\n", printMyName(), printCallerName())
	DumpStacks()
}

func printMyName() string {
	pc, _, _, _ := runtime.Caller(1)
	return runtime.FuncForPC(pc).Name()
}

func printCallerName() string {
	pc, _, _, _ := runtime.Caller(2)
	return runtime.FuncForPC(pc).Name()
}

func trace() {
	pc := make([]uintptr, 10)
	n := runtime.Callers(0, pc)
	fmt.Println("n:", n)
	for i := 0; i < n; i++ {
		f := runtime.FuncForPC(pc[i])
		file, line := f.FileLine(pc[i])
		fmt.Printf("%s:%d %s\n", file, line, f.Name())
	}
}

func trace2() {
	pc := make([]uintptr, 10)
	n := runtime.Callers(0, pc)
	fmt.Println("n:", n, "pc:", pc)
	frames := runtime.CallersFrames(pc[:n])
	for {
		frame, more := frames.Next()
		fmt.Printf("%s:%d %s\n", frame.File, frame.Line, frame.Function)
		if !more {
			break
		}
	}
}

func DumpStacks() {
	buf := make([]byte, 1000)
	buf = buf[:runtime.Stack(buf, true)]
	fmt.Printf("dump:%s", buf)
}
