package main

import (
	"fmt"
	"runtime"
)

func Foo() {
	fmt.Printf("i am %s, %s call me\n", printMyName(), printCallerName())
	Bar()
}

func Bar() {
	fmt.Printf("i am %s, %s call me\n", printMyName(), printCallerName())
	trace2()
}

func printMyName() string {
	pc, _, _, _ := runtime.Caller(0)
	return runtime.FuncForPC(pc).Name()
}

func printCallerName() string {
	pc, _, _, _ := runtime.Caller(1)
	return runtime.FuncForPC(pc).Name()
}

func trace() {
	pc := make([]uintptr, 10)
	n := runtime.Callers(0, pc)
	for i := 0; i < n; i++ {
		f := runtime.FuncForPC(pc[i])
		file, line := f.FileLine(pc[i])
		fmt.Printf("%s:%d %s\n", file, line, f.Name())
	}
}

func DumpStacks() {
	buf := make([]byte, 16384)
	buf = buf[:runtime.Stack(buf, true)]
	fmt.Println(string(buf))
}

func trace2() {
	pc := make([]uintptr, 10)
	n := runtime.Callers(0, pc)
	frames := runtime.CallersFrames(pc[:n])
	for {
		frame, more := frames.Next()
		fmt.Printf("%s:%d %s\n", frame.File, frame.Line, frame.Function)
		if !more {
			break
		}
	}
}

func main() {
	Foo()
}
