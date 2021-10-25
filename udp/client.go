package main

import (
	"fmt"
	"net"
)

func main() {
	sock, err := net.DialUDP("udp", nil, &net.UDPAddr{
		IP:   net.IPv4(0, 0, 0, 0),
		Port: 3000,
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	defer sock.Close()
	sendData := []byte("hello server")
	_, err = sock.Write(sendData)
	if err != nil {
		fmt.Println(err)
		return
	}
	data := make([]byte, 4096)
	n, rAddr, err := sock.ReadFrom(data)

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("recv: %v, addr: %v, count: %v\n", string(data[:n]), rAddr, n)
}
