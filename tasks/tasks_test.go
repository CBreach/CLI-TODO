package tasks

import (
	"testing"
	"time"
)

func TestNewTask(t *testing.T) {
	now := time.Now()
	task := newTask(1, "this is a test", false, "tommorrow")

	if task.Name != "this is a test" {
		t.Errorf("expected Name %q, got %q", "this is a test", task.Name)
	}

	if task.ID != 1 {
		t.Errorf("expected ID %d, got %d", 1, task.ID)
	}

	if task.dueDate != "tommorrow" {
		t.Errorf("expected dueDate %q, got %q", "tommorrow", task.dueDate)
	}

	// Instead of checking exact time equality, check if it's within a short window
	if task.created.Before(now) || task.created.After(now.Add(time.Second)) {
		t.Errorf("expected created time to be ~now (%v), got %v", now, task.created)
	}
}
