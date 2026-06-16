package main

import (
	"github.com/gin-gonic/gin"
)

// learn gin
func main() {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.Use(gin.Logger(), gin.Recovery())
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	router.Run(":8432")
}
