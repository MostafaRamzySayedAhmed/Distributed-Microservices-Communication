package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// Response structure for JSON response
type Response struct {
	Message string `json:"message"`
}

func main() {
	// URL of the server (assuming it's running on localhost:8080)
	serverURL := "http://localhost:8080/greet?name=GoUser"

	// Send a GET request to the server
	resp, err := http.Get(serverURL)
	if err != nil {
		log.Fatalf("Error sending request to server: %v", err)
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error reading response body: %v", err)
	}

	// Decode the JSON response
	var response Response
	err = json.Unmarshal(body, &response)
	if err != nil {
		log.Fatalf("Error decoding response: %v", err)
	}

	// Print the response
	fmt.Println("Server Response:", response.Message)
}
