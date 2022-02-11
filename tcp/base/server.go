package main

import (
	"bufio"
	"fmt"
	"github.com/betNevS/code-examples/tcp/base/proto"
	"io"
	"net"
)

type Person struct {
	name string
	age  int
}

func main() {
	lis, err := net.Listen("tcp", ":8000")
	if err != nil {
		fmt.Println("listen err:", err)
		return
	}
	defer lis.Close()
	for {
		conn, err := lis.Accept()
		if err != nil {
			fmt.Println("accept failed, err:", err)
			continue
		}
		go process(conn)
	}
}

func process(conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewReader(conn)
	for {
		msg, err := proto.Decode(reader)
		if err == io.EOF {
			fmt.Println("get EOF", err)
			return
		}
		if err != nil {
			fmt.Println("decode msg failed, err:", err)
			return
		}
		fmt.Println("get client content:", msg)
	}
}
