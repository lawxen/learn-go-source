package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 定义路由处理函数
func handles(c *gin.Context) {
	// 获取当前请求类型
	method := c.Request.Method
	fmt.Printf("当前请求类型为：%v\n", method)
	// 获取当前请求地址
	urls := c.Request.URL
	fmt.Printf("获取当前请求地址为：%v\n", urls)
	// 获取当前所有Cookies
	cookies := c.Request.Cookies()
	fmt.Printf("获取当前所有Cookies为：%v\n", cookies)
	// 获取Cookies的某个键值对
	cookie, _ := c.Cookie("aa")
	fmt.Printf("获取Cookies的某个键值对：%v\n", cookie)
	// 获取所有请求头信息
	header := c.Request.Header
	fmt.Printf("获取所有请求头信息：%v\n", header)
	// 获取请求头信息
	accept := c.GetHeader("Accept")
	fmt.Printf("获取请求头Accept信息：%v\n", accept)
	// 往请求信息添加数据
	c.Set("address", "GuangZhou")
	// 获取请求信息的新增数据
	address, _ := c.Get("address")
	fmt.Printf("获取请求信息的新增数据：%v\n", address)
	// 判断请求方式
	if method == "GET" {
		// 获取GET请求的请求参数
		// 如果请求参数age不存在，则设置默认值为66
		age := c.DefaultQuery("age", "66")
		fmt.Printf("请求参数name的值为：%v\n", age)
		// 如果请求参数name不存在，则返回空字符串
		name1 := c.Query("name")
		fmt.Printf("请求参数name1的值为：%v\n", name1)
		// 同一请求参数存在2个或以上，使用QueryArray()以切片格式表示
		name2 := c.QueryArray("name")
		fmt.Printf("请求参数name2的值为：%v\n", name2)
	} else {
		// 获取POST请求的请求参数
		// 获取表单的请求参数name
		name := c.PostForm("name")
		fmt.Printf("获取表单的请求参数name：%v\n", name)
		// 如果表单的请求参数age不存在，则设置默认值为66
		age := c.DefaultPostForm("age", "66")
		fmt.Printf("获取表单的请求参数age：%v\n", age)
		// 获取JSON数据
		var body map[string]interface{}
		var body1 struct {
			Name string `json:"name"`
		}
		// 使用GetRawData()读取JSON数据
		data, _ := c.GetRawData()
		// 可以使用结构体或集合方式表示JSON数据
		json.Unmarshal(data, &body)
		name1 := body["name"]
		fmt.Printf("获取JSON的请求参数name1：%v\n", name1)
		// 使用BindJSON()读取JSON数据
		// 可以使用结构体或集合方式表示JSON数据
		c.BindJSON(&body1)
		name2 := body1.Name
		fmt.Printf("获取JSON的请求参数name2：%v\n", name2)
	}

	data := map[string]interface{}{
		"name":  "Golang",
		"value": "操作成功",
	}
	c.JSON(http.StatusOK, data)
}

func main() {
	// 创建Gin对象
	r := gin.Default()
	r.Any("/", handles)
	r.Run(":8000")
}
