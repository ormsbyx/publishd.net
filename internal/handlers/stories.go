package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"publishd.net/internal/database"
	"publishd.net/internal/models"
)

func GetStories(c *gin.Context) {
	stories, err := models.GetAllStories(database.DB)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to fetch stories",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"stories": stories,
		"count":   len(stories),
	})
}

func GetStory(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid story ID",
		})
		return
	}

	story, err := models.GetStoryByID(database.DB, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to fetch story",
		})
		return
	}

	if story == nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Story not found",
		})
		return
	}

	c.JSON(http.StatusOK, story)
}

func CreateStory(c *gin.Context) {
	var storyData models.StoryCreate
	if err := c.ShouldBindJSON(&storyData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request data",
		})
		return
	}

	story, err := models.CreateStory(database.DB, &storyData)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to create story",
		})
		return
	}

	c.JSON(http.StatusCreated, story)
}

func UpdateStory(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid story ID",
		})
		return
	}

	var updates models.StoryUpdate
	if err := c.ShouldBindJSON(&updates); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request data",
		})
		return
	}

	story, err := models.UpdateStory(database.DB, id, &updates)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to update story",
		})
		return
	}

	if story == nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Story not found",
		})
		return
	}

	c.JSON(http.StatusOK, story)
}

func DeleteStory(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid story ID",
		})
		return
	}

	err = models.DeleteStory(database.DB, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to delete story",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Story deleted successfully",
	})
}

func PublishStory(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid story ID",
		})
		return
	}

	story, err := models.PublishStory(database.DB, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to publish story",
		})
		return
	}

	if story == nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Story not found or already published",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Story published successfully",
		"story":   story,
	})
}