package store_test

import (
	"testing"

	"openspec-superpowers-guide/evaluation/go/internal/store"
)

func TestList_Empty(t *testing.T) {
	s := store.New()
	tasks := s.List()
	if len(tasks) != 0 {
		t.Fatalf("expected 0 tasks, got %d", len(tasks))
	}
}

func TestCreate(t *testing.T) {
	s := store.New()
	task := s.Create("buy milk")
	if task.ID != 1 {
		t.Errorf("expected ID 1, got %d", task.ID)
	}
	if task.Title != "buy milk" {
		t.Errorf("expected title 'buy milk', got %q", task.Title)
	}
	if task.Done {
		t.Error("expected Done to be false")
	}
}

func TestDelete_Exists(t *testing.T) {
	s := store.New()
	s.Create("task")
	if !s.Delete(1) {
		t.Error("expected Delete to return true")
	}
	if s.Delete(1) {
		t.Error("expected second Delete to return false")
	}
}

func TestDelete_NotFound(t *testing.T) {
	s := store.New()
	if s.Delete(99) {
		t.Error("expected Delete to return false for non-existent id")
	}
}
