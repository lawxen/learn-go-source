package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	// 创建Gin对象
	r := gin.Default()
	r.Static("/static", "./static")
	// 由于已使用Static设置，无需再使用StaticFS重复设置，因此以代码注释表示
	//r.StaticFS("/static", http.Dir("static"))
	r.StaticFile("/favicon.ico", "./static/favicon.ico")
	// 监听并在 0.0.0.0:8000 上启动服务
	r.Run(":8000")
}
