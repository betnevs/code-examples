package main

import (
	"fmt"
	"time"

	"github.com/betNevS/code-examples/workerpool"
)

func main() {
	p := workerpool.New(5)

	for i := 0; i < 10; i++ {
		err := p.Schedule(func() {
			time.Sleep(time.Second * 3)
		})

		if err != nil {
			fmt.Println("task: ", i, "err: ", err)
		}

	}

	p.Free()
}
