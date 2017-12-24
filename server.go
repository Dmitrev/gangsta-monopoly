package main

import (
	"log"
	"net/http"
)

func startServer() {
	// Create simple file serverinfo info wa to serve static files
	fs := http.FileServer(http.Dir("public"))
	http.Handle("/", fs)

	log.Printf("StartingServer server on port 8000\n")
	err := http.ListenAndServe(":8000", nil)

	if err != nil {
		log.Fatalf("ListenAndServe: %v", err)
	}
}
