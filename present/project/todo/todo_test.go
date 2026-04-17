package todo

import (
	"errors"
	"testing"
)

func TestLevel1AddAndList(t *testing.T) {
	list := New()

	first, err := list.Add("Read chapter 1")
	if err != nil {
		t.Fatalf("Add returned unexpected error: %v", err)
	}
	second, err := list.Add("Write notes")
	if err != nil {
		t.Fatalf("Add returned unexpected error: %v", err)
	}

	if first.ID != 1 || second.ID != 2 {
		t.Fatalf("IDs should be incremental, got %d and %d", first.ID, second.ID)
	}

	items := list.Items()
	if len(items) != 2 {
		t.Fatalf("Items length = %d, want 2", len(items))
	}
	if items[0].Title != "Read chapter 1" {
		t.Fatalf("unexpected first title: %q", items[0].Title)
	}
}

func TestLevel1AddRejectsEmptyTitle(t *testing.T) {
	list := New()
	_, err := list.Add("   ")
	if !errors.Is(err, ErrEmptyTitle) {
		t.Fatalf("expected ErrEmptyTitle, got %v", err)
	}
}

func TestLevel2DoneAndPendingCount(t *testing.T) {
	list := New()
	a, _ := list.Add("Task A")
	_, _ = list.Add("Task B")

	if got := list.PendingCount(); got != 2 {
		t.Fatalf("PendingCount() = %d, want 2", got)
	}

	if err := list.Done(a.ID); err != nil {
		t.Fatalf("Done returned unexpected error: %v", err)
	}

	if got := list.PendingCount(); got != 1 {
		t.Fatalf("PendingCount() = %d, want 1", got)
	}
}

func TestLevel3Remove(t *testing.T) {
	list := New()
	a, _ := list.Add("Task A")
	b, _ := list.Add("Task B")

	if err := list.Remove(a.ID); err != nil {
		t.Fatalf("Remove returned unexpected error: %v", err)
	}

	items := list.Items()
	if len(items) != 1 {
		t.Fatalf("Items length = %d, want 1", len(items))
	}
	if items[0].ID != b.ID {
		t.Fatalf("remaining item id = %d, want %d", items[0].ID, b.ID)
	}
}

func TestLevel4NotFoundErrors(t *testing.T) {
	list := New()

	if err := list.Done(999); !errors.Is(err, ErrNotFound) {
		t.Fatalf("Done should return ErrNotFound, got %v", err)
	}
	if err := list.Remove(999); !errors.Is(err, ErrNotFound) {
		t.Fatalf("Remove should return ErrNotFound, got %v", err)
	}
}
