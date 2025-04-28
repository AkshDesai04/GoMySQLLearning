package db

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

// City represents the structure of a city record
type City struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	CountryCode string `json:"country_code"`
	District    string `json:"district"`
	Population  int    `json:"population"`
}

// APIResponse represents the response structure for API calls
type APIResponse struct {
	API    string `json:"api"`
	Port   int    `json:"port"`
	Query  string `json:"query"`
	Cities []City `json:"cities"`
}

// ConnectDB establishes a connection to the MySQL database
func ConnectDB() (*sql.DB, error) {
	if err := godotenv.Load(); err != nil {
		return nil, fmt.Errorf("error loading .env file: %v", err)
	}

	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	// Set default values if environment variables are not set
	if dbHost == "" {
		dbHost = "localhost"
	}
	if dbPort == "" {
		dbPort = "3306"
	}
	if dbName == "" {
		dbName = "world"
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPassword, dbHost, dbPort, dbName)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, fmt.Errorf("error connecting to database: %v", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("error pinging database: %v", err)
	}

	return db, nil
}

// GetCities returns cities from the database that match the search term
func GetCities(db *sql.DB, searchTerm string) ([]City, error) {
	query := "SELECT ID, Name, CountryCode, District, Population FROM city"
	var args []interface{}
	if searchTerm != "" {
		query += " WHERE Name LIKE ? OR CountryCode LIKE ? OR District LIKE ?"
		searchPattern := "%" + searchTerm + "%"
		args = append(args, searchPattern, searchPattern, searchPattern)
	}
	rows, err := db.Query(query, args...)
	if err != nil {
		return nil, fmt.Errorf("error querying cities: %v", err)
	}
	defer rows.Close()

	var cities []City
	for rows.Next() {
		var city City
		err := rows.Scan(&city.ID, &city.Name, &city.CountryCode, &city.District, &city.Population)
		if err != nil {
			return nil, fmt.Errorf("error scanning city: %v", err)
		}
		cities = append(cities, city)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating rows: %v", err)
	}

	return cities, nil
}
