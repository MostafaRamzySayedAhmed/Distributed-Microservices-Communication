package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"encoding/xml"
)

type GreetingRequest struct {
	XMLName xml.Name `xml:"soapenv:Envelope"`
	Xmlns   string   `xml:"xmlns:soapenv,attr"`
	Body    struct {
		Name string `xml:"name"`
	} `xml:"soapenv:Body"`
}

type GreetingResponse struct {
	XMLName xml.Name `xml:"soapenv:Envelope"`
	Xmlns   string   `xml:"xmlns:soapenv,attr"`
	Body    struct {
		Message string `xml:"message"`
	} `xml:"soapenv:Body"`
}

func main() {
	// SOAP request body
	reqBody := GreetingRequest{
		Xmlns: "http://schemas.xmlsoap.org/soap/envelope/",
	}
	reqBody.Body.Name = "GoUser"

	// Marshal request body into XML
	reqXML, err := xml.Marshal(reqBody)
	if err != nil {
		log.Fatalf("Error marshaling SOAP request: %v", err)
	}

	// Send SOAP request to the server
	url := "http://localhost:8080/greet"
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(reqXML))
	if err != nil {
		log.Fatalf("Error creating HTTP request: %v", err)
	}

	// Set the SOAPAction header (some servers may require it)
	req.Header.Set("Content-Type", "text/xml")
	req.Header.Set("SOAPAction", "http://localhost/greet")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Error sending SOAP request: %v", err)
	}
	defer resp.Body.Close()

	// Read the SOAP response
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error reading SOAP response: %v", err)
	}

	// Unmarshal the SOAP response
	var soapResp GreetingResponse
	err = xml.Unmarshal(body, &soapResp)
	if err != nil {
		log.Fatalf("Error unmarshaling SOAP response: %v", err)
	}

	// Print the response message
	fmt.Printf("Server Response: %s\n", soapResp.Body.Message)
}
