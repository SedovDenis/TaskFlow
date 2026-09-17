package main

import (
	"errors"
	"time"
)

type JobStatus string

const (
	JobStatusPending   JobStatus = "pending"
	JobStatusRunning   JobStatus = "running"
	JobStatusCompleted JobStatus = "completed"
	JobStatusFailed    JobStatus = "failed"
)

type Job struct {
	ID        int
	Name      string
	Status    JobStatus
	CreatedAt time.Time
}

func NewJob(id int, name string) (Job, error) {
	if id <= 0 {
		return Job{}, errors.New("id must be greater than zero")
	}
	if name == "" {
		return Job{}, errors.New("name must not be empty")
	}

	job := Job{
		ID:        id,
		Name:      name,
		Status:    JobStatusPending,
		CreatedAt: time.Now(),
	}

	return job, nil
}
