package main

import (
	"bufio"
	"fmt"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", ":8080")
	if err != nil {
		fmt.Println("dial failed, err:", err)
		return
	}
	reader := bufio.NewReader(conn)
	b := make([]byte, 512)
	defer conn.Close()
	for {
		n, err := reader.Read(b)
		if err != nil {
			fmt.Println("read err: ", err)
			break
		}
		fmt.Println("receive: ", string(b[:n]))
	}
}
