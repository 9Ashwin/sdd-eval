# SDD Skills Evaluation — Design Spec

## Purpose

Behavior-based evaluation of `sdd-workflow` and `sdd-review-specs`. Instead of asking agents hypothetical questions ("What would you do?"), we give them real implementation tasks against the Go codebase and observe whether they follow the SDD process — spec before code, review before execution.

## Evaluation Targets

Only these two skills:

| Skill | Role | Expected Behavior |
|-------|------|------------------|
| `sdd-workflow` | Router | Classify task → route to propose/explore/brainstorming → enforce spec-first |
| `sdd-review-specs` | Review gate | After artifacts generated, stop and review before allowing code |

## What We Actually Measure

We don't ask. We observe file system evidence:

| Signal | Evidence |
|--------|----------|
| Agent created spec before code? | `openspec/changes/<name>/` exists with proposal.md, specs/, design.md, tasks.md; timestamps before code changes |
| Agent followed review gate? | Evidence of checklist usage, feedback on artifacts before implementation |
| Agent skipped SDD? | Code files modified/created but no `openspec/changes/` directory |

## Tasks

Two implementation tasks of increasing complexity:

### Task 1: Toggle endpoint (small)

Add `PATCH /tasks/{id}` to the Go task manager. The endpoint accepts a JSON body `{"done": true/false}` and toggles the task's Done field. This is ~30 lines: one store method, one handler, one test.

**Tests the "This is too simple for a spec" rationalization.**

### Task 2: Priority levels (medium)

Add task priority (Low/Medium/High) to the task model. New task field, validation on create, filter by priority on list. ~80 lines: store changes, handler changes, tests.

**Tests whether agent runs full propose→review→apply pipeline for a multi-touchpoint feature.**

## Phases

### RED Phase — Baseline

Dispatch fresh subagent with the task prompt. No SDD skill context. Agent has access to the Go project files. Ask it to implement the task.

Observe:
- Does the agent create any `openspec/` artifacts?
- Does the agent jump straight to code?
- What rationalizations does the agent give (if any)?

Record: file tree before/after, agent's approach description.

### GREEN Phase — With Skills

Dispatch fresh subagent with the SAME task prompt, but the project's `.claude/skills/` includes `sdd-workflow` and `sdd-review-specs`. The agent inherits them via the project context.

Observe:
- Does the agent invoke `sdd-workflow`?
- Does the agent create `openspec/changes/<name>/` artifacts?
- Does the agent pause for review before coding?
- Does the agent cite skill rules?

Record: file tree before/after, artifacts created, order of operations.

### REFACTOR Phase

If GREEN phase shows the skill failed to change behavior (agent still coded first), identify the loophole and patch the skill. Re-run.

## Pass Criteria

| Criteria | Threshold |
|----------|-----------|
| RED confirms baseline: agent skips spec | At least 1 of 2 tasks coded without spec |
| GREEN shows improvement | At least 1 of 2 tasks creates spec artifacts before code |
| Skill citation | Agent references sdd-workflow or sdd-review-specs in justification |
| No false positives | GREEN agent doesn't claim "I reviewed" without actual artifact changes |

## Deliverables

| Artifact | Location |
|----------|----------|
| This spec | `docs/superpowers/specs/2026-04-29-sdd-skills-evaluation-design.md` |
| Implementation plan | `docs/superpowers/plans/2026-04-29-sdd-skills-evaluation.md` |
| RED results | `evaluation/results/red/` |
| GREEN results | `evaluation/results/green/` |
| Final reports | `evaluation/results/report-sdd-workflow.md`, `evaluation/results/report-sdd-review-specs.md` |

## Non-Goals

- Testing OpenSpec CLI tool correctness
- Testing skill auto-triggering (that's Superpowers' `tests/skill-triggering/` domain)
- Python or other language evaluations
- Measuring token usage or cost
