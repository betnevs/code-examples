package main

import (
	"fmt"
	"net"
)

func main() {
	lis, err := net.ListenUDP("udp", &net.UDPAddr{
		IP:   net.IPv4(0, 0, 0, 0),
		Port: 3000,
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	defer lis.Close()
	for {
		var data [1024]byte
		n, addr, err := lis.ReadFromUDP(data[:])
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Printf("data: %v, addr: %v, count: %v\n", string(data[:n]), addr, n)
		_, err = lis.WriteToUDP(data[:n], addr)
		if err != nil {
			fmt.Println("write err:", err)
			continue
		}
	}
}
