package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	// 创建Gin对象
	r := gin.Default()
	// 定义路由和处理函数
	r.GET("/", func(c *gin.Context) {
		data := map[string]interface{}{
			"name":  "Golang",
			"value": "你好",
		}
		// 输出 : {"name":"Golang","value":"你好"}
		c.JSON(http.StatusOK, data)
	})
	// 监听并在 0.0.0.0:8000 上启动服务
	r.Run(":8000")
}
