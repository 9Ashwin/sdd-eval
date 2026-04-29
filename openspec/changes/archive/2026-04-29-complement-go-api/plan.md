# complement-go-api Implementation Plan

> **For agentic workers:** REQUIRED SUB-SKILL: Use superpowers:subagent-driven-development (recommended) or superpowers:executing-plans to implement this plan task-by-task. Steps use checkbox (`- [ ]`) syntax for tracking.

**Goal:** 补全 Go Task Manager API 的 RESTful CRUD 操作 — 新增 GET /tasks/{id} 和 PATCH /tasks/{id} 端点。

**Architecture:** 三层结构 — Store（内存存储 + sync.Mutex）→ Handler（net/http 路由 + JSON 编解码）→ main（注册路由）。新增 `GetByID` 和 `Update` 两个 Store 方法，对应两个 Handler 端点。遵循现有代码模式。

**Tech Stack:** Go stdlib (net/http, encoding/json, sync, strconv, testing)

---

### Task 1: Store.GetByID — TDD 实现

**Files:**
- Modify: `evaluation/go/internal/store/store.go`
- Modify: `evaluation/go/internal/store/store_test.go`

- [ ] **Step 1.1: 编写 TestGetByID_NotFound（RED）**

在 `store_test.go` 文件末尾追加：

```go
func TestGetByID_NotFound(t *testing.T) {
	s := store.New()
	_, ok := s.GetByID(99)
	if ok {
		t.Error("expected GetByID to return false for non-existent id")
	}
}
```

- [ ] **Step 1.2: 编写 TestGetByID_Exists（RED）**

继续追加：

```go
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
```

- [ ] **Step 1.3: 运行测试，验证失败（RED）**

```bash
cd evaluation/go && go test ./... -run "TestGetByID" -v
```

期望输出：`undefined: Store.GetByID` 编译失败。

- [ ] **Step 1.4: 实现 Store.GetByID（GREEN）**

在 `store.go` 的 `Delete` 方法之后添加：

```go
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
```

- [ ] **Step 1.5: 运行测试，验证通过（GREEN）**

```bash
cd evaluation/go && go test ./... -run "TestGetByID|TestList|TestCreate|TestDelete" -v
```

期望：所有测试 PASS。

---

### Task 2: Store.Update — TDD 实现

**Files:**
- Modify: `evaluation/go/internal/store/store.go`
- Modify: `evaluation/go/internal/store/store_test.go`

- [ ] **Step 2.1: 编写 TestUpdate_NotFound（RED）**

```go
func TestUpdate_NotFound(t *testing.T) {
	s := store.New()
	title := "new"
	_, ok := s.Update(99, &title, nil)
	if ok {
		t.Error("expected Update to return false for non-existent id")
	}
}
```

- [ ] **Step 2.2: 编写 TestUpdate_Title（RED）**

```go
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
```

- [ ] **Step 2.3: 编写 TestUpdate_Done（RED）**

```go
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
```

- [ ] **Step 2.4: 编写 TestUpdate_Both（RED）**

```go
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
```

- [ ] **Step 2.5: 编写 TestUpdate_NoFields（RED）**

```go
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
```

- [ ] **Step 2.6: 运行测试，验证失败（RED）**

```bash
cd evaluation/go && go test ./... -run "TestUpdate" -v
```

期望输出：`undefined: Store.Update` 编译失败。

- [ ] **Step 2.7: 实现 Store.Update（GREEN）**

在 `store.go` 的 `GetByID` 方法之后添加：

```go
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
```

- [ ] **Step 2.8: 运行测试，验证通过（GREEN）**

```bash
cd evaluation/go && go test ./... -run "TestUpdate|TestGetByID|TestList|TestCreate|TestDelete" -v
```

期望：所有测试 PASS。

---

### Task 3: GET /tasks/{id} Handler

**Files:**
- Modify: `evaluation/go/internal/handler/handler.go`

- [ ] **Step 3.1: 注册路由并实现 getByID handler**

在 `handler.go` 的 `Register` 方法中添加路由：

```go
// 在 Register 方法内，mux.HandleFunc("DELETE /tasks/{id}", h.delete) 之后添加：
mux.HandleFunc("GET /tasks/{id}", h.getByID)
```

在 `delete` 方法之后添加 handler：

```go
func (h *Handler) getByID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, `{"error":"invalid id"}`, http.StatusBadRequest)
		return
	}
	task, ok := h.store.GetByID(id)
	if !ok {
		http.Error(w, `{"error":"not found"}`, http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(task)
}
```

- [ ] **Step 3.2: 编译验证**

```bash
cd evaluation/go && go build ./cmd/server/
```

期望：编译成功，无错误。

- [ ] **Step 3.3: 运行全量测试验证无回归**

```bash
cd evaluation/go && go test ./... -v
```

期望：所有已存在测试 PASS。

---

### Task 4: PATCH /tasks/{id} Handler

**Files:**
- Modify: `evaluation/go/internal/handler/handler.go`

- [ ] **Step 4.1: 注册路由并实现 update handler**

在 `Register` 方法中，`getByID` 路由之后添加：

```go
mux.HandleFunc("PATCH /tasks/{id}", h.update)
```

在 `getByID` 方法之后添加 handler：

```go
func (h *Handler) update(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, `{"error":"invalid id"}`, http.StatusBadRequest)
		return
	}
	var req struct {
		Title *string `json:"title"`
		Done  *bool   `json:"done"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, `{"error":"invalid json"}`, http.StatusBadRequest)
		return
	}
	if req.Title == nil && req.Done == nil {
		http.Error(w, `{"error":"no fields to update"}`, http.StatusBadRequest)
		return
	}
	task, ok := h.store.Update(id, req.Title, req.Done)
	if !ok {
		http.Error(w, `{"error":"not found"}`, http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(task)
}
```

- [ ] **Step 4.2: 编译验证**

```bash
cd evaluation/go && go build ./cmd/server/
```

期望：编译成功。

- [ ] **Step 4.3: 运行全量测试验证无回归**

```bash
cd evaluation/go && go test ./... -v
```

期望：所有测试 PASS。

---

### Task 5: 最终验证

- [ ] **Step 5.1: 运行全量测试**

```bash
cd evaluation/go && go test ./... -v
```

期望输出：11 个测试全部 PASS（原有 4 个 + 新增 7 个）。

- [ ] **Step 5.2: 确认 tasks.md 所有项已完成**

检查 `openspec/changes/complement-go-api/tasks.md` 中 7 项全部标记为 `[x]`。
