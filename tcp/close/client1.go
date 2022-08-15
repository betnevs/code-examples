package main

import (
	"log"
	"net"
	"syscall"
	"time"
)

func main() {
	log.Println("begin dial..1.")
	conn, err := net.Dial("tcp", ":8888")
	if err != nil {
		log.Println("dial error:", err)
		return
	}
	conn.Close()
	log.Println("close ok")
	syscall.SetsockoptInt()
	var buf = make([]byte, 32)
	n, err := conn.Read(buf)
	if err != nil {
		log.Println("read error:", err)
	} else {
		log.Printf("read %d bytes, content is %s\n", n, string(buf[:n]))
	}
	n, err = conn.Write(buf)
	if err != nil {
		log.Println("write error:", err)
	} else {
		log.Printf("write %d bytes, content is %s\n", n, string(buf[:n]))
	}
	time.Sleep(1000 * time.Second)
}
