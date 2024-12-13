package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// Response structure for JSON output
type Response struct {
	Message string `json:"message"`
}

func handler(w http.ResponseWriter, r *http.Request) {
	// Extract query parameter from URL
	name := r.URL.Query().Get("name")
	if name == "" {
		name = "World" // Default if name is not provided
	}

	// Create the response message
	response := Response{
		Message: fmt.Sprintf("Hello, %s!", name),
	}

	// Set content type to application/json
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	// Send JSON response
	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		log.Fatalf("Error encoding response: %v", err)
	}
}

func main() {
	// Handle the "/greet" path with the handler function
	http.HandleFunc("/greet", handler)

	// Start the HTTP server
	log.Println("Server started at http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
