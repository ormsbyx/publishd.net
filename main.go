package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"publishd.net/internal/database"
	"publishd.net/internal/handlers"
)

func main() {
	// Initialize database connection
	if err := database.Connect(); err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Run database migrations
	if err := database.RunMigrations(); err != nil {
		log.Fatal("Failed to run migrations:", err)
	}

	r := gin.Default()

	// Home page
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Welcome to Publishd - Premium Reading Platform",
			"version": "0.1.0",
		})
	})

	// Stories API routes
	api := r.Group("/api/v1")
	{
		api.GET("/stories", handlers.GetStories)
		api.GET("/stories/:id", handlers.GetStory)
		api.POST("/stories", handlers.CreateStory)
		api.PUT("/stories/:id", handlers.UpdateStory)
		api.DELETE("/stories/:id", handlers.DeleteStory)
		api.POST("/stories/:id/publish", handlers.PublishStory)
	}

	// Public stories routes (for reading interface)
	r.GET("/stories", handlers.GetStories)
	r.GET("/stories/:id", handlers.GetStory)

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