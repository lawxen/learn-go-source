package main

import (
	_ "MyGin/docs"
	"encoding/json"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
)

// @Summary 首页的GET接口
// @Description 获取首页数据
// @Router /getIndex [get]
// @Param q query string false "关键词"
// @Success 200 {object} map[string]any
// @Failure 200 {object} map[string]any
func getIndex(g *gin.Context) {
	context := gin.H{"state": "success"}
	context["msg"] = "hello " + g.Query("q")
	g.JSON(http.StatusOK, context)
}

// @Summary 首页的POST接口
// @Description 获取首页数据
// @Router /postIndex [post]
// @Param data body string true "请求参数"
// @Success 200 {object} map[string]any
// @Failure 200 {object} map[string]any
func postIndex(g *gin.Context) {
	context := gin.H{"state": "success"}
	data, _ := g.GetRawData()
	var body map[string]interface{}
	json.Unmarshal(data, &body)
	context["msg"] = body
	g.JSON(http.StatusOK, context)
}

// @title MyGin文档
// @version 1.0
// @description 这是我的API文档
// @Schemes http https
// @host 127.0.0.1:8000
// @BasePath /api/v1
func main() {
	r := gin.Default()
	// 配置跨域访问
	config := cors.DefaultConfig()
	// 允许所有域名
	config.AllowAllOrigins = true
	// 允许执行的请求方法
	config.AllowMethods = []string{"GET", "POST"}
	// 允许执行的请求头
	config.AllowHeaders = []string{"tus-resumable", "upload-length",
		"upload-metadata", "cache-control",
		"x-requested-with", "*"}
	r.Use(cors.New(config))
	// 设置Gin运行模式，根据运行模式是否生成文档
	gin.SetMode(gin.DebugMode)
	if gin.Mode() == "debug" {
		r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	}
	// 定义路由
	v1 := r.Group("/api/v1")
	v1.GET("/getIndex", getIndex)
	v1.POST("/postIndex", postIndex)
	// 运行端口必须与@host 127.0.0.1:8000的端口相同
	r.Run(":8000")
}
