package db

import (
	"strings"
)

// MergeSort sorts the cities slice using the merge sort algorithm
func MergeSort(cities []City, field string, direction string) []City {
	if len(cities) <= 1 {
		return cities
	}

	mid := len(cities) / 2
	left := MergeSort(cities[:mid], field, direction)
	right := MergeSort(cities[mid:], field, direction)

	return merge(left, right, field, direction)
}

// merge combines two sorted slices into a single sorted slice
func merge(left, right []City, field, direction string) []City {
	result := make([]City, 0, len(left)+len(right))
	i, j := 0, 0

	for i < len(left) && j < len(right) {
		if compareCity(left[i], right[j], field, direction) {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}

	result = append(result, left[i:]...)
	result = append(result, right[j:]...)
	return result
}

// compareCity compares two cities based on the specified field and sort direction
func compareCity(a, b City, field, direction string) bool {
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
