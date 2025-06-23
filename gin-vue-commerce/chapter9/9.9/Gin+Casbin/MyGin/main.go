package main

import (
	"MyGin/middleware"
	"github.com/gin-gonic/gin"
)

func init() {
	middleware.Setup()
}

func main() {
	r := gin.Default()
	auth := r.Group("/auth/")
	auth.Use(middleware.JWTAuthMiddleware)
	auth.Use(middleware.CasbinMiddleware)
	// 获取JWT
	r.GET("/", func(c *gin.Context) {
		token, _ := middleware.GenToken("admin")
		c.JSON(200, gin.H{"token": token})
	})
	// 验证用户权限
	auth.Any("", func(c *gin.Context) {
		res := gin.H{"state": "success", "msg": "验证成功"}
		c.JSON(200, res)
	})
	r.Run(":8000")
}
