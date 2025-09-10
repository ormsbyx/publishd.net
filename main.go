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

	// Load HTML templates - debug version
	log.Println("🔍 Attempting to load templates...")
	
	// Check what files exist
	if files, err := filepath.Glob("web/templates/*.html"); err == nil && len(files) > 0 {
		log.Printf("Found template files: %v", files)
		r.LoadHTMLGlob("web/templates/*.html")
		log.Println("✅ Templates loaded successfully")
	} else if files, err := filepath.Glob("./web/templates/*.html"); err == nil && len(files) > 0 {
		log.Printf("Found template files: %v", files)
		r.LoadHTMLGlob("./web/templates/*.html")
		log.Println("✅ Templates loaded successfully")
	} else {
		// If templates can't be found, create a simple inline template
		log.Println("⚠️ Template files not found, using inline templates")
		r.SetHTMLTemplate(template.Must(template.New("").Parse(`
{{define "base.html"}}
<!DOCTYPE html>
<html>
<head><title>Publishd - Debug Mode</title></head>
<body>
<h1>Publishd Platform</h1>
<p>Template system is working! (Debug mode)</p>
<a href="/health">Health Check</a>
</body>
</html>
{{end}}
		`)))
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