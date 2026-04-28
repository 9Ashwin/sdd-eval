# sdd-workflow — Evaluation Report

## Task 1: Toggle Endpoint

| Phase | Spec Created? | Code Written? | Order |
|-------|--------------|---------------|-------|
| RED | No | Yes | Code directly — agent read existing files, then immediately edited implementation and test files |
| GREEN | Yes — `openspec/changes/patch-task-toggle/` with proposal.md, design.md, tasks.md, spec.md | Yes | Spec first — agent created proposal and tasks before writing any code; reported "6/6 tasks complete" |

## Task 2: Priority Levels

| Phase | Spec Created? | Code Written? | Order |
|-------|--------------|---------------|-------|
| RED | No | Yes | Code directly — agent read files, implemented feature in a single pass, created handler_test.go |
| GREEN | Yes — `openspec/changes/task-priority-levels/` with full artifact set | Yes | Spec first — agent created spec artifacts before implementation; reported "Schema: spec-driven" and "8/8 tasks complete" |

## Verdict

- [x] RED confirms baseline (agent skips spec without skill)
- [x] GREEN shows improvement (agent creates spec with skill loaded)
- [x] Agent cited skill rules — task tracking with [x] notation, progress reporting format consistent with SDD workflow

### Key findings

1. **CLAUDE.md routing works:** When CLAUDE.md is present with the sdd-workflow routing rule for new features, agents consistently create spec artifacts before writing code. Without it, agents go directly to implementation.

2. **Both RED agents behaved identically:** Neither created `openspec/changes/` artifacts. Both went directly to code — reading existing files and editing them in a single pass.

3. **Both GREEN agents behaved identically:** Both created full proposal → tasks → spec → implementation pipelines. Both tracked task progress.

4. **Test quality:** GREEN agents wrote more comprehensive tests. RED Task 1 agent wrote 3 store tests; GREEN Task 1 agent wrote 6 tests (3 store + 6 handler). RED Task 2 agent wrote handler tests; GREEN Task 2 agent wrote 14 tests total, preserving existing toggle tests.
