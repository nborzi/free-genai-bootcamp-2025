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
		api.GET("/dashboard", handlers.GetDashboard)
		
		// Study activities routes
		api.GET("/study-activities", handlers.GetStudyActivities)
		api.GET("/study-activities/:id", handlers.GetStudyActivity)
		api.POST("/study-activities/:id/start", handlers.StartStudyActivity)
		
		// Study sessions routes
		api.GET("/study-sessions", handlers.GetStudySessions)
		api.GET("/study-sessions/:id", handlers.GetStudySession)
		
		// Words routes
		api.GET("/words", handlers.GetWords)
		api.GET("/words/:id", handlers.GetWord)
		api.POST("/words/:id/review", handlers.ReviewWord)
		
		// Groups routes
		api.GET("/groups", handlers.GetGroups)
		api.GET("/groups/:id", handlers.GetGroup)
		api.GET("/groups/:id/words", handlers.GetGroupWords)
	}

	log.Println("Server starting on http://localhost:8081")
	log.Fatal(r.Run(":8081"))
}
