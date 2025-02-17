package handlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func setupTestRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	return r
}

func TestGetWords(t *testing.T) {
	r := setupTestRouter()
	r.GET("/api/words", GetWords)

	// Test case 1: Default page (page 1)
	t.Run("Default page", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/api/words", nil)
		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
		
		var response map[string]interface{}
		err := json.Unmarshal(w.Body.Bytes(), &response)
		assert.NoError(t, err)
		
		// Check response structure
		assert.Contains(t, response, "items")
		assert.Contains(t, response, "pagination")
	})

	// Test case 2: Specific page
	t.Run("Page 2", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/api/words?page=2", nil)
		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
		
		var response map[string]interface{}
		err := json.Unmarshal(w.Body.Bytes(), &response)
		assert.NoError(t, err)
		
		pagination := response["pagination"].(map[string]interface{})
		assert.Equal(t, float64(2), pagination["current_page"])
	})

	// Test case 3: Invalid page number
	t.Run("Invalid page", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/api/words?page=invalid", nil)
		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code) // Should default to page 1
		
		var response map[string]interface{}
		err := json.Unmarshal(w.Body.Bytes(), &response)
		assert.NoError(t, err)
		
		pagination := response["pagination"].(map[string]interface{})
		assert.Equal(t, float64(1), pagination["current_page"])
	})
}

func TestGetWord(t *testing.T) {
	r := setupTestRouter()
	r.GET("/api/words/:id", GetWord)

	// Test case 1: Existing word
	t.Run("Existing word", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/api/words/1", nil)
		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
		
		var response map[string]interface{}
		err := json.Unmarshal(w.Body.Bytes(), &response)
		assert.NoError(t, err)
		assert.Contains(t, response, "spanish")
		assert.Contains(t, response, "english")
	})

	// Test case 2: Non-existent word
	t.Run("Non-existent word", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/api/words/999999", nil)
		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusNotFound, w.Code)
	})

	// Test case 3: Invalid ID
	t.Run("Invalid ID", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/api/words/invalid", nil)
		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})
}
