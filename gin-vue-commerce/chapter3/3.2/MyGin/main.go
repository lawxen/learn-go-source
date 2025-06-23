package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	// 创建Gin对象
	r := gin.Default()
	// 路由/user/:name设置路由变量name
	r.GET("/user/:name", func(c *gin.Context) {
		data := map[string]interface{}{
			"name":  "Golang1",
			"value": c.Param("name"),
		}
		// 输出 : {"name":"Golang","value":"你好"}
		c.JSON(http.StatusOK, data)
	})
	// 路由/product/*name设置路由变量name
	r.GET("/product/*name", func(c *gin.Context) {
		data := map[string]interface{}{
			"name":  "Golang2",
			"value": c.Param("name"),
		}
		// 输出 : {"name":"Golang","value":"你好"}
		c.JSON(http.StatusOK, data)
	})
	// 监听并在 0.0.0.0:8000 上启动服务
	r.Run(":8000")
}
