package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"publishd.net/internal/database"
)

func main() {
	// Initialize database connection
	if err := database.Connect(); err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

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
		dbStatus := "connected"
		if err := database.DB.Ping(); err != nil {
			dbStatus = "disconnected"
		}
		
		c.JSON(http.StatusOK, gin.H{
			"status":   "healthy",
			"database": dbStatus,
		})
	})

	r.Run(":8080")
}