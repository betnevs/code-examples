package main

import (
	"fmt"

	"github.com/lesismal/nbio"
)

func main() {
	g := nbio.NewGopher(nbio.Config{
		Network: "tcp",
		Addrs:   []string{"localhost:8080"},
	})

	defer g.Stop()

	g.OnData(func(c *nbio.Conn, data []byte) {
		fmt.Println("server get:", string(data))
		c.Write(data)
	})

	err := g.Start()
	if err != nil {
		panic(err)
	}

	g.Wait()
}
