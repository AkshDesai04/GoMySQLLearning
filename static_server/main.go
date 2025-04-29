package main

import (
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func main() {
	// Get the current working directory
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal("Error getting current working directory:", err)
	}

	// Construct the path to the static directory
	staticPath := filepath.Join(cwd, "static")

	// Log the path being used
	log.Printf("Current working directory: %s", cwd)
	log.Printf("Serving static files from: %s", staticPath)

	// Check if the directory exists
	if _, err := os.Stat(staticPath); os.IsNotExist(err) {
		log.Fatalf("Static directory not found at: %s", staticPath)
	}

	fs := http.FileServer(http.Dir(staticPath))
	http.Handle("/", fs)

	log.Println("Static file server starting on :8083")
	log.Fatal(http.ListenAndServe(":8083", nil))
}
