package main

import (
	"log"
	"net"
	"time"
)

func main() {
	conn, err := net.Dial("tcp", ":8888")
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close()
	data := make([]byte, 65536)
	time.Sleep(2 * time.Second)
	conn.Write(data)
	time.Sleep(1000 * time.Second)
}
