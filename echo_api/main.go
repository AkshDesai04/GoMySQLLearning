package main

import (
	"net/http"
	"strconv"

	"github.com/akshd/GoMySQLLearning/core/sort"
	"github.com/akshd/GoMySQLLearning/db"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.CORS())

	e.GET("/cities", handleCities)
	e.PUT("/cities/:id", handleCityUpdate)
	e.Logger.Fatal(e.Start(":8081"))
}

func handleCities(c echo.Context) error {
	searchTerm := c.QueryParam("search")
	sortField := c.QueryParam("sort_field")
	sortDir := c.QueryParam("sort_dir")

	dbConn, err := db.ConnectDB()
	if err != nil {
		return err
	}
	defer dbConn.Close()

	cities, err := db.GetCities(dbConn, searchTerm)
	if err != nil {
		return err
	}

	// Apply sorting if sort field is specified
	if sortField != "" {
		sorter := sort.NewCitySorter()
		cities = sorter.SortCities(cities, sortField, sortDir)
	}

	response := db.APIResponse{
		API:    "Echo",
		Port:   8081,
		Query:  searchTerm,
		Cities: cities,
	}

	return c.JSON(http.StatusOK, response)
}

func handleCityUpdate(c echo.Context) error {
	// Extract city ID from URL
	cityID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid city ID")
	}

	// Parse request body
	var updateData map[string]interface{}
	if err := c.Bind(&updateData); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid request body")
	}

	// Get field and value from request
	if len(updateData) != 1 {
		return echo.NewHTTPError(http.StatusBadRequest, "Only one field can be updated at a time")
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
		return err
	}
	defer dbConn.Close()

	// Update city
	if err := db.UpdateCity(dbConn, cityID, field, value); err != nil {
		return err
	}

	return c.NoContent(http.StatusOK)
}
