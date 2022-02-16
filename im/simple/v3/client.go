package v3

import (
	"fmt"

	"github.com/betNevS/code-examples/im/simple/common"
	"github.com/gorilla/websocket"
)

// 客户端
type Client struct {
	name string
	conn *websocket.Conn
	send chan common.Message
}

func (c *Client) ReadMessage() {
	i := 0
	for {
		message := &common.Message{}
		if err := c.conn.ReadJSON(message); err != nil {
			fmt.Println("occur err:", err)
			c.conn.Close()
			chatRoom.unregister <- c
			return
		}
		if i == 0 {
			c.name = message.Token
			i++
		}
		chatRoom.send <- *message
	}
}

func (c *Client) SendMessage() {
	for {
		fmt.Println("wait send content:", c.name)
		m := <-c.send
		fmt.Println("to: ", c.name, ", from:", m.Token, ", content:", m.Content, ", chan size:", len(c.send))
		if err := c.conn.WriteJSON(m); err != nil {
			fmt.Println("occur err:", err)
			c.conn.Close()
			chatRoom.unregister <- c
			return
		}
	}
}
