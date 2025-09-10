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

	// Load HTML templates
	r.LoadHTMLGlob("web/templates/*")
	
	// Serve static files
	r.Static("/static", "./web/static")

	// Web routes (HTML)
	r.GET("/", handlers.RenderHome)
	r.GET("/stories", handlers.RenderStoriesList)
	r.GET("/stories/:id", handlers.RenderStory)
	r.GET("/admin", handlers.RenderAdmin)

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

	// JSON API endpoint for stories (keep for API access)
	r.GET("/api/stories", handlers.GetStories)
	r.GET("/api/stories/:id", handlers.GetStory)

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