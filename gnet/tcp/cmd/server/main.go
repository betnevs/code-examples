package main

import (
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/betNevS/code-examples/gnet/tcp/pkg/packet"

	"github.com/betNevS/code-examples/gnet/tcp/pkg/frame"

	"github.com/panjf2000/gnet"
	"github.com/panjf2000/gnet/pool/goroutine"
)

type customCodecServer struct {
	*gnet.EventServer
	addr       string
	multicore  bool
	async      bool
	codec      gnet.ICodec
	workerPool *goroutine.Pool
}

func (cs *customCodecServer) OnInitComplete(srv gnet.Server) (action gnet.Action) {
	log.Printf("custom codec server is listening on %s (multi-cores: %t, loops: %d)\n", srv.Addr.String(), srv.Multicore, srv.NumEventLoop)
	return
}

func (cs *customCodecServer) React(frame []byte, c gnet.Conn) (out []byte, action gnet.Action) {
	var p packet.Packet
	var ackFramePayload []byte
	p, err := packet.Decode(frame)
	if err != nil {
		fmt.Println("react: packet decode error:", err)
		action = gnet.Close
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
		ackFramePayload, err = packet.Encode(submitAck)
		if err != nil {
			fmt.Println("handleConn: packet encode error:", err)
			action = gnet.Close
			return
		}
		out = []byte(ackFramePayload)
		return
	default:
		return nil, gnet.Close
	}
}

func customCodecServe(addr string, multicore bool, async bool, codec gnet.ICodec) {
	var err error
	codec = &frame.Frame{}
	cs := &customCodecServer{
		addr:       addr,
		multicore:  multicore,
		async:      async,
		codec:      codec,
		workerPool: goroutine.Default(),
	}
	err = gnet.Serve(cs, addr, gnet.WithMulticore(multicore), gnet.WithTCPKeepAlive(5*time.Minute), gnet.WithCodec(codec))
	if err != nil {
		panic(err)
	}
}

func main() {
	var (
		port      int
		multicore bool
	)
	flag.IntVar(&port, "port", 8888, "server port")
	flag.BoolVar(&multicore, "multicore", true, "multicore")
	flag.Parse()
	addr := fmt.Sprintf("tcp://:%d", port)
	customCodecServe(addr, multicore, false, nil)
}
