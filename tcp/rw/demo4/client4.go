package main

import (
	"log"
	"net"
)

func main() {
	log.Println("begin dial...")
	conn, err := net.Dial("tcp", ":8888")
	if err != nil {
		log.Println("dial error:", err)
		return
	}
	defer conn.Close()
	log.Println("dial ok")
	data := make([]byte, 65536)
	data[0] = 'a'
	data[1] = 'b'
	data[65535] = 'x'
	conn.Write(data)
}
