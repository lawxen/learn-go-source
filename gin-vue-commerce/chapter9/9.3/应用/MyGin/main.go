package main

import (
	"MyGin/limits"
	"fmt"
	"github.com/gin-gonic/gin"
	"strings"
)

// 限流中间件
func rateLimiter() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取用户的IP地址
		ip := strings.Split(c.Request.RemoteAddr, ":")[0]
		fmt.Printf("用户的IP地址%v\n", ip)
		// 使用限流对象集合调用GetLimiter()获取限流对象
		limiter := limits.LimiterHandler.GetLimiter(ip, 2)
		// 没有可用令牌则禁止访问
		if !limiter.Allow() {
			c.JSON(403, gin.H{"status": "Forbidden"})
			c.Abort()
		}
		// 存在可用令牌则绩效执行
		c.Next()
	}
}

func main() {
	r := gin.Default()
	r.Use(rateLimiter())
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "Hello World"})
	})
	r.Run(":8000")
}
