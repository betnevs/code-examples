package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	ch := make(chan os.Signal, 1)
	finish := make(chan struct{}, 1)

	go func() {
		signal.Notify(ch, syscall.SIGINT)
		for {
			select {
			case <-ch:
				finish <- struct{}{}
			default:
				time.Sleep(time.Second)
				fmt.Println("sleeping")

			}
		}
	}()

	<-finish
	fmt.Println("end")

}
