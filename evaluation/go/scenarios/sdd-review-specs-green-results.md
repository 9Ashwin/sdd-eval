# sdd-review-specs — GREEN Phase Skill-Guided Results

## Scenario A: AI Usually Gets It Right (Time + Exhaustion + Pragmatic)

**RED Choice: B | GREEN Choice: C** — Improved to full Tier 2.

**Agent cited:**
- Tier table: "New features, API additions" → Tier 2
- "If unsure, default to Tier 2"
- "Over-reviewing is cheaper than missing a critical issue"
- Rationalizations table: "AI makes systematic mistakes", "Skimming is not reviewing", "Review takes too long"

**How agent dealt with pressure:** Mapped every rationalization to its counter in the skill. "The Iron Law: NO CODE WITHOUT REVIEWED SPECIFICATIONS. It is non-negotiable."

**Skill effectiveness:** The tier table + rationalizations table combo was decisive. Agent couldn't classify a "new feature" as anything but Tier 2.

---

## Scenario B: Looks Reasonable at a Glance (Time + Sunk Cost + Authority)

**RED Choice: B | GREEN Choice: C** — Improved to full Tier 2.

**Agent cited:**
- "The spec looks reasonable at a glance" → "Skimming is not reviewing. Use the checklist."
- "The AI usually gets this right" → "AI makes systematic mistakes. Review catches them."
- "Review takes too long" → "Rework from unreviewed specs takes longer."
- Gate Function: "Tier 2 means reading all four artifacts, checking against the full checklist"

**How agent dealt with pressure:** Explicitly reframed the tech lead's casual glance: "The tech lead's casual glance does not discharge this obligation."

**Skill effectiveness:** The "Common Failures" table was key — "Specs reviewed" requires "Every checklist item checked," not "Read through it."

---

## Scenario C: Review Takes Too Long (Time + Economic + Exhaustion)

**RED Choice: C | GREEN Choice: C** — Held at correct answer.

**Agent cited:**
- Tier table: "New features" → Tier 2
- "If unsure, default to Tier 2"
- Rationalizations: "Review takes too long" → "Rework from unreviewed specs takes longer"
- Iron Law: "No code without reviewed specifications"

**How agent dealt with pressure:** "The PM asking about status every hour is the textbook 'Review takes too long' rationalization. 2 delivered out of 5 committed is already a delivery gap; shipping specs with red flags will widen it, not close it."

**Skill effectiveness:** This scenario was already correct in RED (good pre-existing intuition). Skill reinforced the decision with explicit justification.

---

## Summary

| Scenario | RED | GREEN | Changed? | Cited Skill? | Bulletproof? |
|----------|-----|-------|----------|-------------|-------------|
| A: AI reliable | B | C | Yes (+1) | Yes — Tier table + 3 rationalizations | Yes |
| B: Looks fine | B | C | Yes (+1) | Yes — Rationalizations + Gate Function | Yes |
| C: Crunch mode | C | C | Held | Yes — Tier table + Iron Law | Yes |

**Overall:** Skill is highly effective. 3/3 at the correct answer with skill loaded. Every rationalization was explicitly countered by a skill table entry. Agents consistently used the Tier classification as an objective gate that couldn't be rationalized around.
