package main

import (
	"bufio"
	"fmt"
	"io"
	"net"

	"github.com/betNevS/code-examples/tcp/base/proto"
)

func main() {
	lis, err := net.Listen("tcp", ":9090")
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
