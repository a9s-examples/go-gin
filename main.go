package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
)

func main() {
	fmt.Println("Setting production mode")
	gin.SetMode(gin.ReleaseMode)

	port := os.Getenv("PORT")
	if len(port) < 1 {
		port = "8080"
	}

	router := gin.Default()
	router.GET("/", IndexHandler)
	router.Run(":" + port)
}

func IndexHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"status": "running",
	})
}
