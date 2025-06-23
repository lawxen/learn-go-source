package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 定义路由处理函数
func handles(c *gin.Context) {
	// 返回响应
	fmt.Printf("执行路由处理函数\n")
	c.JSON(http.StatusOK, gin.H{
		"status": "success",
	})
}

func MyMiddleware(c *gin.Context) {
	if c.Query("end") == "1" {
		fmt.Printf("中间件中断了，不再往下执行\n")
		c.JSON(http.StatusOK, gin.H{
			"status": "fail",
		})
		c.Abort()
	} else {
		fmt.Printf("中间件完成执行，继续往下执行\n")
		c.Next()
		fmt.Printf("中间件已完成执行\n")
	}
}

func main() {
	// 创建Gin对象
	r := gin.Default()
	// 路由定义
	// 所有路由使用中间件MyMiddleware
	//r.Use(MyMiddleware)
	// 路由分组user使用中间件MyMiddleware
	user := r.Group("/user/", MyMiddleware)
	//user.Use(MyMiddleware)
	user.GET("/", handles)
	// 单路由使用中间件MyMiddleware
	r.GET("/school/", MyMiddleware, handles)
	r.Run(":8000")
}
