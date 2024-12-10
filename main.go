package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	// Load the environment variables
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading enviornment variables: %v", err)
	}

	// Set API creds and ensure they were properly set
	apiKey := os.Getenv("X_API_KEY")
	apiSecret := os.Getenv("X_API_SECRET")
	bearerToken := os.Getenv("BEARER_TOKEN")
	if len(apiKey) == 0 || len(apiSecret) == 0 || len(bearerToken) == 0 {
		log.Fatal("Error setting API variables or Bearer token not initialized")
	}

	testRequest(bearerToken)

}

func testRequest(bearerToken string) {
	query := "?query=bitcoin&max_results=5"
	requestURL := fmt.Sprintf("https://api.x.com/2/tweets/search/all%s", query)

	req, err := http.NewRequest(http.MethodGet, requestURL, nil)
	if err != nil {
		log.Printf("Client could not create request: %v\n", err)
		os.Exit(1)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", bearerToken))

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Printf("Error making http request: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Got response code: %d\n\n", res.StatusCode)

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		log.Printf("Error reading response body: %v", err)
		os.Exit(1)
	}

	fmt.Printf("Raw response body: %v\n\nDecoded response body: %s", resBody, string(resBody))
}
