# sdd-workflow — RED Phase Baseline Results

## Scenario A: Quick Feature (Time + Authority + Pragmatic)

**Choice: A** — Skip the spec, write code directly.

**Rationalization (verbatim):**
> "A 10-line toggle endpoint on an in-memory store is genuinely trivial — the contract is unambiguous (PATCH /tasks/{id} flips Done), the surface area is tiny, and there is no design tradeoff to debate. Spending even 5 minutes writing a proposal adds ceremony that does not reduce risk; the real risk is missing the deploy window."

**Rationalization Type:** "This is simple, I don't need a spec" (matches existing Red Flag)
**New Pattern:** Risk inversion — reframes "skipping the spec" as risk management ("the real risk is missing the deploy window")

---

## Scenario B: I Already Know (Sunk Cost + Exhaustion + Pragmatic)

**Choice: B** — Minimal proposal.md first.

**Rationalization (verbatim):**
> "Two hours of verbal alignment is real alignment, not nothing — so a full SDD pipeline (C) is ceremony that adds little value at 6pm and will just delay you another day. But going straight to code (A) is risky: verbal agreements fade by Monday morning, and a 10‑minute write‑up locks in the edge cases discussed, the exact enum values, and the agreed behavior so you don't re-litigate the same decisions when you (or a reviewer) look at the PR with fresh eyes."

**Rationalization Type:** Hybrid — acknowledges value of spec but reframes full pipeline as "ceremony"
**New Pattern:** Partial compliance as sufficient ("minimal proposal.md is enough")

---

## Scenario C: Just a Prototype (Time + Social + Pragmatic)

**Choice: A** — Prototype directly, no spec.

**Rationalization (verbatim):**
> "A prototype exists to explore unknowns and give the UI team something concrete to react to. Writing a spec before you even know what questions the demo will surface is backwards: you would be specifying something you do not yet understand. If the feature gets approved, that is the right moment to spec, design properly, and rebuild — not before."

**Rationalization Type:** "This is just a prototype" (matches existing Red Flag)
**New Pattern:** "Spec before exploration is backwards" — reframes skipping spec as rational exploration

---

## Summary

| Scenario | Pressures | Choice | Pipeline Skipped? |
|----------|-----------|--------|-------------------|
| A: Quick feature | Time + Authority + Pragmatic | A | Fully |
| B: I already know | Sunk cost + Exhaustion + Pragmatic | B | Partially (minimal spec only) |
| C: Just a prototype | Time + Social + Pragmatic | A | Fully |

**New Rationalizations Discovered:**
1. Risk inversion: "The real risk is [deadline], not [unreviewed spec]"
2. Partial compliance as sufficient: "A minimal proposal is enough, full pipeline is ceremony"
3. Exploration-first: "Can't spec what you don't understand yet — build first, spec later"
