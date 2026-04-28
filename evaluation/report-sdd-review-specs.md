# sdd-review-specs — Final Evaluation Report

## Methodology

Superpowers TDD-for-docs: 3 multi-pressure RED scenarios → 3 GREEN scenarios (skill loaded) → REFACTOR patches. All subagents dispatched fresh with no shared context.

**Date:** 2026-04-29  
**Evaluator:** Automated subagent dispatch

## RED Phase — Baseline (No Skill)

| Scenario | Pressures | Choice | Review Quality |
|----------|-----------|--------|---------------|
| A: AI usually right | Time + Exhaustion + Pragmatic | **B** (Tier 1) | Partial |
| B: Looks reasonable | Time + Sunk cost + Authority | **B** (partial) | Partial |
| C: Review too long | Time + Economic + Exhaustion | **C** (Tier 2) | Full ✓ |

**Baseline outcome:** 1/3 chose full Tier 2 review. 2/3 chose partial review (Tier 1 or targeted scan).

## GREEN Phase — With Skill Loaded

| Scenario | RED | GREEN | Changed | Skill Cited |
|----------|-----|-------|---------|-------------|
| A: AI reliable | B | **C** | +1 | Tier table + 3 rationalizations |
| B: Looks fine | B | **C** | +1 | Rationalizations + Gate Function |
| C: Crunch mode | C | **C** | Held | Tier table + Iron Law |

**Skill outcome:** 3/3 chose full Tier 2. The Tier classification table was the decisive mechanism — agents could not classify "new feature" as anything but Tier 2.

## Rationalizations Discovered (RED Phase)

| New Rationalization | Type |
|---------------------|------|
| "Full review when tired is performative — fatigue degrades thoroughness" | Low-Yield Review |
| "A targeted partial review is honest and productive" | Partial as Virtue |
| "Unrealistic to do full review in [time window]" | Time-Boxed Review |

## REFACTOR — Patches Applied

2 new entries added to Common Rationalizations table:

1. **"A targeted partial review is honest pragmatism"** → "Partial review is partial compliance. If it's Tier 2 by classification, Tier 1 review is insufficient regardless of how 'honest' it feels."
2. **"Full review when tired is performative"** → "Fatigue degrades thoroughness, but skipping review entirely guarantees zero thoroughness. If you're too tired, defer — don't rationalize a skim."

## Bulletproof Signs Check

| Sign | Status |
|------|--------|
| Agent chooses correct option under maximum pressure | **Pass** — 3/3 correct after skill loaded |
| Agent cites skill sections as justification | **Pass** — Tier table + Rationalizations most cited |
| Agent acknowledges temptation but follows rule | **Pass** — "The PM asking is textbook 'Review takes too long'" |
| Skill survives combined 3+ pressure scenarios | **Pass** — Even crunch mode (3 pressures) held correct |

## Verdict

The skill is highly effective. The Tier classification system acts as an objective gate that resists rationalization — agents cannot reclassify a "new feature" as Tier 1 without explicitly contradicting the skill text. The rationalizations table provides counter-arguments for every common excuse.

**Strongest mechanism:** Tier table + "If unsure, default to Tier 2" — creates an objective classification that resists subjective downgrading.  
**Strongest table:** Common Failures — prevents "I looked at it" from being claimed as "reviewed."  
**Gap (now patched):** The "partial review is honest" rationalization wasn't directly countered.
