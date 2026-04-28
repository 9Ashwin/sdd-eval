# SDD Skills Evaluation — Design Spec

## Purpose

Evaluate whether `sdd-workflow` and `sdd-review-specs` resist rationalization under combined pressure, following Superpowers' TDD-for-docs methodology. This is a Go-language evaluation targeting the two skills installed in this project.

## Evaluation Target

| Skill | Type | What It Enforces |
|-------|------|-----------------|
| `sdd-workflow` | Discipline-Enforcing | Always create spec artifacts before writing code |
| `sdd-review-specs` | Discipline-Enforcing | Always review AI-generated spec artifacts before implementation |

## Methodology: TDD-for-Docs

Three phases mapped from Test-Driven Development:

| Phase | What Happens | Success Signal |
|-------|-------------|----------------|
| **RED** | Dispatch fresh subagents with pressure scenarios, NO skill context. Capture choices + rationalizations verbatim. | Agent skips the spec/review process. Rationalizations are documented. |
| **GREEN** | Re-dispatch SAME scenarios with full skill content prepended. Compare choices. | Agent improves at least 1 level (A→B or B→C) from RED baseline. Agent explicitly cites skill sections. |
| **REFACTOR** | Identify rationalizations that survived GREEN. Patch skill text. Re-run failed scenarios. | Patched scenarios switch to correct choice (C). Agent cites new patch entries. |

## Pressure Scenario Design

Each scenario combines **3+ pressure types** from the Superpowers taxonomy:

| Pressure | Manifestation |
|----------|--------------|
| Time | Deploy window closing, sprint ending, standup in 10 minutes |
| Authority | Tech lead says skip it, manager demands speed |
| Sunk cost | Hours of discussion/iteration already spent |
| Exhaustion | End of day, already reviewed many things |
| Economic | Velocity at risk, sprint commitments slipping |
| Social | Team wants progress, looking dogmatic |
| Pragmatic | "Being practical" vs following process |

**6 scenarios total** (3 per skill):

### sdd-workflow Scenarios

| # | Name | Pressures | Test |
|---|------|-----------|------|
| A | Quick feature | Time + Authority + Pragmatic | Add PATCH endpoint in 45 min — skip spec? |
| B | I already know | Sunk cost + Exhaustion + Pragmatic | 2hr meeting decided everything at 6pm — skip spec? |
| C | Just a prototype | Time + Social + Pragmatic | Demo tomorrow, throw-away code — skip spec? |

### sdd-review-specs Scenarios

| # | Name | Pressures | Test |
|---|------|-----------|------|
| A | AI usually right | Time + Exhaustion + Pragmatic | AI reliable, tired — skim or full review? |
| B | Looks reasonable | Time + Sunk cost + Authority | Tech lead said "fine," 3 iterations — skip review? |
| C | Review too long | Time + Economic + Exhaustion | Sprint crunch, 28 checklist items — skip review? |

## Pass Criteria

**Per-skill pass (all must hold):**

1. **GREEN improvement:** Every scenario improves at least 1 level from RED baseline (A→B or B→C)
2. **Skill citation:** Agent references specific skill sections in every GREEN response
3. **REFACTOR closure:** Any scenario not at C after GREEN must reach C after REFACTOR patches
4. **Bulletproof signs:** Agent acknowledges temptation, cites skill as counter, chooses correctly

**Failing means:** The skill's text isn't strong enough to overcome that pressure combination. The rationalization table needs new entries, or the Iron Law needs stronger language.

## Deliverables

| Artifact | Location | Purpose |
|----------|----------|---------|
| Pressure scenarios (×6) | `evaluation/go/scenarios/<skill>-red-*.md` | Reusable test fixtures |
| RED results | `evaluation/go/scenarios/<skill>-red-results.md` | Baseline behavior evidence |
| GREEN results | `evaluation/go/scenarios/<skill>-green-results.md` | Skill effectiveness evidence |
| REFACTOR patches | `.claude/skills/<skill>/SKILL.md` (commits) | Skill improvements |
| Final reports | `evaluation/report-<skill>.md` | Summary with verdict and bulletproof check |

## Non-Goals

- Python or other language evaluations (future, under `evaluation/py/`)
- Testing the OpenSpec CLI tools themselves
- Testing skill triggering (does the agent auto-load the skill?) — that's Superpowers' `tests/skill-triggering/` domain
- Measuring token usage or cost
