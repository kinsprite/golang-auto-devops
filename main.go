package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	jsoniter "github.com/json-iterator/go"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

func main() {
	println("Hello 'golang-auto-devops'")
	mainRouter().Run(":5000") // 监听并在 0.0.0.0:5000 上启动服务
}

func mainRouter() *gin.Engine {
	engine := gin.New()

	engine.GET("/healthz", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "OK",
		})
	})

	engine.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	engine.GET("/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "world",
		})
	})

	engine.GET("/add", func(c *gin.Context) {
		a := 2
		b := 8
		result := add(a, b)

		c.JSON(http.StatusOK, gin.H{
			"result": result,
		})
	})

	return engine
}
