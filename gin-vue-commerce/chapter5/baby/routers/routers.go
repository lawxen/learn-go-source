package routers

import (
	"baby/middleware"
	v1 "baby/servers/v1"
	"baby/settings"
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitRouter() *gin.Engine {
	gin.SetMode(settings.Mode)
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.StaticFS("/static", http.Dir("static"))
	apiv1 := r.Group("/api/v1/")
	// 路由分组，部分路由设置中间件JWTAuthMiddleware验证用户
	commodity := apiv1.Group("")
	{
		// 网站首页
		commodity.GET("home/", v1.Home)
		// 商品列表
		commodity.GET("commodity/list/", v1.CommodityList)
		// 商品详细
		commodity.GET("commodity/detail/:id/", v1.CommodityDetail)
		// 用户注册登录
		commodity.POST("shopper/login/", v1.ShopperLogin)
	}
	//shopper := apiv1.Group("")
	//shopper.Use(middleware.JWTAuthMiddleware())
	shopper := apiv1.Group("", middleware.JWTAuthMiddleware)
	{
		// 商品收藏
		shopper.POST("commodity/collect/", v1.CommodityCollect)
		// 退出登录
		shopper.POST("shopper/logout/", v1.ShopperLogout)
		// 个人主页
		shopper.GET("shopper/home/", v1.ShopperHome)
		// 加入购物车
		shopper.POST("shopper/shopcart/", v1.ShopperShopCart)
		// 购物车列表
		shopper.GET("shopper/shopcart/", v1.ShopperShopCart)
		// 在线支付
		shopper.POST("shopper/pays/", v1.ShopperPays)
		// 删除购物车商品
		shopper.POST("shopper/delete/", v1.ShopperDelete)
	}
	return r
}
