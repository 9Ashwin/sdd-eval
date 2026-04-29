package store_test

import (
	"testing"

	"sdd-eval/evaluation/go/internal/store"
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

func TestGetByID_NotFound(t *testing.T) {
	s := store.New()
	_, ok := s.GetByID(99)
	if ok {
		t.Error("expected GetByID to return false for non-existent id")
	}
}

func TestGetByID_Exists(t *testing.T) {
	s := store.New()
	s.Create("buy milk")
	task, ok := s.GetByID(1)
	if !ok {
		t.Fatal("expected task to exist")
	}
	if task.ID != 1 {
		t.Errorf("expected ID 1, got %d", task.ID)
	}
	if task.Title != "buy milk" {
		t.Errorf("expected title 'buy milk', got %q", task.Title)
	}
}

func TestUpdate_NotFound(t *testing.T) {
	s := store.New()
	title := "new"
	_, ok := s.Update(99, &title, nil)
	if ok {
		t.Error("expected Update to return false for non-existent id")
	}
}

func TestUpdate_Title(t *testing.T) {
	s := store.New()
	s.Create("original")
	newTitle := "updated"
	task, ok := s.Update(1, &newTitle, nil)
	if !ok {
		t.Fatal("expected Update to succeed")
	}
	if task.Title != "updated" {
		t.Errorf("expected title 'updated', got %q", task.Title)
	}
	got, _ := s.GetByID(1)
	if got.Title != "updated" {
		t.Errorf("expected persisted title 'updated', got %q", got.Title)
	}
}

func TestUpdate_Done(t *testing.T) {
	s := store.New()
	s.Create("task")
	done := true
	task, ok := s.Update(1, nil, &done)
	if !ok {
		t.Fatal("expected Update to succeed")
	}
	if !task.Done {
		t.Error("expected Done to be true")
	}
	got, _ := s.GetByID(1)
	if !got.Done {
		t.Error("expected persisted Done to be true")
	}
}

func TestUpdate_Both(t *testing.T) {
	s := store.New()
	s.Create("original")
	newTitle := "both"
	done := true
	task, ok := s.Update(1, &newTitle, &done)
	if !ok {
		t.Fatal("expected Update to succeed")
	}
	if task.Title != "both" || !task.Done {
		t.Errorf("expected title='both' and done=true, got title=%q done=%v", task.Title, task.Done)
	}
}

func TestUpdate_NoFields(t *testing.T) {
	s := store.New()
	s.Create("task")
	task, ok := s.Update(1, nil, nil)
	if !ok {
		t.Fatal("expected Update to succeed (no-op)")
	}
	if task.Title != "task" {
		t.Errorf("expected title unchanged, got %q", task.Title)
	}
}
