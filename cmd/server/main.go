package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Health check
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	})

	// Submit feedback
	r.POST("/feedback", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "feedback received",
		})
	})

	r.Run(":8080")
}
