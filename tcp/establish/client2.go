package main

import (
	"fmt"
	"log"
	"net"
	"time"
)

func main() {
	var sl []net.Conn
	for i := 1; i < 1000; i++ {
		conn := establishConn(i)
		if conn != nil {
			sl = append(sl, conn)
		}
	}
	fmt.Println("ready to end")
	time.Sleep(1000000 * time.Second)
}

func establishConn(i int) net.Conn {
	conn, err := net.Dial("tcp", ":8888")
	if err != nil {
		log.Printf("%d: dial error: %v", i, err)
	}
	log.Println(i, ":connect to server ok")
	return conn
}
