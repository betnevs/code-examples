package v2

import (
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/betNevS/code-examples/im/simple/common"

	"github.com/gorilla/websocket"
)

var (
	room *Room
	ug   = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
)

// 聊天室
type Room struct {
	sync.RWMutex
	Conns []*websocket.Conn
}

func (r *Room) batchSendMessage(message *common.Message) {
	r.Lock()
	for i := 0; i < len(r.Conns); i++ {
		conn := r.Conns[i]
		if err := conn.WriteJSON(message); err != nil {
			log.Println("batch write error:", err)
			conn.Close()
			r.Conns = append(r.Conns[:i], r.Conns[i+1:]...)
			i--
		}
	}
	time.Sleep(100 * time.Second)
	r.Unlock()
}

func chatHandle(w http.ResponseWriter, r *http.Request) {
	// 协议升级
	conn, err := ug.Upgrade(w, r, nil)
	if err != nil {
		log.Println("ws upgrade error:", err)
		return
	}
	defer conn.Close()
	// 连接加入池中
	room.Lock()
	room.Conns = append(room.Conns, conn)
	room.Unlock()
	// 处理聊天逻辑
	for {
		// 获取消息
		message := &common.Message{}
		if err := conn.ReadJSON(message); err != nil {
			log.Println("read json error:", err)
			return
		}
		// 处理消息，防止阻塞
		go room.batchSendMessage(message)
	}

}

func StartChatRoom() {
	log.Println("chatroom start!!!")
	room = &Room{}
	http.HandleFunc("/chatroom", chatHandle)
	http.ListenAndServe(":8081", nil)
}
