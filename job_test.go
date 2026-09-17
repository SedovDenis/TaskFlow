package main

import "testing"

func TestNewJobSuccess(t *testing.T) {
	job, err := NewJob(1, "Test job")

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if job.ID != 1 {
		t.Errorf("unexpected ID: got %d, want 1", job.ID)
	}

	if job.Name != "Test job" {
		t.Errorf("unexpected name: got %q, want %q", job.Name, "Test job")
	}

	if job.Status != JobStatusPending {
		t.Errorf("unexpected status: got %q, want %q", job.Status, JobStatusPending)
	}

	if job.CreatedAt.IsZero() {
		t.Error("CreatedAt is not set")
	}
}

func TestIncorrectID(t *testing.T) {
	job, err := NewJob(0, "Test job")

	if err == nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if job != (Job{}) {
		t.Errorf("unexpected ID: got %d, want 1", job.ID)
	}
}

func TestIncorrectName(t *testing.T) {
	job, err := NewJob(1, "")

	if err == nil {
		t.Fatal("expected error, got nil")
	}

	if job != (Job{}) {
		t.Errorf("expected empty Job, got %+v", job)
	}
}
