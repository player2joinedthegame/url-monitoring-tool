// main.go
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
	"time"

	"./config"
	"./http"
	"./notification"
	"./scheduler"
)

func main() {
	// Load configuration
	config := config.LoadConfig()

	// Read URLs from file
	urls, err := readURLsFromFile("urls/urls.txt")
	if err != nil {
		log.Fatalf("Failed to read URLs from file: %v", err)
	}

	// Create HTTP client
	httpClient := http.NewClient(10 * time.Second)

	// Create email notifier
	notifier := notification.NewEmailNotifier(config.SenderEmail, config.RecipientEmail, config.EmailService)

	// Create URL monitoring job
	job := scheduler.NewJob(urls, httpClient, notifier)

	// Create scheduler
	scheduler := scheduler.NewScheduler(config.MonitoringInterval)
	scheduler.RegisterJob(job)

	// Run scheduler
	scheduler.Run()
}

// readURLsFromFile reads URLs from a file and returns a slice of URLs.
func readURLsFromFile(filePath string) ([]string, error) {
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read URLs file: %v", err)
	}

	urls := strings.Split(string(content), "\n")
	var cleanedURLs []string
	for _, url := range urls {
		url = strings.TrimSpace(url)
		if url != "" {
			cleanedURLs = append(cleanedURLs, url)
		}
	}

	return cleanedURLs, nil
}