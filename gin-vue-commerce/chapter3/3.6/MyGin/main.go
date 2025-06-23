package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 定义路由处理函数
func HtmlHandles(c *gin.Context) {
	// 将变量title写入模版文件index.tmpl
	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"name": "Golang",
	})
}

func JsonHandles(c *gin.Context) {
	data := gin.H{
		"name": "Golang",
	}
	c.JSON(http.StatusOK, data)
}
func JsonpHandles(c *gin.Context) {
	data := gin.H{
		"name": "Golang",
	}
	c.JSONP(http.StatusOK, data)
}

func YamlHandles(c *gin.Context) {
	data := gin.H{
		"name": "Golang",
	}
	c.YAML(http.StatusOK, data)
}

func XmlHandles(c *gin.Context) {
	data := gin.H{
		"name": "Golang",
	}
	c.XML(http.StatusOK, data)
}

func TomlHandles(c *gin.Context) {
	data := gin.H{
		"name": "Golang",
	}
	c.TOML(http.StatusOK, data)
}

func main() {
	// 创建Gin对象
	r := gin.Default()
	// 使用LoadHTMLGlob()或LoadHTMLFiles()设置HTML模版文件所在文件夹
	r.LoadHTMLGlob("templates/*")
	r.GET("/html/", HtmlHandles)
	r.GET("/json/", JsonHandles)
	r.GET("/jsonp/", JsonpHandles)
	r.GET("/yaml/", YamlHandles)
	r.GET("/xml/", XmlHandles)
	r.GET("/toml/", TomlHandles)
	// 重定向
	r.GET("", func(c *gin.Context) {
		data := map[string]interface{}{
			"name": "Golang",
		}
		c.JSON(http.StatusOK, data)
	})
	// HTTP重定向
	r.GET("/reset/", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/")
	})
	// 路由重定向
	r.GET("/reset2/", func(c *gin.Context) {
		// 在上下文中改变路由的路径
		c.Request.URL.Path = "/"
		// 把修改后的上下文写入路由
		r.HandleContext(c)
	})
	r.Run(":8000")
}
