package main

import (
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		appEnv := os.Getenv("APP_ENV")
		c.JSON(200, gin.H{
			"message": appEnv,
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}
