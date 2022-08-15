package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "www.baidu.com:http")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(conn.LocalAddr().String())
}
