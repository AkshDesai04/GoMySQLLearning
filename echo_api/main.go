package main

import (
	"net/http"

	"github.com/akshd/GoMySQLLearning/db"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Response().Header().Set("Access-Control-Allow-Origin", "*")
			c.Response().Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
			c.Response().Header().Set("Access-Control-Allow-Headers", "Content-Type")
			if c.Request().Method == http.MethodOptions {
				return c.NoContent(http.StatusOK)
			}
			return next(c)
		}
	})

	e.GET("/cities", handleCities)

	e.Logger.Fatal(e.Start(":8081"))
}

func handleCities(c echo.Context) error {
	dbConn, err := db.ConnectDB()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	defer dbConn.Close()

	searchTerm := c.QueryParam("search")
	cities, err := db.GetCities(dbConn, searchTerm)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	response := db.APIResponse{
		API:    "Echo",
		Port:   8081,
		Query:  searchTerm,
		Cities: cities,
	}

	return c.JSON(http.StatusOK, response)
}
