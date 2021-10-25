package main

import (
	"fmt"
	"net"

	"github.com/betNevS/code-examples/tcp/base/proto"
)

func main() {
	conn, err := net.Dial("tcp", ":8000")
	if err != nil {
		fmt.Println("dial failed, err:", err)
		return
	}
	defer conn.Close()
	for i := 0; i < 20; i++ {
		msg := `hello, hello, hello`
		data, err := proto.Encode(msg)
		if err != nil {
			fmt.Println("encode msg, err:", err)
			return
		}
		conn.Write(data)
	}
}
