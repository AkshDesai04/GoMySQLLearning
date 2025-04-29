package sort

import (
	"strings"

	"github.com/akshd/GoMySQLLearning/db"
)

// CitySorter handles sorting operations for cities
type CitySorter struct{}

// NewCitySorter creates a new CitySorter instance
func NewCitySorter() *CitySorter {
	return &CitySorter{}
}

// SortCities sorts the cities slice using merge sort
func (s *CitySorter) SortCities(cities []db.City, field string, direction string) []db.City {
	if len(cities) <= 1 {
		return cities
	}

	// Create a copy of the slice to maintain stability
	citiesCopy := make([]db.City, len(cities))
	copy(citiesCopy, cities)

	mid := len(citiesCopy) / 2
	left := s.SortCities(citiesCopy[:mid], field, direction)
	right := s.SortCities(citiesCopy[mid:], field, direction)

	return s.merge(left, right, field, direction)
}

// merge combines two sorted slices into a single sorted slice
func (s *CitySorter) merge(left, right []db.City, field, direction string) []db.City {
	result := make([]db.City, 0, len(left)+len(right))
	i, j := 0, 0

	for i < len(left) && j < len(right) {
		// For equal elements, prefer the left one to maintain stability
		if s.compareCity(right[j], left[i], field, direction) {
			result = append(result, right[j])
			j++
		} else {
			result = append(result, left[i])
			i++
		}
	}

	result = append(result, left[i:]...)
	result = append(result, right[j:]...)
	return result
}

// compareCity compares two cities based on the specified field and sort direction
func (s *CitySorter) compareCity(a, b db.City, field, direction string) bool {
	isAsc := direction != "desc"

	switch field {
	case "id":
		if isAsc {
			return a.ID < b.ID
		}
		return a.ID > b.ID
	case "name":
		if isAsc {
			return strings.ToLower(a.Name) < strings.ToLower(b.Name)
		}
		return strings.ToLower(a.Name) > strings.ToLower(b.Name)
	case "country_code":
		if isAsc {
			return strings.ToLower(a.CountryCode) < strings.ToLower(b.CountryCode)
		}
		return strings.ToLower(a.CountryCode) > strings.ToLower(b.CountryCode)
	case "district":
		if isAsc {
			return strings.ToLower(a.District) < strings.ToLower(b.District)
		}
		return strings.ToLower(a.District) > strings.ToLower(b.District)
	case "population":
		if isAsc {
			return a.Population < b.Population
		}
		return a.Population > b.Population
	default:
		return true
	}
}
