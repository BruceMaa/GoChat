package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 初始化路由配置
func initRouterOutter(middleware ...gin.HandlerFunc) http.Handler {
	router := gin.Default()
	router.Use(middleware...)

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Welcome to GoChat!"})
	})

	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	wechatMpServer(router)

	return router
}

// 对内网路由初始化
func initRouterInner(middleware ...gin.HandlerFunc) http.Handler {
	router := gin.Default()
	router.Use(middleware...)

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Welcome to GoChat!"})
	})

	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	apisRouter(router)

	return router
}

