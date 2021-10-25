package main

import (
	"context"
	_ "net/http/pprof"
	"time"
)

var datas []string

type People struct {
	Name string
	Age  int
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	dl, ok := ctx.Deadline()
	if ok {
		timeout := time.Now().Add(3 * time.Second)
		dl.Before(timeout)
		timeout = dl
	}
}
