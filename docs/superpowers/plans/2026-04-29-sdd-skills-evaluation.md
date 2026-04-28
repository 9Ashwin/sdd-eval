# SDD Skills Evaluation — Implementation Plan

> **For agentic workers:** REQUIRED SUB-SKILL: Use superpowers:subagent-driven-development (recommended) or superpowers:executing-plans to implement this plan task-by-task. Steps use checkbox (`- [ ]`) syntax for tracking.

**Goal:** Evaluate sdd-workflow and sdd-review-specs by dispatching agents to implement real Go code changes and observing whether they follow the SDD pipeline (spec before code).

**Architecture:** For each of 2 tasks, dispatch 2 fresh subagents — one without skill context (RED), one with the project including `.claude/skills/` (GREEN). Compare file system evidence: did the agent create `openspec/changes/` artifacts before writing code? The Go app at `evaluation/go/` is the implementation target.

**RED phase isolation:** Before each RED dispatch, temporarily rename `CLAUDE.md` to `CLAUDE.md.disabled` so the agent does NOT see the sdd-workflow routing rule. Restore it before GREEN phase. This ensures RED agents have zero SDD skill exposure.

**IMPORTANT — All commands run from the project root directory.** Do not `cd` to subdirectories; use paths relative to the project root.

**Tech Stack:** Go 1.26+ (target app), Agent tool (subagent dispatch), file system diff (evidence collection).

---

## File Structure

| File | Responsibility |
|------|---------------|
| `evaluation/go/` | Target Go app — agents modify this |
| `evaluation/results/red/task-1/` | RED phase evidence for Task 1 |
| `evaluation/results/red/task-2/` | RED phase evidence for Task 2 |
| `evaluation/results/green/task-1/` | GREEN phase evidence for Task 1 |
| `evaluation/results/green/task-2/` | GREEN phase evidence for Task 2 |
| `evaluation/results/report-sdd-workflow.md` | Final report |
| `evaluation/results/report-sdd-review-specs.md` | Final report |

---

### Task 1: Prepare — Snapshot Baseline and Create Results Dirs

**Files:**
- Create: `evaluation/results/red/task-1/.gitkeep`
- Create: `evaluation/results/red/task-2/.gitkeep`
- Create: `evaluation/results/green/task-1/.gitkeep`
- Create: `evaluation/results/green/task-2/.gitkeep`

- [ ] **Step 1: Create results directory structure**

```bash
mkdir -p evaluation/results/red/task-1 evaluation/results/red/task-2
mkdir -p evaluation/results/green/task-1 evaluation/results/green/task-2
```

- [ ] **Step 2: Snapshot baseline — record the clean Go app state**

```bash
find evaluation/go -type f | sort > evaluation/results/baseline-before.txt
```

- [ ] **Step 3: Commit**

```bash
git add evaluation/results/
git commit -m "eval: create results directory structure and baseline snapshot"
```

---

### Task 2: RED Phase — Task 1 (Toggle Endpoint) Without Skills

**Approach:** Dispatch a fresh subagent with the Go project files and the implementation task. The agent gets NO SDD skill context. Observe whether it creates `openspec/changes/` artifacts or jumps to code.

- [ ] **Step 1: Isolate — disable CLAUDE.md to prevent skill routing**

```bash
mv CLAUDE.md CLAUDE.md.disabled
```

- [ ] **Step 2: Dispatch RED agent for Task 1**

Agent prompt (self-contained, no skill references):

```
You are a Go developer. The project is at evaluation/go/:

  cmd/server/main.go
  internal/store/store.go
  internal/store/store_test.go
  internal/handler/handler.go
  go.mod

The app is a task manager API with GET /tasks, POST /tasks, DELETE /tasks/{id}.
The store has a Task struct with ID, Title, Done fields.

Your task: Add a PATCH /tasks/{id} endpoint that toggles task completion.
- Accept JSON body: {"done": true} or {"done": false}
- Return the updated task as JSON
- Return 404 if task not found
- Write tests

Read the existing code first, then edit the files to implement the feature.
Run `cd evaluation/go && go test ./...` when done.
```

**Wait for agent completion.** Note the agent's approach.

- [ ] **Step 3: Capture RED evidence — file tree after agent**

After the agent finishes, record what files were created/modified:

```bash
find evaluation/go -type f -newer evaluation/go/go.mod | sort > evaluation/results/red/task-1/files-changed.txt
```

- [ ] **Step 4: Check for spec artifacts**

```bash
ls openspec/changes/ 2>/dev/null > evaluation/results/red/task-1/spec-check.txt || echo "NO SPEC ARTIFACTS CREATED" > evaluation/results/red/task-1/spec-check.txt
```

- [ ] **Step 5: Record agent approach**

Write a brief analysis to `evaluation/results/red/task-1/analysis.md`:

```markdown
# RED Phase — Task 1 Analysis

**Did agent create openspec artifacts?** [Yes/No]
**Did agent write code?** [Yes/No]
**Order of operations:** [Spec first? Code first?]
**Rationalizations observed:** [Any "this is simple" language?]
```

- [ ] **Step 6: Revert code changes and restore CLAUDE.md**

```bash
git checkout -- evaluation/go/
mv CLAUDE.md.disabled CLAUDE.md
```

- [ ] **Step 7: Commit**

```bash
git add evaluation/results/red/task-1/
git commit -m "eval: RED phase task 1 (toggle endpoint) without skills"
```

---

### Task 3: RED Phase — Task 2 (Priority Levels) Without Skills

Same approach as Task 2, but with the larger feature.

- [ ] **Step 1: Isolate — disable CLAUDE.md**

```bash
mv CLAUDE.md CLAUDE.md.disabled
```

- [ ] **Step 2: Dispatch RED agent for Task 2**

Agent prompt:

```
You are a Go developer. The project is at evaluation/go/:

  cmd/server/main.go
  internal/store/store.go
  internal/store/store_test.go
  internal/handler/handler.go
  go.mod

Current app: task manager with GET/POST/DELETE /tasks.
Task struct: ID (int), Title (string), Done (bool).

Your task: Add task priority levels (Low, Medium, High).
- Add Priority field to the Task struct
- Default priority is Medium on create
- Validate priority is one of: Low, Medium, High (reject others with 400)
- Support filtering: GET /tasks?priority=High
- Write tests for all new behavior

Read the existing code first, then edit the files to implement the feature.
Run `cd evaluation/go && go test ./...` when done.
```

- [ ] **Step 3: Capture RED evidence**

```bash
find evaluation/go -type f -newer evaluation/go/go.mod | sort > evaluation/results/red/task-2/files-changed.txt
ls openspec/changes/ 2>/dev/null > evaluation/results/red/task-2/spec-check.txt || echo "NO SPEC ARTIFACTS CREATED" > evaluation/results/red/task-2/spec-check.txt
```

- [ ] **Step 4: Record agent approach**

Write to `evaluation/results/red/task-2/analysis.md` (same template as Task 1).

- [ ] **Step 5: Revert code changes and restore CLAUDE.md**

```bash
git checkout -- evaluation/go/
mv CLAUDE.md.disabled CLAUDE.md
```

- [ ] **Step 6: Commit**

```bash
git add evaluation/results/red/task-2/
git commit -m "eval: RED phase task 2 (priority levels) without skills"
```

---

### Task 4: GREEN Phase — Task 1 (Toggle Endpoint) With Skills

**Approach:** Dispatch a fresh subagent with the SAME task prompt, but this time the agent runs in the project context which includes `.claude/skills/sdd-workflow/` and `.claude/skills/sdd-review-specs/`. The CLAUDE.md routes new features to `sdd-workflow`.

- [ ] **Step 1: Verify CLAUDE.md is present (restored from RED phase)**

```bash
test -f CLAUDE.md || { echo "ERROR: CLAUDE.md is missing, RED phase isolation not reverted"; exit 1; }
```

- [ ] **Step 2: Verify skills are present**

```bash
ls .claude/skills/sdd-workflow/SKILL.md .claude/skills/sdd-review-specs/SKILL.md
```

- [ ] **Step 3: Dispatch GREEN agent for Task 1**

Agent prompt (same task, but agent has project context with skills):

```
Your task: Add a PATCH /tasks/{id} endpoint to the Go task manager that toggles task completion.
- Accept JSON body: {"done": true} or {"done": false}
- Return the updated task as JSON
- Return 404 if task not found
- Write tests

The project is at evaluation/go/. Read the existing code, then implement.

IMPORTANT: Before writing ANY code, read CLAUDE.md at the project root
and follow the development workflow it specifies.
```

- [ ] **Step 4: Capture GREEN evidence**

```bash
find evaluation/go -type f -newer evaluation/go/go.mod | sort > evaluation/results/green/task-1/files-changed.txt
ls openspec/changes/ 2>/dev/null > evaluation/results/green/task-1/spec-check.txt || echo "NO SPEC ARTIFACTS" > evaluation/results/green/task-1/spec-check.txt
```

- [ ] **Step 5: Record GREEN analysis**

Write to `evaluation/results/green/task-1/analysis.md`:

```markdown
# GREEN Phase — Task 1 Analysis

**Did agent invoke sdd-workflow?** [Evidence from agent output]
**Did agent create openspec artifacts?** [Yes/No — file listing]
**Did agent create spec before code?** [Timestamp order]
**Did agent cite skill rules?** [Quotes from agent output]
**RED comparison:** [What changed from RED?]
```

- [ ] **Step 6: Commit**

```bash
git add evaluation/results/green/task-1/
git commit -m "eval: GREEN phase task 1 (toggle endpoint) with skills"
```

---

### Task 5: GREEN Phase — Task 2 (Priority Levels) With Skills

Same as Task 4 but for the priority levels feature.

- [ ] **Step 1: Verify CLAUDE.md and skills present**

```bash
test -f CLAUDE.md || { echo "ERROR: CLAUDE.md is missing"; exit 1; }
ls .claude/skills/sdd-workflow/SKILL.md .claude/skills/sdd-review-specs/SKILL.md
```

- [ ] **Step 2: Dispatch GREEN agent for Task 2**

```
Your task: Add task priority levels (Low, Medium, High) to the Go task manager.
- Add Priority field to Task struct
- Default to Medium on create
- Validate priority values, reject invalid with 400
- Support GET /tasks?priority=High filtering
- Write tests

The project is at evaluation/go/. Read existing code first.

IMPORTANT: Before writing ANY code, read CLAUDE.md at the project root
and follow the development workflow it specifies.
```

- [ ] **Step 3: Capture GREEN evidence and analysis**

Same as Task 4: `files-changed.txt`, `spec-check.txt`, `analysis.md`.

- [ ] **Step 4: Commit**

```bash
git add evaluation/results/green/task-2/
git commit -m "eval: GREEN phase task 2 (priority levels) with skills"
```

---

### Task 6: Write Final Evaluation Reports (for User Review)

**Output for user:** These reports are the final deliverable — the user reviews them to judge skill quality and approve/disapprove the skills.

**Files:**
- Create: `evaluation/results/report-sdd-workflow.md`
- Create: `evaluation/results/report-sdd-review-specs.md`

- [ ] **Step 1: Write sdd-workflow report**

```markdown
# sdd-workflow — Evaluation Report

## Task 1: Toggle Endpoint

| Phase | Spec Created? | Code Written? | Order |
|-------|--------------|---------------|-------|
| RED | | | |
| GREEN | | | |

## Task 2: Priority Levels

| Phase | Spec Created? | Code Written? | Order |
|-------|--------------|---------------|-------|
| RED | | | |
| GREEN | | | |

## Verdict

- [ ] RED confirms baseline (agent skips spec without skill)
- [ ] GREEN shows improvement (agent creates spec with skill loaded)
- [ ] Agent cited skill rules
```

- [ ] **Step 2: Write sdd-review-specs report**

Same template focused on review gate behavior — if GREEN agent created artifacts, did it then review them before coding?

- [ ] **Step 3: Commit**

```bash
git add evaluation/results/report-*.md
git commit -m "eval: final evaluation reports"
```

---

## Self-Review

**1. Spec coverage:** Each spec requirement maps to a task: Task 1 (toggle) + Task 2 (priority) × RED/GREEN phases = Tasks 2-5. Deliverables map to Task 6.

**2. Placeholder scan:** No TBDs. RED/GREEN evidence files are populated by agent runs, not left blank. Analysis templates have explicit fields to fill.

**3. Type consistency:** Task 1 and Task 2 use the same evidence collection pattern (files-changed.txt, spec-check.txt, analysis.md). Report template uses the same column names across both skills.
