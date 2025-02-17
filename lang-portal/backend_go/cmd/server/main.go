package main

import (
	"github.com/gin-gonic/gin"
	"lang-portal/internal/handlers"
	"lang-portal/internal/service"
	"log"
)

func main() {
	// Initialize database
	err := service.InitDB("words.db")
	if err != nil {
		log.Fatal("Failed to initialize database:", err)
	}
	defer service.CloseDB()

	r := gin.Default()

	// Enable CORS
	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	})

	// API routes
	api := r.Group("/api")
	{
		// Dashboard routes
		api.GET("/dashboard/last_study_session", handlers.GetLastStudySession)
		api.GET("/dashboard/study_progress", handlers.GetStudyProgress)
		api.GET("/dashboard/quick-stats", handlers.GetQuickStats)

		// Study activities routes
		api.GET("/study_activities/:id", handlers.GetStudyActivity)
		api.GET("/study_activities/:id/study_sessions", handlers.GetStudyActivitySessions)
		api.POST("/study_activities", handlers.CreateStudyActivity)

		// Words routes
		api.GET("/words", handlers.GetWords)
		api.GET("/words/:id", handlers.GetWord)
	}

	log.Println("Server starting on http://localhost:8081")
	log.Fatal(r.Run(":8081"))
}
