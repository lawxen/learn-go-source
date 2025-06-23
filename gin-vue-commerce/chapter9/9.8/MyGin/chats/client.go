package chats

import (
	"github.com/gorilla/websocket"
)

// 客户端对象
type Client struct {
	Chat *Chat
	Conn *websocket.Conn
}

// 读取数据通道
func (c *Client) GetData() {
	defer func() {
		c.Chat.Unregister <- c
		c.Conn.Close()
	}()
	for {
		_, message, err := c.Conn.ReadMessage()
		if err == nil {
			c.Chat.Broadcast <- message
		} else {
			break
		}
	}
}
