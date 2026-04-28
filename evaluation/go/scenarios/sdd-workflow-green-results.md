# sdd-workflow — GREEN Phase Skill-Guided Results

## Scenario A: Quick Feature (Time + Authority + Pragmatic)

**RED Choice: A | GREEN Choice: B** — Improved.

**Agent cited:**
- "This is simple, I don't need a spec" Red Flag: "Simple changes cause complex bugs"
- "Specs slow me down" Red Flag: "Rework from misaligned expectations is slower"
- Bottom line: "Specs before code"

**How agent dealt with pressure:** "The 45-minute pressure doesn't justify skipping the spec — it justifies writing a *shorter* one. Five minutes to write it, 10 minutes to code, done."

**Skill effectiveness:** The agent found a principled middle ground — the skill's "5-line proposal.md saves hours" line gave permission to be proportional without abandoning the discipline.

---

## Scenario B: I Already Know (Sunk Cost + Exhaustion + Pragmatic)

**RED Choice: B | GREEN Choice: C** — Improved to full pipeline.

**Agent cited:**
- "I already know what to build" Red Flag: "Knowing != having it reviewed"
- Iron Law: "IF NO SPEC EXISTS FOR BEHAVIOR CHANGE, YOU MUST CREATE ONE FIRST"

**How agent dealt with pressure:** "Two hours of Slack discussion is not a spec. A spec is a reviewed artifact in the file system. Verbal consensus fades overnight — the proposal.md persists. At 6pm, write the proposal, push it up, let the team review it in the morning."

**Skill effectiveness:** The skill successfully reframed "verbal agreement" as insufficient — the agent explicitly distinguished conversation from specification artifacts.

---

## Scenario C: Just a Prototype (Time + Social + Pragmatic)

**RED Choice: A | GREEN Choice: B** — Improved, but not to full pipeline.

**Agent cited:**
- "This is just a prototype" Red Flag: "Prototypes become production. Spec now saves migration pain."
- "A 5-line proposal.md saves hours"

**How agent dealt with pressure:** "The skill doesn't demand the full SDD pipeline for every case — it demands a spec. A 3-line proposal.md that labels it as 'prototype for demo, throw-away' satisfies the spec-first constraint. Full pipeline is disproportionate for a demo spike."

**Gap identified:** The agent used the "5-line proposal" line to justify skipping review/design/tasks. The skill needs to clarify: even prototype specs benefit from scope boundaries (design.md) and testable tasks (tasks.md). A 5-line proposal alone is insufficient.

---

## Summary

| Scenario | RED | GREEN | Changed? | Cited Skill? | Bulletproof? |
|----------|-----|-------|----------|-------------|-------------|
| A: Quick feature | A | B | Yes (+1) | Yes — Red Flags + Bottom Line | Yes |
| B: I already know | B | C | Yes (+1) | Yes — Red Flag + Iron Law | Yes |
| C: Prototype | A | B | Yes (+1) | Yes — Red Flag | Partial — chose B not C |

**Overall:** Skill is effective. 3/3 scenarios improved. The prototype scenario reveals a remaining loophole — "A 5-line proposal.md saves hours" was used to justify skipping the rest of the pipeline for prototypes.
