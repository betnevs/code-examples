package main

import (
	"log"
	"net"
	"time"
)

func main() {
	l, err := net.Listen("tcp", ":8888")
	if err != nil {
		log.Println("err listen:", err)
	}
	defer l.Close()
	log.Println("listen ok")

	var i int
	for {
		time.Sleep(10 * time.Second)
		//if _, err := l.Accept(); err != nil {
		//	log.Println("accept err:", err)
		//	break
		//}
		i++
		log.Printf("%d: accept a new conn\n", i)
	}
}
