package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/noazaj/btc-pipeline/requests"
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

	err = os.Mkdir("data", 0755)
	if err != nil {
		log.Printf("Error creating dir: %v", err)
	}

	_, err = os.Create("x-data.json")
	if err != nil {
		log.Printf("Error creating JSON data from source X: %v", err)
	}

	dataX, err := requests.XRequestAPI(bearerToken, os.Getenv("X_QUERY"), os.Getenv("X_URL"))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(dataX)

}
