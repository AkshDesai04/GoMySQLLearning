package main

import (
	"net/http"

	"github.com/akshd/GoMySQLLearning/core/sort"
	"github.com/akshd/GoMySQLLearning/db"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.CORS())

	e.GET("/cities", handleCities)
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
