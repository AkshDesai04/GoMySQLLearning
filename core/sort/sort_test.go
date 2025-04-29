package sort

import (
	"testing"

	"github.com/akshd/GoMySQLLearning/db"
)

func TestCitySorter_SortCities(t *testing.T) {
	// Test data
	cities := []db.City{
		{ID: 3, Name: "Cairo", CountryCode: "EGY", District: "Cairo", Population: 6789479},
		{ID: 1, Name: "Tokyo", CountryCode: "JPN", District: "Tokyo-to", Population: 7980230},
		{ID: 2, Name: "London", CountryCode: "GBR", District: "England", Population: 7285000},
		{ID: 4, Name: "Paris", CountryCode: "FRA", District: "Île-de-France", Population: 2125246},
	}

	tests := []struct {
		name      string
		field     string
		direction string
		expected  []int // Expected IDs in order
	}{
		{
			name:      "Sort by ID ascending",
			field:     "id",
			direction: "asc",
			expected:  []int{1, 2, 3, 4},
		},
		{
			name:      "Sort by ID descending",
			field:     "id",
			direction: "desc",
			expected:  []int{4, 3, 2, 1},
		},
		{
			name:      "Sort by Name ascending",
			field:     "name",
			direction: "asc",
			expected:  []int{3, 2, 4, 1},
		},
		{
			name:      "Sort by Name descending",
			field:     "name",
			direction: "desc",
			expected:  []int{1, 4, 2, 3},
		},
		{
			name:      "Sort by Country Code ascending",
			field:     "country_code",
			direction: "asc",
			expected:  []int{3, 4, 2, 1},
		},
		{
			name:      "Sort by Country Code descending",
			field:     "country_code",
			direction: "desc",
			expected:  []int{1, 2, 4, 3},
		},
		{
			name:      "Sort by Population ascending",
			field:     "population",
			direction: "asc",
			expected:  []int{4, 3, 2, 1},
		},
		{
			name:      "Sort by Population descending",
			field:     "population",
			direction: "desc",
			expected:  []int{1, 2, 3, 4},
		},
	}

	sorter := NewCitySorter()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Make a copy of the test data
			testCities := make([]db.City, len(cities))
			copy(testCities, cities)

			// Sort the cities
			sorted := sorter.SortCities(testCities, tt.field, tt.direction)

			// Verify the order
			for i, city := range sorted {
				if city.ID != tt.expected[i] {
					t.Errorf("Expected city ID %d at position %d, got %d", tt.expected[i], i, city.ID)
				}
			}
		})
	}
}

func TestCitySorter_SortCitiesEmpty(t *testing.T) {
	sorter := NewCitySorter()

	// Test empty slice
	emptyCities := []db.City{}
	sorted := sorter.SortCities(emptyCities, "id", "asc")
	if len(sorted) != 0 {
		t.Errorf("Expected empty slice, got %d elements", len(sorted))
	}

	// Test single element
	singleCity := []db.City{{ID: 1, Name: "Test"}}
	sorted = sorter.SortCities(singleCity, "id", "asc")
	if len(sorted) != 1 || sorted[0].ID != 1 {
		t.Errorf("Expected single city with ID 1, got %v", sorted)
	}
}

func TestCitySorter_SortCitiesStability(t *testing.T) {
	sorter := NewCitySorter()

	// Test data with duplicate names
	cities := []db.City{
		{ID: 1, Name: "London", CountryCode: "GBR", District: "England", Population: 7285000},
		{ID: 2, Name: "London", CountryCode: "CAN", District: "Ontario", Population: 339917},
		{ID: 3, Name: "London", CountryCode: "USA", District: "Kentucky", Population: 7996},
	}

	// Sort by name (should maintain original order for equal names)
	sorted := sorter.SortCities(cities, "name", "asc")

	// Verify the order
	expectedIDs := []int{1, 2, 3}
	for i, city := range sorted {
		if city.ID != expectedIDs[i] {
			t.Errorf("Expected city ID %d at position %d, got %d", expectedIDs[i], i, city.ID)
		}
	}
}
