package main

import (
	"log"
	"net"
	"time"
)

func main() {
	l, err := net.Listen("tcp", ":8888")
	if err != nil {
		log.Println("listen error: ", err)
		return
	}
	for {
		c, err := l.Accept()
		if err != nil {
			log.Println("accept error:", err)
			break
		}
		log.Println("accept a new connection")
		go handleConn4(c)
	}
}

func handleConn4(c net.Conn) {
	defer c.Close()
	for {
		time.Sleep(10 * time.Second)
		buf := make([]byte, 65536)
		log.Println("start to read from conn")
		c.SetReadDeadline(time.Now().Add(time.Microsecond * 10))
		n, err := c.Read(buf)
		if err != nil {
			log.Printf("conn read %d bytes, error: %s\n", n, err)
			if nerr, ok := err.(net.Error); ok && nerr.Timeout() {
				continue
			}
			return
		}
		log.Printf("read %d bytes, content is %s\n", n, string(buf[:n]))
	}
}
