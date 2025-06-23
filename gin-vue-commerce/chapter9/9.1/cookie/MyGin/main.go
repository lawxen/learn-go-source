package main

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	// 设置Session的存储方式
	store := cookie.NewStore([]byte("你这个老6"))
	// 自定义Session设置
	// 设置有效期一天
	store.Options(sessions.Options{Path: "/", MaxAge: 60 * 60 * 24})
	// 实例化Session并以中间件写入Gin
	r.Use(sessions.Sessions("token", store))
	// 设置Session数据
	r.GET("/", func(c *gin.Context) {
		s := sessions.Default(c)
		s.Set("name", "Tom")
		s.Save()
		c.JSON(200, gin.H{"name": "Tom"})
	})
	// 获取Session数据
	r.GET("/get/", func(c *gin.Context) {
		s := sessions.Default(c)
		res := gin.H{"token_name": s.Get("name")}
		// 设置生命周期，-1代表已过期，相当于清空Session数据
		s.Options(sessions.Options{Path: "/", MaxAge: -1})
		// 删除Session某个数据
		s.Delete("name")
		// 删除Session全部数据
		s.Clear()
		// 保存Session，否则删除失效
		s.Save()
		c.JSON(200, res)
	})
	r.Run(":8000")
}