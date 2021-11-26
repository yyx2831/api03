package main

import (
	router2 "api03/router"
	"github.com/gin-gonic/gin"
)

func main() {
	//gin
	router := gin.Default()

	ApiGroup := router.Group("/v1")
	router2.InitUserRouter(ApiGroup)
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
}
