package service

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUser(c *gin.Context) {
	name := c.Param("name")
	c.JSON(http.StatusOK, gin.H{
		"username": name,
		"surename": "Duisen",
		"age":      20,
	})
}

func CreateUser(c *gin.Context) {
	name := c.Param("name")
	comment := c.Param("comment")
	message := name + ": " + comment
	c.JSON(http.StatusCreated, gin.H{
		"username": "Alex",
		"surename": "Duisen",
		"age":      20,
		"message":  message,
	})
}

func UpdateUser(c *gin.Context) {
	c.JSON(http.StatusAccepted, gin.H{
		"username": "Alex",
		"surename": "Duisen",
		"age":      20,
	})
}

func DeleteUser(c *gin.Context) {
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
