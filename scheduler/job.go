package scheduler

import (
	"time"
)

// Job represents a URL monitoring job.
type Job struct {
	URL        string
	Interval   time.Duration
	handler    func(string)
	lastRun    time.Time
	lastResult bool
}

// NewJob creates a new URL monitoring job.
func NewJob(url string, interval time.Duration, handler func(string)) *Job {
	return &Job{
		URL:      url,
		Interval: interval,
		handler:  handler,
	}
}

// Run runs the URL monitoring job at the specified interval until the stop signal is received.
func (j *Job) Run(stop <-chan struct{}) {
	ticker := time.NewTicker(j.Interval)
	defer ticker.Stop()

	for {
		select {
		case <-stop:
			return
		case <-ticker.C:
			// Perform URL check
			// Implement your logic to check the availability and response time of the URL
			// Set the result to true if the URL is up, and false otherwise
			result := true

			// If the result has changed since the last run, handle the URL check
			if result != j.lastResult {
				j.handler(j.URL)
			}

			j.lastResult = result
			j.lastRun = time.Now()
		}
	}
}