package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/akshd/GoMySQLLearning/db"
	"github.com/gin-gonic/gin"
)

func TestHandleCities(t *testing.T) {
	// Setup
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	// Call the handler
	handleCities(c)

	// Check the status code
	if w.Code != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			w.Code, http.StatusOK)
	}

	// Check the response body
	var response db.APIResponse
	err := json.NewDecoder(w.Body).Decode(&response)
	if err != nil {
		t.Fatal(err)
	}

	if len(response.Cities) == 0 {
		t.Error("Expected non-empty cities array, got empty")
	}
}
