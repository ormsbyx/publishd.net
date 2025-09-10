package handlers

import (
	"html/template"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"publishd.net/internal/database"
	"publishd.net/internal/models"
)

func RenderHome(c *gin.Context) {
	// Get latest 6 stories for homepage
	stories, err := models.GetAllStories(database.DB)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "home.html", gin.H{
			"Title":  "Error",
			"Error":  "Failed to load stories",
			"Stories": []models.Story{},
		})
		return
	}

	// Limit to 6 stories for homepage
	if len(stories) > 6 {
		stories = stories[:6]
	}

	c.HTML(http.StatusOK, "home.html", gin.H{
		"Title":   "Home",
		"Stories": stories,
	})
}

func RenderStoriesList(c *gin.Context) {
	stories, err := models.GetAllStories(database.DB)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "stories.html", gin.H{
			"Title":   "Stories",
			"Error":   "Failed to load stories",
			"Stories": []models.Story{},
		})
		return
	}

	c.HTML(http.StatusOK, "stories.html", gin.H{
		"Title":   "Stories",
		"Stories": stories,
	})
}

func RenderStory(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.HTML(http.StatusBadRequest, "story.html", gin.H{
			"Title": "Invalid Story",
			"Error": "Invalid story ID",
		})
		return
	}

	story, err := models.GetStoryByID(database.DB, id)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "story.html", gin.H{
			"Title": "Error",
			"Error": "Failed to load story",
		})
		return
	}

	if story == nil {
		c.HTML(http.StatusNotFound, "story.html", gin.H{
			"Title": "Story Not Found",
			"Error": "Story not found",
		})
		return
	}

	// Calculate reading time (rough estimate: 200 words per minute)
	wordCount := len(strings.Fields(story.Content))
	readingTime := (wordCount / 200) + 1
	if readingTime < 1 {
		readingTime = 1
	}

	// Determine if this is paid content and if user should see preview
	showPreview := story.Price > 0 // For now, always show preview for paid content
	var preview string
	var formattedContent template.HTML

	if showPreview {
		// Show first paragraph as preview
		paragraphs := strings.Split(story.Content, "\n\n")
		if len(paragraphs) > 0 {
			preview = paragraphs[0]
			if len(paragraphs) > 1 {
				preview += "..."
			}
		}
	} else {
		// Format content with basic HTML (convert newlines to <br>)
		formatted := strings.ReplaceAll(story.Content, "\n\n", "</p><p>")
		formatted = "<p>" + formatted + "</p>"
		formattedContent = template.HTML(formatted)
	}

	c.HTML(http.StatusOK, "story.html", gin.H{
		"Title":            story.Title,
		"Story":            story,
		"ReadingTime":      readingTime,
		"ShowPreview":      showPreview,
		"Preview":          preview,
		"FormattedContent": formattedContent,
		// TODO: Add previous/next story navigation
	})
}

func RenderAdmin(c *gin.Context) {
	c.HTML(http.StatusOK, "admin.html", gin.H{
		"Title": "Admin - Story Management",
	})
}