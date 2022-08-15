package main

import (
	"fmt"
	"sync"
)

func main() {

	var s = "ab"

	var wg sync.WaitGroup
	for i := 0; i < 1000000; i++ {

		wg.Add(1)
		go func() {
			defer wg.Done()
			_ = s
		}()

		wg.Add(1)
		go func() {
			defer wg.Done()
			s = "xyzw"
		}()

		if s != "ab" && s != "xyzw" {
			fmt.Println("error:", s)
		}
	}
	wg.Wait()

	fmt.Println(s)

}
