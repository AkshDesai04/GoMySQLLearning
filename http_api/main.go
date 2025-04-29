package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/akshd/GoMySQLLearning/core/sort"
	"github.com/akshd/GoMySQLLearning/db"
)

func main() {
	http.HandleFunc("/cities", handleCities)
	http.HandleFunc("/cities/", handleCityUpdate)
	log.Println("HTTP API server starting on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleCities(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	searchTerm := r.URL.Query().Get("search")
	sortField := r.URL.Query().Get("sort_field")
	sortDir := r.URL.Query().Get("sort_dir")

	dbConn, err := db.ConnectDB()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer dbConn.Close()

	cities, err := db.GetCities(dbConn, searchTerm)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Apply sorting if sort field is specified
	if sortField != "" {
		sorter := sort.NewCitySorter()
		cities = sorter.SortCities(cities, sortField, sortDir)
	}

	response := db.APIResponse{
		API:    "HTTP",
		Port:   8080,
		Query:  searchTerm,
		Cities: cities,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func handleCityUpdate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "PUT, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method != http.MethodPut {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Extract city ID from URL
	cityIDStr := r.URL.Path[len("/cities/"):]
	cityID, err := strconv.Atoi(cityIDStr)
	if err != nil {
		http.Error(w, "Invalid city ID", http.StatusBadRequest)
		return
	}

	// Parse request body
	var updateData map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&updateData); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Get field and value from request
	if len(updateData) != 1 {
		http.Error(w, "Only one field can be updated at a time", http.StatusBadRequest)
		return
	}

	var field string
	var value interface{}
	for k, v := range updateData {
		field = k
		value = v
		break
	}

	// Connect to database
	dbConn, err := db.ConnectDB()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer dbConn.Close()

	// Update city
	if err := db.UpdateCity(dbConn, cityID, field, value); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
