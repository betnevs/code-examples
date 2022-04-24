package main

import (
	"fmt"
)

func main() {
	a := make(chan int)
	b := make(chan int)
	out := make(chan int)
	go func() {
		for i := 0; i < 5; i++ {
			a <- i
		}
		close(a)
	}()

	go func() {
		for i := 5; i < 1000; i++ {
			b <- i
		}
		close(b)
	}()
	go func() {
		merge(out, a, b)
	}()
	for v := range out {
		fmt.Print(v, " ")
	}
}

func merge(out chan<- int, a, b <-chan int) {
	for a != nil || b != nil {
		select {
		case v, ok := <-a:
			if !ok {
				a = nil
				fmt.Println("a is now closed")
				continue
			}
			out <- v
		case v, ok := <-b:
			if !ok {
				b = nil
				continue
			}
			out <- v
		}
	}
	close(out)
}
