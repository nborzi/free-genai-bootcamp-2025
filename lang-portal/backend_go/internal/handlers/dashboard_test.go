package handlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestGetDashboard(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.GET("/api/dashboard", GetDashboard)

	t.Run("Get Dashboard Data", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/api/dashboard", nil)
		r.ServeHTTP(w, req)

		// Check status code
		assert.Equal(t, http.StatusOK, w.Code)

		// Parse response
		var response map[string]interface{}
		err := json.Unmarshal(w.Body.Bytes(), &response)
		assert.NoError(t, err)

		// Check response structure
		assert.Contains(t, response, "last_session")
		assert.Contains(t, response, "progress")
		assert.Contains(t, response, "stats")

		// Check stats structure
		stats, ok := response["stats"].(map[string]interface{})
		assert.True(t, ok)
		assert.Contains(t, stats, "total_words")
		assert.Contains(t, stats, "words_learned")
		assert.Contains(t, stats, "learning_streak")
	})
}
