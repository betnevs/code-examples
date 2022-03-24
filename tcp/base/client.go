package main

import (
	"fmt"
	"net"
	"time"

	"github.com/betNevS/code-examples/tcp/base/proto"
)

func main() {
	conn, err := net.Dial("tcp", ":8999")
	if err != nil {
		fmt.Println("dial failed, err:", err)
		return
	}
	defer conn.Close()
	for i := 0; i < 100000; i++ {
		msg := `hello, hello, hello`
		data, err := proto.Encode(msg)
		if err != nil {
			fmt.Println("encode msg, err:", err)
			return
		}
		conn.Write(data)
		time.Sleep(time.Second)
	}
}
