package main

import (
	"net/http"
	"strconv"

	"github.com/akshd/GoMySQLLearning/db"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, PUT, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	})

	r.GET("/cities", handleCities)
	r.PUT("/cities/:id", handleCityUpdate)

	r.Run(":8082")
}

func handleCities(c *gin.Context) {
	dbConn, err := db.ConnectDB()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer dbConn.Close()

	searchTerm := c.Query("search")
	sortField := c.Query("sort_field")
	sortDir := c.Query("sort_dir")

	cities, err := db.GetCities(dbConn, searchTerm)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Apply sorting if sort field is specified
	if sortField != "" {
		cities = db.MergeSort(cities, sortField, sortDir)
	}

	response := db.APIResponse{
		API:    "Gin",
		Port:   8082,
		Query:  searchTerm,
		Cities: cities,
	}

	c.JSON(http.StatusOK, response)
}

func handleCityUpdate(c *gin.Context) {
	// Extract city ID from URL
	cityID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid city ID"})
		return
	}

	// Parse request body
	var updateData map[string]interface{}
	if err := c.BindJSON(&updateData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	// Get field and value from request
	if len(updateData) != 1 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Only one field can be updated at a time"})
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
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer dbConn.Close()

	// Update city
	if err := db.UpdateCity(dbConn, cityID, field, value); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusOK)
}
