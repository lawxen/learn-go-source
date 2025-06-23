package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func CasbinMiddleware(c *gin.Context) {
	// 请求地址
	p := c.Request.URL.Path
	// 请求方法
	m := c.Request.Method
	// 获取用户名
	u, _ := c.Get("username")
	// 加载策略
	E.LoadPolicy()
	allowed, _ := E.Enforce(u, p, m)
	if allowed {
		c.Next()
	} else {
		c.JSON(http.StatusOK, gin.H{
			"state": "fail",
			"msg":   "暂无权限",
		})
		c.Abort()
		return
	}
}
