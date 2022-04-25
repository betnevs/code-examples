package main

import "fmt"

func main() {
	ch := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			ch <- 0
		}
		close(ch)
	}()

	//for v := range ch {
	//	fmt.Println(v)
	//}
	for {
		v, ok := <-ch
		fmt.Println(v, ok)
		//if !ok {
		//	break
		//}
	}
}
