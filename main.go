package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"

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

	// Load HTML templates with proper nested template support
	log.Println("🔍 Loading nested templates...")
	
	// The KEY difference: Parse all templates together into one template set
	// This allows template inheritance to work properly
	templatePaths := []string{
		"web/templates/*.html",
		"./web/templates/*.html", 
		"templates/*.html",
	}
	
	var templatesLoaded bool
	for _, pattern := range templatePaths {
		if files, err := filepath.Glob(pattern); err == nil && len(files) > 0 {
			log.Printf("Found template files: %v", files)
			
			// This is the CORRECT way to handle nested templates in Gin:
			// Parse all templates together so they can reference each other
			tmpl := template.Must(template.ParseGlob(pattern))
			r.SetHTMLTemplate(tmpl)
			
			templatesLoaded = true
			log.Println("✅ Nested templates loaded successfully")
			break
		}
	}
	
	if !templatesLoaded {
		log.Fatal("❌ Failed to load HTML templates from any path")
	}
	
	// Serve static files with fallback paths
	staticPaths := []string{"./web/static", "web/static", "static"}
	for _, path := range staticPaths {
		if _, err := os.Stat(path); err == nil {
			r.Static("/static", path)
			log.Printf("✅ Serving static files from: %s", path)
			break
		}
	}

	// Web routes (HTML)
	r.GET("/", handlers.RenderHome)
	r.GET("/stories", handlers.RenderStoriesList)
	r.GET("/stories/:id", handlers.RenderStory)
	r.GET("/admin", handlers.RenderAdmin)
	r.GET("/admin/unpublished", handlers.RenderUnpublishedStories)

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

	// Get port from environment variable (Render sets this)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	
	log.Printf("Server starting on port %s", port)
	r.Run(":" + port)
}