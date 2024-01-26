package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	// Print a message indicating the purpose of the program
	fmt.Println("Tiny URL API consumption")

	// Check if a command-line argument is provided
	if len(os.Args) != 2 {
		// If no URL is provided, print an error message to standard error
		fmt.Fprintf(os.Stderr, "Error: Please provide a URL")
		// Exit the program with a non-zero status code
		os.Exit(1)
	}

	// Base URL for the TinyURL API
	baseURL := "http://tinyurl.com/api-create.php?url="

	// Extract the URL to be shortened from the command-line arguments
	urlToShorten := os.Args[1]

	// Construct the complete URL for the HTTP GET request
	getReqURL := baseURL + urlToShorten

	// Send an HTTP GET request to the TinyURL API
	response, err := http.Get(getReqURL)
	if err != nil {
		// If an error occurs during the HTTP request, log the error and exit the program
		log.Fatal(err)
	} else {
		// If the request is successful, defer closing the response body until the function returns
		defer response.Body.Close()
		// Copy the response body (shortened URL) to standard output
		_, err := io.Copy(os.Stdout, response.Body)
		if err != nil {
			// If an error occurs while copying the response body, log the error and exit the program
			log.Fatal(err)
		}
	}
}
