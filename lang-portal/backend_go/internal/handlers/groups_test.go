package handlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestGetGroups(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.GET("/api/groups", GetGroups)

	t.Run("List Groups", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/api/groups", nil)
		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)

		var response map[string]interface{}
		err := json.Unmarshal(w.Body.Bytes(), &response)
		assert.NoError(t, err)
		
		groups, ok := response["groups"].([]interface{})
		assert.True(t, ok)
		assert.NotNil(t, groups)
	})
}

func TestGetGroup(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.GET("/api/groups/:id", GetGroup)

	tests := []struct {
		name     string
		groupID  string
		wantCode int
	}{
		{"Valid Group", "1", http.StatusOK},
		{"Invalid Group ID", "999", http.StatusNotFound},
		{"Invalid ID Format", "abc", http.StatusBadRequest},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			req, _ := http.NewRequest("GET", "/api/groups/"+tt.groupID, nil)
			r.ServeHTTP(w, req)

			assert.Equal(t, tt.wantCode, w.Code)

			if tt.wantCode == http.StatusOK {
				var response map[string]interface{}
				err := json.Unmarshal(w.Body.Bytes(), &response)
				assert.NoError(t, err)
				assert.Contains(t, response, "id")
				assert.Contains(t, response, "name")
				assert.Contains(t, response, "description")
			}
		})
	}
}

func TestGetGroupWords(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.GET("/api/groups/:id/words", GetGroupWords)

	tests := []struct {
		name     string
		groupID  string
		wantCode int
	}{
		{"Valid Group Words", "1", http.StatusOK},
		{"Invalid Group ID", "999", http.StatusNotFound},
		{"Invalid ID Format", "abc", http.StatusBadRequest},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			req, _ := http.NewRequest("GET", "/api/groups/"+tt.groupID+"/words", nil)
			r.ServeHTTP(w, req)

			assert.Equal(t, tt.wantCode, w.Code)

			if tt.wantCode == http.StatusOK {
				var response map[string]interface{}
				err := json.Unmarshal(w.Body.Bytes(), &response)
				assert.NoError(t, err)
				
				words, ok := response["words"].([]interface{})
				assert.True(t, ok)
				assert.NotNil(t, words)
			}
		})
	}
}
