package main

import (
	"fmt"
	"time"

	"github.com/betNevS/code-examples/workerpool"
)

func main() {
	p := workerpool.New(5, workerpool.WithBlock(false), workerpool.WithPreAllocWokers(false))

	time.Sleep(2 * time.Second)

	for i := 0; i < 10; i++ {
		err := p.Schedule(func() {
			time.Sleep(3 * time.Second)
		})
		if err != nil {
			fmt.Printf("task[%d]: error: %s\n", i, err.Error())
		}
	}

	p.Free()
}
