package handlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestGetStudyActivities(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.GET("/api/study-activities", GetStudyActivities)

	t.Run("List Study Activities", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/api/study-activities", nil)
		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)

		var response map[string]interface{}
		err := json.Unmarshal(w.Body.Bytes(), &response)
		assert.NoError(t, err)
		
		activities, ok := response["activities"].([]interface{})
		assert.True(t, ok)
		assert.NotNil(t, activities)
	})
}

func TestGetStudyActivity(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.GET("/api/study-activities/:id", GetStudyActivity)

	tests := []struct {
		name       string
		activityID string
		wantCode   int
	}{
		{"Valid Activity", "1", http.StatusOK},
		{"Invalid Activity ID", "999", http.StatusNotFound},
		{"Invalid ID Format", "abc", http.StatusBadRequest},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			req, _ := http.NewRequest("GET", "/api/study-activities/"+tt.activityID, nil)
			r.ServeHTTP(w, req)

			assert.Equal(t, tt.wantCode, w.Code)

			if tt.wantCode == http.StatusOK {
				var response map[string]interface{}
				err := json.Unmarshal(w.Body.Bytes(), &response)
				assert.NoError(t, err)
				assert.Contains(t, response, "id")
				assert.Contains(t, response, "type")
				assert.Contains(t, response, "created_at")
			}
		})
	}
}

func TestStartStudyActivity(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.POST("/api/study-activities/:id/start", StartStudyActivity)

	tests := []struct {
		name       string
		activityID string
		wantCode   int
	}{
		{"Start Valid Activity", "1", http.StatusOK},
		{"Start Invalid Activity", "999", http.StatusNotFound},
		{"Invalid ID Format", "abc", http.StatusBadRequest},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			req, _ := http.NewRequest("POST", "/api/study-activities/"+tt.activityID+"/start", nil)
			r.ServeHTTP(w, req)

			assert.Equal(t, tt.wantCode, w.Code)
		})
	}
}

func TestGetStudySessions(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.GET("/api/study-sessions", GetStudySessions)

	t.Run("List Study Sessions", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/api/study-sessions", nil)
		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)

		var response map[string]interface{}
		err := json.Unmarshal(w.Body.Bytes(), &response)
		assert.NoError(t, err)
		
		sessions, ok := response["sessions"].([]interface{})
		assert.True(t, ok)
		assert.NotNil(t, sessions)
	})
}
