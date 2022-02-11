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
	for {
		time.Sleep(10 * time.Second)
		var buf = make([]byte, 65536)
		log.Println("start to read from connection")
		c.SetReadDeadline(time.Now().Add(5 * time.Microsecond))
		n, err := c.Read(buf)
		if err != nil {
			log.Printf("conn read %d bytes, error:%s", n, err)
			if nerr, ok := err.(net.Error); ok && nerr.Timeout() {
				continue
			}
			return
		}
		log.Printf("read %d bytes, content is %s\n", n, string(buf[:n]))
	}
}
