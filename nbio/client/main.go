package main

import (
	"fmt"

	"github.com/lesismal/nbio"
)

func main() {
	g := nbio.NewGopher(nbio.Config{})

	g.OnData(func(c *nbio.Conn, data []byte) {
		fmt.Println("rev:", string(data))
	})

	err := g.Start()
	if err != nil {
		panic(err)
	}
	defer g.Stop()

	c, err := nbio.Dial("tcp", "localhost:8080")
	if err != nil {
		panic(err)
	}

	g.AddConn(c)

	c.Write([]byte("hello yangjie"))
}
