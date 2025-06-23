package settings

import (
	"github.com/gin-gonic/gin"
	"time"
)

type Database struct {
	User     string
	Password string
	Host     string
	Name     string
}

var MySQLSetting = &Database{
	User:     "root",
	Password: "1234",
	Host:     "127.0.0.1:3306",
	Name:     "baby",
}

// gin.DebugMode：表示开发环境模式
// gin.ReleaseMode：表示生产环境模式
// gin.TestMode：表示测试环境模式
var Mode = gin.ReleaseMode

// JWT的有效时间
var TokenExpireDuration = time.Minute * 30

// JWT的加密盐
var Secret = []byte("你这个老六")

// 分页功能，每一页的数据量
var PageSize = 6

// 支付信息
var AppId = "xxxx"
var AlipayPublicKeyString = `xxx`
var AppPrivateKeyString = `xxx`
