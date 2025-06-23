package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 定义路由处理函数
func handles(c *gin.Context) {
	data := map[string]interface{}{
		"name":  "Golang",
		"value": "操作成功",
	}
	c.JSON(http.StatusOK, data)
}

func main() {
	// 创建Gin对象
	r := gin.Default()
	// 设置路由的版本信息
	apiv1 := r.Group("/api/v1/")
	// 路由分组user
	user := apiv1.Group("user/")
	{
		user.POST("register/", handles)
		user.POST("login/", handles)
		home := user.Group("home/")
		// 在分组user里面再分组home
		{
			home.GET("query/", handles)
			home.POST("modify/", handles)
		}
	}
	// 路由分组product
	product := apiv1.Group("product/")
	{
		product.GET("query/", handles)
		product.POST("modify/", handles)
	}
	// 监听并在 0.0.0.0:8000 上启动服务
	r.Run(":8000")
}
