package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Home page
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Welcome to Publishd - Premium Reading Platform",
			"version": "0.1.0",
		})
	})

	// Stories routes
	r.GET("/stories", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"stories": []string{},
			"message": "Story listing coming soon",
		})
	})

	r.GET("/stories/:id", func(c *gin.Context) {
		id := c.Param("id")
		c.JSON(http.StatusOK, gin.H{
			"story_id": id,
			"message":  "Individual story view coming soon",
		})
	})

	// Health check
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "healthy",
		})
	})

	r.Run(":8080")
}