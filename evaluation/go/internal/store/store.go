package store

import "sync"

type Task struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Done  bool   `json:"done"`
}

type Store struct {
	mu     sync.Mutex
	tasks  []Task
	nextID int
}

func New() *Store {
	return &Store{tasks: []Task{}, nextID: 1}
}

func (s *Store) List() []Task {
	s.mu.Lock()
	defer s.mu.Unlock()
	out := make([]Task, len(s.tasks))
	copy(out, s.tasks)
	return out
}

func (s *Store) Create(title string) Task {
	s.mu.Lock()
	defer s.mu.Unlock()
	t := Task{ID: s.nextID, Title: title, Done: false}
	s.nextID++
	s.tasks = append(s.tasks, t)
	return t
}

func (s *Store) Delete(id int) bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	for i, t := range s.tasks {
		if t.ID == id {
			s.tasks = append(s.tasks[:i], s.tasks[i+1:]...)
			return true
		}
	}
	return false
}

func (s *Store) GetByID(id int) (Task, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()
	for _, t := range s.tasks {
		if t.ID == id {
			return t, true
		}
	}
	return Task{}, false
}

func (s *Store) Update(id int, title *string, done *bool) (Task, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()
	for i, t := range s.tasks {
		if t.ID == id {
			if title != nil {
				s.tasks[i].Title = *title
			}
			if done != nil {
				s.tasks[i].Done = *done
			}
			return s.tasks[i], true
		}
	}
	return Task{}, false
}
