package main

import (
	"bufio"
	"fmt"
	"net"

	"github.com/betNevS/code-examples/tcp/stream/demo1/pkg/packet"

	"github.com/betNevS/code-examples/tcp/stream/demo1/pkg/frame"
)

func main() {
	l, err := net.Listen("tcp", ":8888")
	if err != nil {
		fmt.Println("listen err:", err)
		return
	}
	for {
		c, err := l.Accept()
		if err != nil {
			fmt.Println("accept err:", err)
			break
		}
		go handleConn(c)
	}
}

func handleConn(c net.Conn) {
	defer c.Close()
	frameCodec := frame.NewMyFrameCodec()
	rbuf := bufio.NewReader(c)
	for {
		framePayload, err := frameCodec.Decode(rbuf)
		if err != nil {
			fmt.Println("handleConn: frame decode err:", err)
			return
		}
		ackFramePayload, err := handlePacket(framePayload)
		if err != nil {
			fmt.Println("handleConn: handle packet err:", err)
			return
		}
		err = frameCodec.Encode(c, ackFramePayload)
		if err != nil {
			fmt.Println("handleConn: frame encode err:", err)
			return
		}
	}
}

func handlePacket(framePayload []byte) (ackFramePayload []byte, err error) {
	var p packet.Packet
	p, err = packet.Decode(framePayload)
	if err != nil {
		fmt.Println("handleConn: packet decode err:", err)
		return
	}
	switch p.(type) {
	case *packet.Submit:
		submit := p.(*packet.Submit)
		fmt.Printf("recv submit: id = %s, payload = %s\n", submit.ID, string(submit.Payload))
		submitAck := &packet.SubmitAck{
			ID:     submit.ID,
			Result: 0,
		}
		ackFramePayload, err := packet.Encode(submitAck)
		if err != nil {
			fmt.Println("handleConn: packet encode err:", err)
			return nil, err
		}
		return ackFramePayload, nil
	default:
		return nil, fmt.Errorf("unknown packet type")
	}
}
