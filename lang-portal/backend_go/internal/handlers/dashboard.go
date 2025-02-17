package handlers

import (
	"github.com/gin-gonic/gin"
	"lang-portal/internal/service"
	"net/http"
)

// GetDashboard returns all dashboard data
func GetDashboard(c *gin.Context) {
	// Get last study session
	lastSession, err := service.GetLastStudySession()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Get study progress
	progress, err := service.GetStudyProgress()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Get quick stats
	stats, err := service.GetQuickStats()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Combine all data
	dashboard := gin.H{
		"last_session": lastSession,
		"progress":     progress,
		"stats":        stats,
	}

	c.JSON(http.StatusOK, dashboard)
}

func GetLastStudySession(c *gin.Context) {
	session, err := service.GetLastStudySession()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if session == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "No study sessions found"})
		return
	}
	c.JSON(http.StatusOK, session)
}

func GetStudyProgress(c *gin.Context) {
	progress, err := service.GetStudyProgress()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, progress)
}

func GetQuickStats(c *gin.Context) {
	stats, err := service.GetQuickStats()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, stats)
}
