package main

import (
	"context"
	"fmt"
	"time"
)

func main() {

	fmt.Println(time.Now().Unix())
	d := time.Now().Add(5 * time.Second)
	fmt.Println(d.Unix())

	ctx, cancel := context.WithDeadline(context.Background(), d)
	defer cancel()

	select {
	case <-time.After(time.Second):
		dd, ok := ctx.Deadline()
		fmt.Println("game over: ", dd.Unix(), ok)
	case <-ctx.Done():
		fmt.Println(ctx.Err())
	}

}
