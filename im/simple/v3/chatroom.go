package v3

import (
	"fmt"
	"log"
	"net/http"

	"github.com/betNevS/code-examples/im/simple/common"
	"github.com/gorilla/websocket"
)

var (
	chatRoom *Room
	ug       = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
)

type Room struct {
	register   chan *Client
	unregister chan *Client
	send       chan common.Message
	cliPool    map[*Client]bool
}

func chatHandle(w http.ResponseWriter, r *http.Request) {
	conn, err := ug.Upgrade(w, r, nil)
	if err != nil {
		log.Println("upgrade error:", err)
		return
	}
	c := &Client{
		conn: conn,
		send: make(chan common.Message, 5),
	}
	go c.ReadMessage()
	go c.SendMessage()
	chatRoom.register <- c
}

func (room *Room) ProcessTask() {
	log.Println("start process task")
	for {
		select {
		case c := <-room.register:
			log.Println("new user connect")
			room.cliPool[c] = true
		case c := <-room.unregister:
			log.Println("user leave")
			if room.cliPool[c] {
				close(c.send)
				delete(room.cliPool, c)
			}
		case m := <-room.send:
			fmt.Println("room get msg:", m, ", pool size:", len(room.cliPool))
			for c := range room.cliPool {
				fmt.Println("begin send:", m)
				c.send <- m
				fmt.Println("end send:", m)
			}
		}
	}
}

func StartChatRoom() {
	log.Println("start chatroom!!!")
	chatRoom = &Room{
		register:   make(chan *Client),
		unregister: make(chan *Client),
		cliPool:    map[*Client]bool{},
		send:       make(chan common.Message),
	}
	http.HandleFunc("/chatroom", chatHandle)
	go chatRoom.ProcessTask()
	http.ListenAndServe(":8081", nil)
}
