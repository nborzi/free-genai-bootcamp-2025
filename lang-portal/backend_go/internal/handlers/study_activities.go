package handlers

import (
	"github.com/gin-gonic/gin"
	"lang-portal/internal/service"
	"net/http"
	"strconv"
)

func GetStudyActivity(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid activity ID"})
		return
	}

	activity, err := service.GetStudyActivity(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if activity == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Activity not found"})
		return
	}
	c.JSON(http.StatusOK, activity)
}

func GetStudyActivitySessions(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid activity ID"})
		return
	}

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	if page < 1 {
		page = 1
	}

	sessions, total, err := service.GetStudyActivitySessions(id, page)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"items": sessions,
		"pagination": gin.H{
			"current_page":   page,
			"total_pages":    (total + 99) / 100, // ceiling division
			"total_items":    total,
			"items_per_page": 100,
		},
	})
}

func CreateStudyActivity(c *gin.Context) {
	var req struct {
		GroupID         int `json:"group_id" binding:"required"`
		StudyActivityID int `json:"study_activity_id" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	session, err := service.CreateStudySession(req.GroupID, req.StudyActivityID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, session)
}
