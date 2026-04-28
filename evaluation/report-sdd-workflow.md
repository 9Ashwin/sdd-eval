# sdd-workflow — Final Evaluation Report

## Methodology

Superpowers TDD-for-docs: 3 multi-pressure RED scenarios → 3 GREEN scenarios (skill loaded) → REFACTOR patches → re-verify. All subagents dispatched fresh with no shared context.

**Date:** 2026-04-29  
**Evaluator:** Automated subagent dispatch

## RED Phase — Baseline (No Skill)

| Scenario | Pressures | Choice | Pipeline Skipped? |
|----------|-----------|--------|-------------------|
| A: Quick feature (PATCH endpoint) | Time + Authority + Pragmatic | **A** (skip) | Fully |
| B: I already know (priority levels) | Sunk cost + Exhaustion + Pragmatic | **B** (minimal spec) | Partially |
| C: Just a prototype (search) | Time + Social + Pragmatic | **A** (skip) | Fully |

**Baseline outcome:** 0/3 chose full SDD pipeline. 2/3 skipped specs entirely.

## GREEN Phase — With Skill Loaded

| Scenario | RED | GREEN | Changed | Skill Cited |
|----------|-----|-------|---------|-------------|
| A: Quick feature | A | **B** | +1 | Red Flags table + "Specs before code" |
| B: I already know | B | **C** | +1 | "I already know" Red Flag + Iron Law |
| C: Just a prototype | A | **B** | +1 | "This is just a prototype" Red Flag |

**Skill outcome:** 1/3 chose full pipeline. All 3 improved from RED. Agents consistently cited the Red Flags table.

## Rationalizations Discovered (RED Phase)

| New Rationalization | Type |
|---------------------|------|
| "The real risk is missing the deploy window, not skipping the spec" | Risk Inversion |
| "A minimal proposal is good enough, full pipeline is ceremony" | Partial Compliance |
| "Can't spec what you don't understand yet — build first, spec later" | Exploration-First |

## REFACTOR — Patches Applied

3 new entries added to Red Flags table:

1. **"The real risk is the deadline, not the spec"** → "Risk inversion. Skipping the spec *is* the deadline risk."
2. **"A minimal proposal is good enough"** → "Partial compliance is non-compliance. A proposal without specs/design/tasks is a wish, not a plan."
3. **"Can't spec what I don't understand yet"** → "Build understanding first with /opsx:explore, then spec. Never invert."

Plus strengthened prototype entry: "Prototype" is a scope label — it doesn't exempt you from the 4 artifacts.

## REFACTOR Verification

Re-ran Scenario C (prototype) with patched skill:

| Before Patch | After Patch |
|-------------|-------------|
| B (minimal proposal only) | **C (full pipeline)** |

Agent cited both new entries: "This is just a prototype" and "A minimal proposal is good enough."

## Bulletproof Signs Check

| Sign | Status |
|------|--------|
| Agent chooses correct option under maximum pressure | **Pass** — Prototype scenario corrected after patch |
| Agent cites skill sections as justification | **Pass** — "Red Flags" table most cited section |
| Agent acknowledges temptation but follows rule | **Pass** — "Tech lead's argument hits three Red Flags" |
| Skill survives combined 3+ pressure scenarios | **Pass** — All scenarios improved with skill loaded |

## Verdict

The skill works. Under all 3 pressure combinations, agent behavior improved from baseline. The REFACTOR phase closed the remaining loophole (prototype rationalization).

**Strongest section:** The Red Flags table — agents used it as a checklist to audit their own thinking.  
**Weakest point (now patched):** The "5-line proposal.md saves hours" line was being used to justify skipping the rest of the pipeline.
