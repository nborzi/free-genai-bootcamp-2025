package handlers

import (
	"github.com/gin-gonic/gin"
	"lang-portal/internal/service"
	"net/http"
	"strconv"
)

func GetWords(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	if page < 1 {
		page = 1
	}

	words, total, err := service.GetWords(page)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"items": words,
		"pagination": gin.H{
			"current_page":   page,
			"total_pages":    (total + 99) / 100, // ceiling division
			"total_items":    total,
			"items_per_page": 100,
		},
	})
}

func GetWord(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	word, err := service.GetWord(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if word == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Word not found"})
		return
	}
	c.JSON(http.StatusOK, word)
}

func ReviewWord(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	var review struct {
		Correct bool `json:"correct"`
	}
	if err := c.BindJSON(&review); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	err = service.ReviewWord(id, review.Correct)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "success"})
}
