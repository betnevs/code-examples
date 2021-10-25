package main

import (
	"log"
	"net"
)

func main() {
	l, err := net.Listen("tcp", ":8888")
	if err != nil {
		log.Println("listen error:", err)
		return
	}
	for {
		c, err := l.Accept()
		if err != nil {
			log.Println("accept error:", err)
			break
		}
		log.Println("accept a new conn")
		go handleConn(c)
	}
}

func handleConn(conn net.Conn) {
	defer conn.Close()
	for {
		buf := make([]byte, 10)
		log.Println("start to read from conn")
		n, err := conn.Read(buf)
		if err != nil {
			log.Println("conn read error:", err)
			return
		}
		log.Printf("read %d bytes, content is %v\n", n, string(buf[:n]))
	}
}
