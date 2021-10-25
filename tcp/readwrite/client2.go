package main

import (
	"log"
	"net"
	"os"
	"time"
)

func main() {
	if len(os.Args) <= 1 {
		return
	}
	log.Println("begin dial..., args: ", os.Args)
	conn, err := net.Dial("tcp", ":8888")
	if err != nil {
		log.Println("dial error: ", err)
		return
	}
	defer conn.Close()
	log.Println("dial ok")
	time.Sleep(2 * time.Second)
	data := os.Args[1]
	conn.Write([]byte(data))
	time.Sleep(5 * time.Second)

}
