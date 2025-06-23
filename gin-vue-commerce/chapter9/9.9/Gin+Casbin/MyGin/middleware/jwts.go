package middleware

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"time"
)

var TokenExpireDuration = time.Minute * 30
var Secret = []byte("你这个老六")

type CustomClaims struct {
	// 自行添加字段
	Username string `json:"username"`
	// 内嵌JWT
	jwt.RegisteredClaims
}

// GenToken生成JWT
func GenToken(username string) (string, error) {
	expire := time.Now().Add(TokenExpireDuration)
	// 创建一个我们自己的声明
	claims := CustomClaims{
		username, // 自定义的用户名字段
		jwt.RegisteredClaims{
			Issuer: "奥力给", // 签发人
		},
	}
	// 使用指定的签名方法创建签名对象
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := t.SignedString(Secret)
	// 使用指定的secret签名并获得完整的编码后的字符串token
	// Token写入数据库
	j := Jwts{Token: token, Expire: expire}
	DB.Create(&j)
	return token, err
}

// ParseToken 解析JWT
func ParseToken(tokenString string) (*CustomClaims, error) {
	// 解析token
	// 如果是自定义Claim结构体则需要使用 ParseWithClaims 方法
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{},
		func(token *jwt.Token) (i interface{}, err error) {
			return Secret, nil
		})
	if err != nil {
		return nil, err
	}
	// 校验token
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("invalid token")
}

func JWTAuthMiddleware(c *gin.Context) {
	// 客户端携带Token有三种方式 1.放在请求头 2.放在请求体 3.放在URI
	// 将Token放在Header的Authorization中
	authHeader := c.Request.Header.Get("Authorization")
	if authHeader == "" {
		c.JSON(http.StatusOK, gin.H{
			"state": "fail",
			"msg":   "请求头的Authorization为空",
		})
		c.Abort()
		return
	}
	mc, err := ParseToken(authHeader)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"state": "fail",
			"msg":   "无效的Token",
		})
		c.Abort()
		return
	}
	var jwts Jwts
	DB.Where("token = ?", authHeader).First(&jwts)
	if jwts.Token != "" {
		if jwts.Expire.After(time.Now()) {
			jwts.Expire = time.Now().Add(TokenExpireDuration)
			DB.Save(&jwts)
		} else {
			// 强制删除表数据
			DB.Unscoped().Delete(&jwts)
		}
	} else {
		c.JSON(http.StatusOK, gin.H{
			"state": "fail",
			"msg":   "无效的Token",
		})
		c.Abort()
		return
	}
	// 将当前请求的username信息保存到请求的上下文c上
	// 路由处理函数通过c.Get("username")来获取用户信息
	c.Set("username", mc.Username)
	c.Next()
}
