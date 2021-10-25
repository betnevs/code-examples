package main

import (
	"log"
	"net"
)

func main() {
	l, err := net.Listen("tcp", ":8888")
	if err != nil {
		log.Println(err)
		return
	}
	for {
		c, err := l.Accept()
		if err != nil {
			log.Println(err)
			break
		}
		log.Println("accept a new conn")
		go handleConn(c)
	}
}

func handleConn(conn net.Conn) {
	defer conn.Close()
	var buf = make([]byte, 20)
	log.Println("start to read from conn")
	n, err := conn.Read(buf)
	if err != nil {
		log.Println("conn read error:", err)
	} else {
		log.Printf("read %d bytes, content is %s\n", n, string(buf[:n]))
	}
	n, err = conn.Write(buf)
	if err != nil {
		log.Println("conn write error:", err)
	} else {
		log.Printf("write %d bytes, content is %s\n", n, string(buf[:n]))
	}
}
