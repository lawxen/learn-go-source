package main

import (
	"MyGin/chats"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
)

// 实例化结构体Upgrader，设置结构体属性
// 解决跨域问题
var upGrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// 路由视图函数
func createWs(c *gin.Context, chat *chats.Chat) {
	// 创建WebSocket对象
	conn, err := upGrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		return
	}
	client := &chats.Client{Chat: chat, Conn: conn}
	go client.GetData()
	client.Chat.Register <- client
}

func main() {
	chat := chats.NewChat()
	go chat.Run()
	router := gin.Default()
	router.LoadHTMLFiles("index.html")
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})
	router.GET("/chat/", func(ctx *gin.Context) {
		createWs(ctx, chat)
	})
	router.Run(":8000")
}
