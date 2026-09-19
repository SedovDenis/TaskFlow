package main

import "errors"

type JobStore struct {
	jobs []Job
}

func (s *JobStore) Add(job Job) {
	s.jobs = append(s.jobs, job)
}

func (s *JobStore) GetByID(id int) (Job, error) {
	for _, job := range s.jobs {
		if job.ID == id {
			return job, nil
		}
	}

	return Job{}, errors.New("job with this id not found")
}

func (s *JobStore) List() []Job {
	cpSlice := make([]Job, len(s.jobs))
	copy(cpSlice, s.jobs)

	return cpSlice
}
