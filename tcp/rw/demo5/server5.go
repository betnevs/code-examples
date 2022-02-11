package main

import (
	"log"
	"net"
	"time"
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
		log.Println("accept a new connection")
		go handleConn(c)
	}
}

func handleConn(c net.Conn) {
	defer c.Close()
	time.Sleep(10 * time.Second)
	for {
		time.Sleep(5 * time.Second)
		buf := make([]byte, 60000)
		log.Println("start to read from conn")
		n, err := c.Read(buf)
		if err != nil {
			log.Printf("conn read %d bytes, error: %s", n, err)
			if nerr, ok := err.(net.Error); ok && nerr.Timeout() {
				continue
			}
			break
		}
		log.Printf("read %d bytes, content is %s\n", n, string(buf[:n]))
	}
}
