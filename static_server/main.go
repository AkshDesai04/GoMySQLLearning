package main

import (
	"log"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("../static"))
	http.Handle("/", fs)

	log.Println("Static file server starting on :8083")
	log.Fatal(http.ListenAndServe(":8083", nil))
} 