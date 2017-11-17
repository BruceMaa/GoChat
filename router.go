package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/BruceMaa/GoChat/wechat/mp"
	"log"
)

// 初始化路由配置
func initRouterOutter() http.Handler {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Welcome to GoChat!"})
	})

	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	wechatMpServer(router)

	return router
}

// 设置微信公众号服务
func wechatMpServer(router *gin.Engine) {
	wechatMpConfig := &ChatConfig.WechatConfig.WechatMpConfig
	wechatMp := &mp.WechatMp{
		Configure: *wechatMpConfig,
	}
	accessToken, err := wechatMp.AccessToken()
	if err != nil {
		panic(err)
	}
	log.Printf("wechat mp accessToken: %+v\n", accessToken)

	router.GET("/wechat", func(c *gin.Context) {
		log.Printf("URL:%s\n", c.Request.URL)
		log.Println(c.Request.URL.Query())
		log.Println(c.Request.PostForm)
		c.String(http.StatusOK, wechatMp.AuthEchostr(wechatMpConfig.Token, c.Request))
	})

	router.POST("/wechat", func(c *gin.Context) {
		log.Println(c.Request.URL)
		log.Println(c.Request.URL.Query())
		log.Println(c.Request.PostForm)
	})
}

// 对内网路由初始化
func initRouterInner() http.Handler {
	router := gin.Default()

	apisRouter(router)

	return router
}

// 设置对内网API路由
func apisRouter(router *gin.Engine) {
	apis := router.Group("/apis")
	{
		apis.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"messages": "Welcome GoChat Server API",
			})
		})
	}
}
