package main

import "testing"

func TestGetByIDNormal(t *testing.T) {
	store := JobStore{}

	job := Job{
		ID:     1,
		Name:   "Test job",
		Status: JobStatusPending,
	}

	store.Add(job)

	foundJob, err := store.GetByID(1)

	if err != nil {
		t.Fatalf("Error get Job by id")
	}

	if foundJob.ID != job.ID {
		t.Fatalf("unexpected ID")
	}

	if foundJob.Name != job.Name {
		t.Fatalf("unexpected Name")
	}

	if foundJob.Status != job.Status {
		t.Fatalf("unexpected Status")
	}
}

func TestGetByIDInvalidID(t *testing.T) {
	store := JobStore{}

	job := Job{
		ID:     1,
		Name:   "Test job",
		Status: JobStatusPending,
	}

	store.Add(job)

	foundJob, err := store.GetByID(999)

	if err == nil {
		t.Fatalf("Must be error because 999 ID does not exist")
	}

	if foundJob != (Job{}) {
		t.Fatalf("Must be empty Job")
	}
}

func TestList(t *testing.T) {
	store := JobStore{}

	job1 := Job{
		ID:     1,
		Name:   "First job",
		Status: JobStatusPending,
	}

	job2 := Job{
		ID:     2,
		Name:   "Second job",
		Status: JobStatusPending,
	}

	store.Add(job1)
	store.Add(job2)

	jobs := store.List()

	if len(jobs) != 2 {
		t.Fatalf("Lenght jobs must be 2")
	}
}
