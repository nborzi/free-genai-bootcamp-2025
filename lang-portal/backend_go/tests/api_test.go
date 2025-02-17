package tests

import (
	"encoding/json"
	"fmt"
	"lang-portal/internal/handlers"
	"lang-portal/internal/models"
	"lang-portal/internal/service"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type APITestSuite struct {
	suite.Suite
	router *gin.Engine
}

func (s *APITestSuite) SetupSuite() {
	// Set Gin to test mode
	gin.SetMode(gin.TestMode)

	// Initialize the database with test data
	err := service.InitDB("../words.db")
	if err != nil {
		s.T().Fatal("Failed to initialize database:", err)
	}

	// Setup router with all routes
	s.router = gin.New()
	api := s.router.Group("/api")
	{
		api.GET("/dashboard", handlers.GetDashboard)
		api.GET("/study-activities", handlers.GetStudyActivities)
		api.GET("/study-activities/:id", handlers.GetStudyActivity)
		api.POST("/study-activities/:id/start", handlers.StartStudyActivity)
		api.GET("/study-sessions", handlers.GetStudySessions)
		api.GET("/study-sessions/:id", handlers.GetStudySession)
		api.GET("/words", handlers.GetWords)
		api.GET("/words/:id", handlers.GetWord)
		api.POST("/words/:id/review", handlers.ReviewWord)
		api.GET("/groups", handlers.GetGroups)
		api.GET("/groups/:id", handlers.GetGroup)
		api.GET("/groups/:id/words", handlers.GetGroupWords)
	}
}

func (s *APITestSuite) TearDownSuite() {
	service.CloseDB()
}

// Test GET /api/dashboard
func (s *APITestSuite) TestGetDashboard() {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/dashboard", nil)
	s.router.ServeHTTP(w, req)

	assert.Equal(s.T(), http.StatusOK, w.Code)

	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(s.T(), err)

	// Verify response structure
	assert.Contains(s.T(), response, "last_session")
	assert.Contains(s.T(), response, "progress")
	assert.Contains(s.T(), response, "stats")
}

// Test GET /api/words with pagination
func (s *APITestSuite) TestGetWords() {
	testCases := []struct {
		name           string
		page           string
		expectedStatus int
		checkResponse  func(*testing.T, []byte)
	}{
		{
			name:           "Valid page 1",
			page:           "1",
			expectedStatus: http.StatusOK,
			checkResponse: func(t *testing.T, body []byte) {
				var response map[string]interface{}
				err := json.Unmarshal(body, &response)
				assert.NoError(t, err)
				assert.Contains(t, response, "items")
				assert.Contains(t, response, "pagination")
			},
		},
		{
			name:           "Invalid page -1",
			page:           "-1",
			expectedStatus: http.StatusOK, // Returns page 1
			checkResponse: func(t *testing.T, body []byte) {
				var response map[string]interface{}
				err := json.Unmarshal(body, &response)
				assert.NoError(t, err)
				pagination := response["pagination"].(map[string]interface{})
				assert.Equal(t, float64(1), pagination["current_page"])
			},
		},
	}

	for _, tc := range testCases {
		s.T().Run(tc.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			req, _ := http.NewRequest("GET", fmt.Sprintf("/api/words?page=%s", tc.page), nil)
			s.router.ServeHTTP(w, req)

			assert.Equal(t, tc.expectedStatus, w.Code)
			tc.checkResponse(t, w.Body.Bytes())
		})
	}
}

// Test GET /api/words/:id
func (s *APITestSuite) TestGetWord() {
	testCases := []struct {
		name           string
		wordID         string
		expectedStatus int
		checkResponse  func(*testing.T, []byte)
	}{
		{
			name:           "Existing word",
			wordID:         "1",
			expectedStatus: http.StatusOK,
			checkResponse: func(t *testing.T, body []byte) {
				var word service.WordWithStats
				err := json.Unmarshal(body, &word)
				assert.NoError(t, err)
				assert.NotEmpty(t, word.Spanish)
				assert.NotEmpty(t, word.English)
			},
		},
		{
			name:           "Non-existent word",
			wordID:         "999999",
			expectedStatus: http.StatusNotFound,
			checkResponse: func(t *testing.T, body []byte) {
				var response map[string]interface{}
				err := json.Unmarshal(body, &response)
				assert.NoError(t, err)
				assert.Contains(t, response, "error")
			},
		},
		{
			name:           "Invalid word ID",
			wordID:         "invalid",
			expectedStatus: http.StatusBadRequest,
			checkResponse: func(t *testing.T, body []byte) {
				var response map[string]interface{}
				err := json.Unmarshal(body, &response)
				assert.NoError(t, err)
				assert.Contains(t, response, "error")
			},
		},
	}

	for _, tc := range testCases {
		s.T().Run(tc.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			req, _ := http.NewRequest("GET", fmt.Sprintf("/api/words/%s", tc.wordID), nil)
			s.router.ServeHTTP(w, req)

			assert.Equal(t, tc.expectedStatus, w.Code)
			tc.checkResponse(t, w.Body.Bytes())
		})
	}
}

// Test POST /api/words/:id/review
func (s *APITestSuite) TestReviewWord() {
	testCases := []struct {
		name           string
		wordID         string
		body           string
		expectedStatus int
	}{
		{
			name:           "Valid review - correct",
			wordID:         "1",
			body:           `{"correct": true}`,
			expectedStatus: http.StatusOK,
		},
		{
			name:           "Valid review - incorrect",
			wordID:         "1",
			body:           `{"correct": false}`,
			expectedStatus: http.StatusOK,
		},
		{
			name:           "Invalid word ID",
			wordID:         "invalid",
			body:           `{"correct": true}`,
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:           "Invalid body",
			wordID:         "1",
			body:           `{"invalid": true}`,
			expectedStatus: http.StatusBadRequest,
		},
	}

	for _, tc := range testCases {
		s.T().Run(tc.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			req, _ := http.NewRequest("POST", fmt.Sprintf("/api/words/%s/review", tc.wordID), strings.NewReader(tc.body))
			req.Header.Set("Content-Type", "application/json")
			s.router.ServeHTTP(w, req)

			assert.Equal(t, tc.expectedStatus, w.Code)
		})
	}
}

func TestAPISuite(t *testing.T) {
	suite.Run(t, new(APITestSuite))
}
