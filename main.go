package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RootHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello World!",
	})
}

func main() {
	router := gin.Default()

	router.GET("/", RootHandler)
	router.Run()
}