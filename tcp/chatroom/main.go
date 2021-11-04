package main

import (
	"log"
	"net"
)

func main() {
	s := newServer()
	go s.run()
	addr := ":8888"
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("unable to start server: %s", err.Error())
	}
	defer listener.Close()
	log.Printf("start server on %s", addr)
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("can not accept: %s", err.Error())
			continue
		}
		go s.newClient(conn)
	}
}
