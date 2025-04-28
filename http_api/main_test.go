package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/akshd/GoMySQLLearning/db"
)

func TestHandleCities(t *testing.T) {
	// Create a request to pass to our handler
	req, err := http.NewRequest("GET", "/cities", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a ResponseRecorder to record the response
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handleCities)

	// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
	handler.ServeHTTP(rr, req)

	// Check the status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body
	var response db.APIResponse
	err = json.NewDecoder(rr.Body).Decode(&response)
	if err != nil {
		t.Fatal(err)
	}

	// Check if we got any cities
	if len(response.Cities) == 0 {
		t.Error("Expected non-empty cities array, got empty")
	}
}
