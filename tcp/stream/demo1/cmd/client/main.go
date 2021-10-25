package main

import (
	"fmt"
	"log"
	"net"
	"sync"
	"time"

	"github.com/betNevS/code-examples/tcp/stream/demo1/pkg/frame"
	"github.com/betNevS/code-examples/tcp/stream/demo1/pkg/packet"
	"github.com/lucasepe/codename"
)

func main() {
	var wg sync.WaitGroup
	var num int = 5
	wg.Add(num)
	for i := 0; i < num; i++ {
		go func() {
			defer wg.Done()
			startClient()
		}()
	}
	wg.Wait()
}

func startClient() {
	conn, err := net.Dial("tcp", ":8888")
	if err != nil {
		log.Println("dial err:", err)
		return
	}
	defer conn.Close()
	log.Println("dial ok")
	rng, err := codename.DefaultRNG()
	if err != nil {
		panic(err)
	}
	frameCodec := frame.NewMyFrameCodec()
	var count int
	go func() {
		for {
			ackFramePayload, err := frameCodec.Decode(conn)
			if err != nil {
				panic(err)
			}
			p, err := packet.Decode(ackFramePayload)
			submitAck, ok := p.(*packet.SubmitAck)
			if !ok {
				panic("not submit ack")
			}
			fmt.Printf("the result of submit ack[%s] is %d\n", submitAck.ID, submitAck.Result)
		}
	}()
	for {
		count++
		id := fmt.Sprintf("%08d", count)
		payload := codename.Generate(rng, 4)
		s := &packet.Submit{
			ID:      id,
			Payload: []byte(payload),
		}
		framePayLoad, err := packet.Encode(s)
		if err != nil {
			panic(err)
		}
		fmt.Printf("send sumbit id = %s, payload = %s, frame length = %d\n", s.ID, s.Payload, len(framePayLoad)+4)
		err = frameCodec.Encode(conn, framePayLoad)
		if err != nil {
			panic(err)
		}
		time.Sleep(time.Second)
	}
}
