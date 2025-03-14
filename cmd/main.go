package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	runApp()
}

func runApp() {
	r := gin.Default()

	{
		v1 := r.Group("/v1")
		v1.GET("/:name", func(c *gin.Context) {

			c.JSON(http.StatusOK, gin.H{
				"message": fmt.Sprintf("Hello, %s!", c.Param("name")),
			})
		})
	}

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	if err := r.Run("0.0.0.0:9009"); err != nil {
		panic(err)
	}
}
