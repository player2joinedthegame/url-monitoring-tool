package scheduler

import (
	"sync"
	"time"
)

// Scheduler represents a URL monitoring scheduler.
type Scheduler struct {
	jobs    []*Job
	stop    chan struct{}
	wait    *sync.WaitGroup
	running bool
}

// NewScheduler creates a new scheduler instance.
func NewScheduler() *Scheduler {
	return &Scheduler{
		jobs:    make([]*Job, 0),
		stop:    make(chan struct{}),
		wait:    &sync.WaitGroup{},
		running: false,
	}
}

// ScheduleJob schedules a job for URL monitoring.
func (s *Scheduler) ScheduleJob(job *Job) {
	s.jobs = append(s.jobs, job)
}

// Start starts the scheduler and begins monitoring URLs.
func (s *Scheduler) Start() {
	if s.running {
		return
	}

	s.running = true

	for _, job := range s.jobs {
		s.wait.Add(1)
		go func(job *Job) {
			defer s.wait.Done()
			job.Run(s.stop)
		}(job)
	}
}

// Stop stops the scheduler and waits for all jobs to finish.
func (s *Scheduler) Stop() {
	if !s.running {
		return
	}

	close(s.stop)
	s.wait.Wait()
	s.running = false
}