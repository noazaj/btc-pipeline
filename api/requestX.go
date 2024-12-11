package api

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

func XRequestAPI(bearerToken, query, url string) (string, error) {
	requestURL := fmt.Sprintf("%s%s", url, query)

	var (
		res *http.Response
		err error
	)

	maxRetries := 5
	for attempts := 0; attempts < maxRetries; attempts++ {
		req, err := http.NewRequest(http.MethodGet, requestURL, nil)
		if err != nil {
			log.Printf("Client could not create request: %v\n", err)
			os.Exit(1)
		}

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", bearerToken))

		res, err = http.DefaultClient.Do(req)
		if err != nil {
			return "", fmt.Errorf("error making HTTP request: %v\n", err)
		}

		if res != nil {
			defer res.Body.Close()
		}

		if res.StatusCode == 429 {
			retryAfter := res.Header.Get("X-Rate-Limit-Limit")
			if retryAfter != "" {
				waitTime, _ := strconv.Atoi(retryAfter)
				log.Printf("Rate limited. Retrying after %d seconds...\n", waitTime)
				time.Sleep(time.Duration(waitTime) * time.Second)
			} else {
				resetTime, _ := strconv.ParseInt(res.Header.Get("X-Rate-Limit-Reset"), 10, 64)
				currentTime := time.Now().Unix()
				waitTime := time.Duration(resetTime-currentTime) * time.Second

				if waitTime > 0 {
					log.Printf("Rate limited. Waiting until %v...\n", time.Unix(resetTime, 0))
					time.Sleep(waitTime)
				}
			}
			continue
		}
		break
	}

	if res == nil {
		return "", fmt.Errorf("failed to get a response after retries.")
	}

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		log.Printf("Error reading response body: %v", err)
		os.Exit(1)
	}
	defer res.Body.Close()

	return string(resBody), nil
}
