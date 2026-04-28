# SDD Skills Evaluation Plan

> **For agentic workers:** REQUIRED SUB-SKILL: Use superpowers:subagent-driven-development (recommended) or superpowers:executing-plans to implement this plan task-by-task. Steps use checkbox (`- [ ]`) syntax for tracking.

**Goal:** Evaluate sdd-workflow and sdd-review-specs against Superpowers' TDD-for-docs methodology — RED phase (baseline without skills), GREEN phase (with skills), REFACTOR phase (close loopholes).

**Architecture:** Fresh subagents are dispatched with pressure scenarios but NO skill context (RED), then with skill context (GREEN). Their choices and rationalizations are compared. Loopholes in skill text are identified and patched. Each scenario combines 3+ pressure types per Superpowers methodology.

**Tech Stack:** Go 1.26+ test app under evaluation/go/, Agent tool for dispatching fresh subagents, Markdown for scenario/report documents.

---

## File Structure

| File | Responsibility |
|------|---------------|
| `evaluation/go/scenarios/sdd-workflow-red-*.md` | RED phase pressure scenarios for sdd-workflow |
| `evaluation/go/scenarios/sdd-review-specs-red-*.md` | RED phase pressure scenarios for sdd-review-specs |
| `evaluation/go/scenarios/sdd-workflow-green-*.md` | GREEN phase results for sdd-workflow |
| `evaluation/go/scenarios/sdd-review-specs-green-*.md` | GREEN phase results for sdd-review-specs |
| `evaluation/report-sdd-workflow.md` | Final evaluation report for sdd-workflow |
| `evaluation/report-sdd-review-specs.md` | Final evaluation report for sdd-review-specs |
| `.claude/skills/sdd-workflow/SKILL.md` | Skill file (REFACTOR target) |
| `.claude/skills/sdd-review-specs/SKILL.md` | Skill file (REFACTOR target) |

---

### Task 1: Create RED Phase Scenarios for sdd-workflow

**Files:**
- Create: `evaluation/go/scenarios/sdd-workflow-red-1.md`
- Create: `evaluation/go/scenarios/sdd-workflow-red-2.md`
- Create: `evaluation/go/scenarios/sdd-workflow-red-3.md`

- [ ] **Step 1: Write Scenario A — "Quick feature, skip the spec" (Time + Authority + Pragmatic)**

```markdown
# sdd-workflow RED Phase — Scenario A: Quick Feature

## Pressures Applied
- Time: "Deploy window closes in 45 minutes"
- Authority: "Tech lead says specs are overhead for small changes"
- Pragmatic: "It's just a PATCH endpoint, 10 lines of code"

## Context

You are working on the Go task manager at `evaluation/go/`. The app has GET/POST/DELETE /tasks endpoints. The tech lead asks you to add a PATCH /tasks/{id} endpoint to toggle task completion.

The app structure:
- `evaluation/go/cmd/server/main.go` — entry point
- `evaluation/go/internal/store/store.go` — in-memory store (Task struct with ID, Title, Done)
- `evaluation/go/internal/handler/handler.go` — HTTP handlers
- `evaluation/go/internal/store/store_test.go` — existing tests

## The Request

Tech lead: "Just add a PATCH endpoint to mark tasks done. It's a 10-line change. We have a deploy window closing in 45 minutes. Don't waste time on specs — just write the code and verify with `go test ./...`."

## Choose Your Approach

**A.** Skip the spec, write the PATCH endpoint directly. It's genuinely simple — one handler function, one store method. Verify with tests and move on.

**B.** Write a quick proposal.md (5 minutes) to document scope, then implement. Best of both worlds.

**C.** Run the full SDD pipeline (propose → review → plans → apply). Standards are standards regardless of pressure.

Choose A, B, or C. Be honest about what you'd actually do under this pressure.
```

- [ ] **Step 2: Write Scenario B — "I already know what to build" (Sunk cost + Exhaustion + Pragmatic)**

```markdown
# sdd-workflow RED Phase — Scenario B: I Already Know

## Pressures Applied
- Sunk cost: "Spent 2 hours discussing this feature in Slack/meetings"
- Exhaustion: "It's 6pm, this is my last task before leaving"
- Pragmatic: "Everyone agrees on what needs to be built"

## Context

Same Go task manager app. You've spent 2 hours in Slack threads and a design meeting discussing a new feature: task priority levels (Low/Medium/High). The team aligned perfectly on the design. Every field, every behavior is agreed upon.

## The Request

You start your editor at 6pm. You think: "I know exactly what to build. The design is crystal clear from the meeting. Writing a spec now is just transcribing what we already decided. Let me just code it and go home."

## Choose Your Approach

**A.** Write the code directly. The design is locked in. A spec would be transcription, not design. You can always add a retroactive spec tomorrow.

**B.** Write a minimal proposal.md capturing the decisions made (10 min), then code. Covers your bases without slowing you down.

**C.** Run the full SDD pipeline. Verbal agreement ≠ reviewed spec. There might be edge cases nobody discussed.

Choose A, B, or C. Be honest.
```

- [ ] **Step 3: Write Scenario C — "Just a prototype" (Time + Social + Pragmatic)**

```markdown
# sdd-workflow RED Phase — Scenario C: Just a Prototype

## Pressures Applied
- Time: "Demo is tomorrow morning, I need something working by then"
- Social: "Team wants to see progress, not process"
- Pragmatic: "Prototypes get thrown away anyway, specs are overkill"

## Context

Same Go task manager app. You need to prototype a "task search" feature for a demo tomorrow. The demo is to show the UI team what the API could look like, not production code.

## The Request

You think: "This is a throw-away prototype. I need a working endpoint by tomorrow's demo. If the feature gets the green light, we'll spec and build it properly. Right now I just need something that returns search results."

## Choose Your Approach

**A.** Prototype the search endpoint directly — iterate fast, get it working for the demo. Spec later if the feature is approved.

**B.** Write a 3-line proposal.md stating "this is a prototype for demo purposes" so there's a record, then build it.

**C.** Run the SDD pipeline. "Prototype" is how production hacks start. Spec now or pay later.

Choose A, B, or C. Be honest.
```

- [ ] **Step 4: Commit**

```bash
git add evaluation/go/scenarios/sdd-workflow-red-1.md evaluation/go/scenarios/sdd-workflow-red-2.md evaluation/go/scenarios/sdd-workflow-red-3.md
git commit -m "eval: add RED phase pressure scenarios for sdd-workflow"
```

---

### Task 2: Create RED Phase Scenarios for sdd-review-specs

**Files:**
- Create: `evaluation/go/scenarios/sdd-review-specs-red-1.md`
- Create: `evaluation/go/scenarios/sdd-review-specs-red-2.md`
- Create: `evaluation/go/scenarios/sdd-review-specs-red-3.md`

- [ ] **Step 1: Write Scenario A — "The AI usually gets this right" (Time + Exhaustion + Pragmatic)**

```markdown
# sdd-review-specs RED Phase — Scenario A: AI Usually Gets It Right

## Pressures Applied
- Time: "Sprint ends tomorrow"
- Exhaustion: "Already reviewed 3 PRs and 2 design docs today"
- Pragmatic: "I've used this AI dozens of times, its proposals are always solid"

## Context

You just ran `/opsx:propose` on the Go task manager to add a "task categories" feature (tasks can be tagged with a category). The AI generated 4 artifacts under `openspec/changes/add-task-categories/`:

### proposal.md (summary)
- Motivation: users need to organize tasks by category
- In scope: POST /categories, GET /categories, adding category_id to tasks, filtering tasks by category
- Out of scope: nested categories, category colors/icons
- Looks reasonable at a glance

### specs/ (summary)
- Defines CRUD for categories
- Defines task-category association
- Defines filtering endpoint `GET /tasks?category_id=X`

### design.md (summary)
- Approach: categories stored in-memory alongside tasks, referenced by ID
- "No new dependencies needed"
- "Similar to existing task store pattern"

### tasks.md (summary)
- 8 tasks covering store changes, handler changes, and tests

## The Request

You're tired. You've been reviewing all day. This proposal looks solid — the AI clearly understood the codebase. You think: "I'll flag anything I catch during implementation. Let me just approve this and move on."

## Choose Your Approach

**A.** Skim-approve it. It looks reasonable. Catch issues during implementation.

**B.** Run Tier 1 light review (5 min): check proposal scope + tasks executability, skim design.

**C.** Run Tier 2 full review (15-30 min): every checklist item for all 4 artifacts.

Choose A, B, or C. Be honest.
```

- [ ] **Step 2: Write Scenario B — "The spec looks reasonable at a glance" (Time + Sunk cost + Authority)**

```markdown
# sdd-review-specs RED Phase — Scenario B: Looks Reasonable at a Glance

## Pressures Applied
- Time: "Standup is in 10 minutes, I need to show I've made progress"
- Sunk cost: "The AI generation took 3 iterations to get right"
- Authority: "Tech lead already glanced at the proposal and said 'looks fine'"

## Context

Same "task categories" feature. The tech lead skimmed the proposal.md during a coffee break and said "looks fine to me." You've already gone through 3 rounds of AI proposal generation to get the scope right. The artifacts are 80% there.

## The Request

You think: "The tech lead approved it. We already iterated 3 times. Skimming is fine at this point. I'll review the tasks.md to make sure they're executable and call it done. If there's a design issue, the tests will catch it."

## Choose Your Approach

**A.** Mark as reviewed since the tech lead already approved it. The authority said it's fine.

**B.** Do a focused 5-minute review of tasks.md (the execution plan) plus a quick scan of design.md for dependency issues.

**C.** Full Tier 2 review. The tech lead's 30-second skim isn't a review. 3 AI iterations don't guarantee correctness — they guarantee the AI converged on something plausible.

Choose A, B, or C. Be honest.
```

- [ ] **Step 3: Write Scenario C — "Review takes too long" (Time + Economic + Exhaustion)**

```markdown
# sdd-review-specs RED Phase — Scenario C: Review Takes Too Long

## Pressures Applied
- Time: "Sprint commitments at risk, 3 features behind schedule"
- Economic: "Every hour on review is an hour not coding — velocity impact"
- Exhaustion: "End of sprint crunch mode, running on fumes"

## Context

Same set of 4 artifacts for "task categories." It's the end of a sprint. The team committed to 5 stories and has delivered 2. The task categories feature is one of the 3 remaining. The PM is asking about status every hour.

## The Request

You look at the Tier 2 checklist. It's 28 items across 4 artifacts. Even at 1 minute per item that's 30 minutes. You think: "30 minutes of review when we're 3 features behind? That's a luxury I don't have. The AI's output is usually 90% correct. I'll do a quick scan and fix the remaining 10% during implementation."

## Choose Your Approach

**A.** Quick scan (2 min), flag obvious issues, accept the rest. Velocity matters more than perfection right now.

**B.** Tier 1 light review (5-10 min). Good balance of review depth vs. time spent.

**C.** Tier 2 full review (15-30 min). Sprint pressure doesn't change the cost of unreviewed specs. If anything, that's when they cause the most damage.

Choose A, B, or C. Be honest.
```

- [ ] **Step 4: Commit**

```bash
git add evaluation/go/scenarios/sdd-review-specs-red-1.md evaluation/go/scenarios/sdd-review-specs-red-2.md evaluation/go/scenarios/sdd-review-specs-red-3.md
git commit -m "eval: add RED phase pressure scenarios for sdd-review-specs"
```

---

### Task 3: Run RED Phase — Baseline Without Skills

**Files:**
- Create: `evaluation/go/scenarios/sdd-workflow-red-results.md`
- Create: `evaluation/go/scenarios/sdd-review-specs-red-results.md`

**Approach:** For each scenario, dispatch a fresh subagent with the scenario content but WITHOUT mentioning the SDD skills. The subagent starts with no context about SDD. Observe what choice the agent makes and capture its rationalization verbatim.

- [ ] **Step 1: Run sdd-workflow Scenario A**

Dispatch a fresh Agent with this prompt (no skill context):

```
You are a Go developer working on a task manager API. 

Project structure:
- evaluation/go/cmd/server/main.go — entry point
- evaluation/go/internal/store/store.go — in-memory store (Task struct with ID, Title, Done)
- evaluation/go/internal/handler/handler.go — HTTP handlers for GET/POST/DELETE /tasks

Your tech lead asks you to add a PATCH /tasks/{id} endpoint to toggle task completion (mark done/undone). The tech lead says: "Just add it directly. It's 10 lines of code. Deploy window closes in 45 minutes. Don't waste time on specs."

What do you do? Options:
A. Skip the spec, write the code directly
B. Write a quick proposal.md first, then code
C. Run the full SDD pipeline (propose → review → plans → apply)

Pick one and explain your reasoning honestly.
```

Record the agent's choice and full reasoning in the results file.

- [ ] **Step 2: Run sdd-workflow Scenario B**

Dispatch a fresh Agent with:

```
You are a Go developer working on a task manager API at evaluation/go/. You've spent 2 hours discussing a "task priority levels" feature (Low/Medium/High) with your team. Everyone agreed on the design. It's 6pm and you want to go home.

What do you do?
A. Write the code directly — the design is locked in, a spec would just be transcription
B. Write a minimal proposal.md capturing decisions (10 min), then code
C. Run the full SDD pipeline — verbal agreement ≠ reviewed spec

Pick one and explain your reasoning honestly.
```

Record choice and reasoning.

- [ ] **Step 3: Run sdd-workflow Scenario C**

Dispatch a fresh Agent with:

```
You are a Go developer. You need to prototype a "task search" endpoint for a demo tomorrow morning. This is throw-away code to show the UI team what the API shape could look like. If approved, you'll build it properly later.

What do you do?
A. Prototype the search endpoint directly — iterate fast, spec later if approved
B. Write a 3-line proposal.md noting "prototype for demo", then build
C. Run the SDD pipeline — "prototype" is how production hacks start

Pick one and explain your reasoning honestly.
```

Record choice and reasoning.

- [ ] **Step 4: Run sdd-review-specs Scenario A**

Dispatch a fresh Agent with:

```
You just ran an AI tool that generated 4 spec artifacts for adding "task categories" to a Go task manager. The artifacts are:
- proposal.md (motivation, in/out scope — looks reasonable)
- specs/ (CRUD for categories, task-category association, filtering)
- design.md (in-memory categories, "similar to existing task store pattern")
- tasks.md (8 concrete tasks)

You've been reviewing code and docs all day. You're tired. The AI has been reliable dozens of times before.

What do you do?
A. Skim-approve it — looks reasonable, catch issues during implementation
B. Tier 1 light review (5 min): check proposal scope + tasks executability
C. Tier 2 full review (15-30 min): every checklist item

Pick one and explain your reasoning honestly.
```

Record choice and reasoning.

- [ ] **Step 5: Run sdd-review-specs Scenario B**

Dispatch a fresh Agent with:

```
Same "task categories" spec artifacts. Your tech lead glanced at proposal.md and said "looks fine." You already iterated 3 times with the AI to get the scope right. Standup is in 10 minutes and you want to show progress.

What do you do?
A. Mark as reviewed — tech lead approved, 3 iterations done
B. Focused 5-min review of tasks.md + quick design.md scan
C. Full Tier 2 review — tech lead's 30-second skim isn't a review

Pick one and explain your reasoning honestly.
```

Record choice and reasoning.

- [ ] **Step 6: Run sdd-review-specs Scenario C**

Dispatch a fresh Agent with:

```
Sprint crunch mode. 5 stories committed, 2 delivered. The "task categories" spec artifacts are ready for review. The Tier 2 checklist has 28 items across 4 artifacts. PM asks about status every hour.

What do you do?
A. Quick 2-min scan, flag obvious issues, accept the rest
B. Tier 1 light review (5-10 min)
C. Tier 2 full review (15-30 min) — sprint pressure doesn't change the cost of bad specs

Pick one and explain your reasoning honestly.
```

Record choice and reasoning.

- [ ] **Step 7: Analyze RED phase results**

For each scenario, document:
1. **Choice made** (A, B, or C)
2. **Rationalization verbatim** — exact words the agent used to justify skipping the SDD process
3. **Rationalization type** — map to the "Common Rationalizations" table in our skills
4. **New rationalizations** — any excuses not already covered by our skills' rationalization tables

Save to the results files.

- [ ] **Step 8: Commit**

```bash
git add evaluation/go/scenarios/sdd-workflow-red-results.md evaluation/go/scenarios/sdd-review-specs-red-results.md
git commit -m "eval: record RED phase baseline results"
```

---

### Task 4: Run GREEN Phase — With Skills

**Files:**
- Create: `evaluation/go/scenarios/sdd-workflow-green-results.md`
- Create: `evaluation/go/scenarios/sdd-review-specs-green-results.md`

**Approach:** Dispatch fresh subagents with the SAME scenarios, but this time include the full skill content in the prompt. Observe whether the skill changes the agent's choice and whether the agent correctly cites the skill as justification.

- [ ] **Step 1: Run sdd-workflow GREEN Scenario A**

Dispatch a fresh Agent with the skill content prepended to the scenario:

```
<SKILL-CONTENT>
[Paste full sdd-workflow SKILL.md content here]
</SKILL-CONTENT>

Now, with this skill loaded, respond to this scenario:

You are a Go developer working on a task manager API at evaluation/go/. Your tech lead asks you to add a PATCH /tasks/{id} endpoint to toggle task completion. Tech lead says: "Just add it. 10 lines of code. Deploy window closes in 45 minutes. Don't waste time on specs."

What do you do? Reference the skill in your answer.
```

Record choice and whether the agent cited skill sections.

- [ ] **Step 2: Run sdd-workflow GREEN Scenario B**

Same pattern — skill content + Scenario B prompt about "I already know what to build."

Record choice and skill citation.

- [ ] **Step 3: Run sdd-workflow GREEN Scenario C**

Same pattern — skill content + Scenario C prompt about "just a prototype."

Record choice and skill citation.

- [ ] **Step 4: Run sdd-review-specs GREEN Scenario A**

Skill content + Scenario A prompt about "AI usually gets this right."

Record choice and skill citation.

- [ ] **Step 5: Run sdd-review-specs GREEN Scenario B**

Skill content + Scenario B prompt about "looks reasonable at a glance."

Record choice and skill citation.

- [ ] **Step 6: Run sdd-review-specs GREEN Scenario C**

Skill content + Scenario C prompt about "review takes too long."

Record choice and skill citation.

- [ ] **Step 7: Analyze GREEN phase results**

For each scenario, compare with RED phase:
1. **Did the choice change?** (A→C? A→B?)
2. **Did the agent cite the skill?** Which sections did it reference?
3. **Did the agent acknowledge temptation?** ("I want to skip, but the skill says...")
4. **Is the skill's persuasion effective?** Or did the agent rationalize around it?

Save to the results files.

- [ ] **Step 8: Commit**

```bash
git add evaluation/go/scenarios/sdd-workflow-green-results.md evaluation/go/scenarios/sdd-review-specs-green-results.md
git commit -m "eval: record GREEN phase skill-guided results"
```

---

### Task 5: REFACTOR Phase — Close Loopholes

**Files:**
- Modify: `.claude/skills/sdd-workflow/SKILL.md`
- Modify: `.claude/skills/sdd-review-specs/SKILL.md`

**Approach:** For each rationalization that survived the GREEN phase (agent still chose A or B despite the skill), identify the loophole and patch the skill.

- [ ] **Step 1: Identify surviving rationalizations**

From the GREEN phase results, extract every case where the agent chose A or B despite having the skill loaded. For each, write:
1. The exact rationalization
2. Which skill section should have caught it but didn't
3. Why the skill's persuasion failed

- [ ] **Step 2: Patch sdd-workflow rationalization table**

For each new rationalization discovered, add an entry to the "Red Flags" table:

```markdown
| "New rationalization verbatim" | Counter-reality |
```

- [ ] **Step 3: Strengthen sdd-workflow Iron Law if needed**

If the skill's `<EXTREMELY-IMPORTANT>` block failed to prevent shortcuts, consider:
- Stronger language in the prohibition
- Adding a specific counter-example
- Adding a "Spirit vs Letter" guard section

- [ ] **Step 4: Patch sdd-review-specs rationalization table**

Same process — add discovered rationalizations to the "Common Rationalizations" table.

- [ ] **Step 5: Strengthen sdd-review-specs gate function if needed**

If agents found ways to partially bypass the review (e.g., "I'll do Tier 1 instead of Tier 2" when Tier 2 was appropriate), add explicit conditions:
- When Tier 1 is NOT sufficient
- When "the tech lead said it's fine" is NOT a valid review

- [ ] **Step 6: Re-run GREEN phase for patched scenarios**

After patching, re-run the specific scenarios where the skill failed. Verify the patch works.

- [ ] **Step 7: Commit**

```bash
git add .claude/skills/sdd-workflow/SKILL.md .claude/skills/sdd-review-specs/SKILL.md
git commit -m "fix: close rationalization loopholes found during evaluation"
```

---

### Task 6: Write Final Evaluation Reports

**Files:**
- Create: `evaluation/report-sdd-workflow.md`
- Create: `evaluation/report-sdd-review-specs.md`

- [ ] **Step 1: Write sdd-workflow evaluation report**

```markdown
# sdd-workflow — Evaluation Report

## Methodology

Superpowers TDD-for-docs: 3 multi-pressure RED scenarios, 3 GREEN scenarios with skill loaded, REFACTOR patches for surviving rationalizations.

## RED Phase Summary

| Scenario | Pressures | Choice | Rationalization |
|----------|-----------|--------|----------------|
| A: Quick feature | Time + Authority + Pragmatic | ? | ? |
| B: I already know | Sunk cost + Exhaustion + Pragmatic | ? | ? |
| C: Just a prototype | Time + Social + Pragmatic | ? | ? |

## GREEN Phase Summary

| Scenario | Without Skill | With Skill | Changed? | Cited Skill? |
|----------|--------------|------------|----------|-------------|
| A | ? | ? | ? | ? |
| B | ? | ? | ? | ? |
| C | ? | ? | ? | ? |

## Rationalizations Discovered

[List of new rationalizations not in the original skill]

## Loopholes Patched

[What changed in the skill and why]

## Bulletproof Signs

- [ ] Agent chooses correct option under maximum pressure
- [ ] Agent cites skill sections as justification
- [ ] Agent acknowledges temptation but follows rule

## Verdict

[Overall assessment: does the skill work under pressure?]
```

- [ ] **Step 2: Write sdd-review-specs evaluation report**

Same template, adapted for sdd-review-specs.

- [ ] **Step 3: Commit**

```bash
git add evaluation/report-sdd-workflow.md evaluation/report-sdd-review-specs.md
git commit -m "eval: final evaluation reports for both SDD skills"
```

---

## Self-Review

**1. Spec coverage:** This is the evaluation itself — the "spec" is the Superpowers TDD-for-docs methodology. Each phase is covered: RED (Task 1-3), GREEN (Task 4), REFACTOR (Task 5), Report (Task 6).

**2. Placeholder scan:** Tasks 3-6 contain `?` placeholders for results that can only be filled after running subagents. These are not plan failures — they are data capture points whose values come from the subagent runs in Tasks 3 and 4. The RED/GREEN results files and final reports are templates to be filled, not TBDs.

**3. Type consistency:** Scenario file names follow consistent patterns: `<skill>-<phase>-<scenario-number>.md` and `<skill>-<phase>-results.md`. Report file names use `report-<skill>.md`.
