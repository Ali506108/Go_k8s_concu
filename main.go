package main

import (
	"log"

	"github.com/Ali506108/Go_k8s_concu/model"
	"github.com/Ali506108/Go_k8s_concu/service"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.Use(gin.Logger(), gin.Recovery())
	router.GET("/user/:name", service.GetUser)
	router.POST("/user/:name/*comment", service.CreateUser)
	router.PUT("/user", service.UpdateUser)
	router.PATCH("/user", service.UpdateUser)
	router.DELETE("/user", service.DeleteUser)
	router.HEAD("/user", service.HeadUser)
	router.OPTIONS("/user", service.OptionUser)

	// new commit changes
	router.POST("/user/form", func(c *gin.Context) {
		var req model.FormRequest

		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(400, gin.H{
				"error": "Invalid JSON provided",
			})
			return
		}

		if req.Nick == "" {
			req.Nick = "Anonymous"
		}

		c.JSON(200, gin.H{
			"status":  "posted",
			"message": req.Message,
			"nick":    req.Nick,
		})
	})

	if err := router.Run("8432"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
