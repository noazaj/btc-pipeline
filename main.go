package main

import (
	"log"
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
	if len(apiKey) == 0 || len(apiSecret) == 0 {
		log.Fatal("Error setting API variables")
	}

}
