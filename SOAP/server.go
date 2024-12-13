package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/nathany/soap"
)

// Define the request and response structures
type GreetingRequest struct {
	Name string `xml:"name"`
}

type GreetingResponse struct {
	Message string `xml:"message"`
}

// SOAP service that generates a greeting message
func greet(w http.ResponseWriter, r *http.Request) {
	// Parse the SOAP request
	var req GreetingRequest
	err := soap.ReadRequest(r.Body, &req)
	if err != nil {
		http.Error(w, "Failed to read SOAP request", http.StatusInternalServerError)
		return
	}

	// Generate a response
	response := GreetingResponse{
		Message: fmt.Sprintf("Hello, %s!", req.Name),
	}

	// Write the SOAP response
	soap.WriteResponse(w, &response)
}

func main() {
	// Create a new SOAP server with a greet method
	http.HandleFunc("/greet", greet)

	// Start the HTTP server
	log.Println("SOAP server is running on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
