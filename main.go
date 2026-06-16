package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func getUser(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"username": "Alex",
		"surename": "Duisen",
		"age":      20,
	})
}

func createUser(c *gin.Context) {
	c.JSON(http.StatusCreated, gin.H{
		"username": "Alex",
		"surename": "Duisen",
		"age":      20,
	})
}

func updateUser(c *gin.Context) {
	c.JSON(http.StatusAccepted, gin.H{
		"username": "Alex",
		"surename": "Duisen",
		"age":      20,
	})
}

func deleteUser(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"username": "Alex",
		"surename": "Duisen",
		"age":      20,
	})
}

func HeadUser(c *gin.Context) {
	c.Status(http.StatusOK)
}

func OptionUser(c *gin.Context) {
	c.Status(http.StatusOK)
}

func main() {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.Use(gin.Logger(), gin.Recovery())
	router.GET("/user", getUser)
	router.POST("/user", createUser)
	router.PUT("/user", updateUser)
	router.PATCH("/user", updateUser)
	router.DELETE("/user", deleteUser)
	router.HEAD("/user", HeadUser)
	router.OPTIONS("/user", OptionUser)

	router.Run(":8432")
}
