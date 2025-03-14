package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	runApp()
}

func runApp() {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	if err := r.Run("0.0.0.0:9009"); err != nil {
		panic(err)
	}
}
