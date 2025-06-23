package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"path"
)

// 定义路由处理函数
func indexHandles(c *gin.Context) {
	// 将模版文件作为网页内容
	c.HTML(http.StatusOK, "index.tmpl", gin.H{})
}
func uploadHandles(c *gin.Context) {
	// 获取表单的字段file
	file, _ := c.FormFile("file")
	// 创建文件目录
	dir := "./static/upload/"
	os.MkdirAll(dir, 0666)
	// 构建文件路径
	dst := path.Join(dir, file.Filename)
	// 上传文件至指定目录
	c.SaveUploadedFile(file, dst)
	// 返回响应数据
	c.JSON(http.StatusOK, gin.H{
		"status": "success",
	})
}

func uploadsHandles(c *gin.Context) {
	// 以切片方式获取文件
	form, _ := c.MultipartForm()
	files := form.File["file"]
	// 创建文件目录
	dir := "./static/upload/"
	os.MkdirAll(dir, 0666)
	for _, file := range files {
		// 构建文件路径
		dst := path.Join(dir, file.Filename)
		// 上传文件至指定目录
		c.SaveUploadedFile(file, dst)
	}
	// 返回响应数据
	c.JSON(http.StatusOK, gin.H{
		"status": "success",
	})
}

func main() {
	// 创建Gin对象
	r := gin.Default()
	// 配置静态资源
	r.Static("/static", "./static")
	// 使用LoadHTMLGlob()或LoadHTMLFiles()设置HTML模版文件所在文件夹
	r.LoadHTMLGlob("templates/*")
	// 路由定义
	r.GET("/", indexHandles)
	r.POST("/upload/", uploadHandles)
	r.POST("/uploads/", uploadsHandles)
	r.Run(":8000")
}
