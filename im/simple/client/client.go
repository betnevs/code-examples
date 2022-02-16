package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/url"
	"strconv"
	"time"

	"github.com/betNevS/code-examples/im/simple/common"
	"github.com/gorilla/websocket"
)

type Sender struct {
	conn *websocket.Conn
	send chan common.Message
}

func (s *Sender) sendMessage(index int, str string) {
	msg := common.Message{
		Token:   strconv.Itoa(index),
		Content: str,
	}
	s.send <- msg
}

func (s *Sender) loopSendMessage() {
	for {
		m := <-s.send
		if err := s.conn.WriteJSON(m); err != nil {
			log.Println("write error:", err)
		}
		log.Println("send msg:", m)
	}
}

func createClients(num int) []*Sender {
	var clients []*Sender
	u := url.URL{Scheme: "ws", Host: "127.0.0.1:8081", Path: "/chatroom"}
	for i := 0; i < num; i++ {
		c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
		if err != nil {
			panic(err)
		}
		sender := &Sender{
			conn: c,
			send: make(chan common.Message, 10),
		}
		go sender.loopSendMessage()
		clients = append(clients, sender)
	}
	return clients
}

func process(clients []*Sender) {
	rand.Seed(time.Now().UnixNano())
	n := len(clients) - 1
	flag := 0

	for {
		time.Sleep(time.Millisecond)
		randIndex := rand.Intn(n)
		go func() {
			clients[randIndex].sendMessage(randIndex, strconv.Itoa(flag))
			flag++
		}()
	}
}

func main() {
	clients := createClients(10)
	go readmsg(clients)
	process(clients)
}

func readmsg(clients []*Sender) {
	for {
		for _, c := range clients {
			resp := common.Message{}
			if err := c.conn.ReadJSON(&resp); err != nil {
				fmt.Println("get resp error:", resp)
				continue
			}
			fmt.Println("get resp:", resp)
		}
	}
}
