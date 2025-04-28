package main

import (
	"net/http"

	"github.com/akshd/GoMySQLLearning/db"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		if c.Request.Method == http.MethodOptions {
			c.AbortWithStatus(200)
			return
		}
		c.Next()
	})

	r.GET("/cities", handleCities)

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
	cities, err := db.GetCities(dbConn, searchTerm)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	response := db.APIResponse{
		API:    "Gin",
		Port:   8082,
		Query:  searchTerm,
		Cities: cities,
	}

	c.JSON(http.StatusOK, response)
}
