package main

import "fmt"

func main() {

	for range []int(nil) {
		fmt.Println("Hello")
	}

	for range map[string]string(nil) {
		fmt.Println("world")
	}

	for i := range (*[5]int)(nil) {
		fmt.Println(i)
	}

	for range chan bool(nil) { // 阻塞在此
		fmt.Println("Bye")
	}
}
