---
name: sdd-workflow
description: Use when starting any development task beyond single-line fixes — classifies the request, detects current project phase, and routes to the right specification or execution tool
---

<SUBAGENT-STOP>
If you were dispatched as a subagent to execute a specific task, skip this skill.
</SUBAGENT-STOP>

<CRITICAL-OVERRIDE>
**Two superpowers skills have defaults that conflict with sdd-workflow. Both are OVERRIDDEN:**

1. **Brainstorming** says its terminal state is `writing-plans`.
   - OVERRIDDEN: After brainstorming → invoke `/opsx:propose "<name>"`, NOT `writing-plans`.
   - The pipeline is: `brainstorming → /opsx:propose → review → /opsx:verify → writing-plans`.

2. **Writing-plans** says its output path is `docs/superpowers/plans/YYYY-MM-DD-<name>.md`.
   - OVERRIDDEN: Output MUST go to `openspec/changes/<name>/plan.md`.
   - `docs/superpowers/plans/` is a legacy path — do NOT use it.

If you invoke `writing-plans` immediately after brainstorming, you skip Steps 2-4 (the entire OpenSpec specification phase). If you write `plan.md` to `docs/superpowers/plans/`, it lives outside the OpenSpec traceability system.
</CRITICAL-OVERRIDE>

<EXTREMELY-IMPORTANT>
Spec-driven development means specs live in the file system, not in chat history. OpenSpec manages specification artifacts. Superpowers enforces execution discipline. This skill routes between them.

IF A SPEC EXISTS, YOU MUST READ IT BEFORE WRITING CODE. IF NO SPEC EXISTS FOR BEHAVIOR CHANGE, YOU MUST CREATE ONE FIRST.

This is not negotiable. This is not optional. You cannot rationalize your way out of this.
</EXTREMELY-IMPORTANT>

## Instruction Priority

1. **User's explicit instructions** (CLAUDE.md, AGENTS.md, direct requests) — highest priority
2. **OpenSpec artifacts** (proposal.md, specs/, design.md, tasks.md) — the authoritative spec baseline
3. **SDD workflow skills** — route and enforce process
4. **Default system prompt** — lowest priority

If the user says "skip the spec, just write code," follow the user's instructions. The user is in control.

# SDD Workflow — Spec-Driven Development Router

## The Rule

**Before any code, human and AI agree on what to build.** Specifications are files in `openspec/`. Every behavior change is traceable from proposal through archive. Run `openspec init` if `openspec/` doesn't exist.

**Announce at start:** "I'm using the sdd-workflow skill to route this development task."

## The 10-Step Pipeline

OpenSpec provides the specification skeleton (what). Superpowers enforces execution discipline (how). They connect in sequence with no overlap:

```
0. [Optional] superpowers:brainstorming  — Exploratory design for greenfield/fuzzy requirements.
                                           Output: design doc → feeds into Step 2.
                                           ⛔ After approval → go to Step 2, NOT writing-plans.

1. [User request]
     ↓
2. /opsx:propose <name>            — OpenSpec: Create proposal. Generate all 4 artifacts.
     ↓                               proposal.md + specs/ + design.md + tasks.md
3. [Manual review + iterate]       — OpenSpec: Refine specs. Review and revise each artifact.
     ↓                               Optional: /opsx:continue (step-by-step), /opsx:ff (fast-forward).
4. /opsx:verify                    — OpenSpec: Verify specs. 3-dimension validation gate.
     ↓
5. superpowers:writing-plans       — Superpowers: Generate implementation plan.
     ↓                               Output: openspec/changes/<name>/plan.md
6. /opsx:apply +                   — Superpowers: TDD execution. apply = scheduler, TDD = executor.
   @test-driven-development          RED → GREEN → REFACTOR per task.
     ↓
7. @requesting-code-review         — Superpowers: Code review.
     ↓
8. @verification-before-           — Superpowers: Pre-completion verification. Fresh test evidence.
   completion
     ↓
9. /opsx:archive <name>            — OpenSpec: Archive change. Delta merge + move to archive/.
     ↓
10. [Delivered]                    — Ship it.
```

## Artifact Ownership

OpenSpec and Superpowers each produce plan-like files — they serve different roles and both belong in `openspec/changes/<name>/`:

| Artifact | Owner | Granularity | Purpose | Example |
|----------|-------|-------------|---------|---------|
| `tasks.md` | OpenSpec (`/opsx:propose`) | Coarse checkbox items | WHAT to implement | `- [ ] Implement Store interface` |
| `plan.md` | Superpowers (`writing-plans`) | 2-5min subtasks | HOW to implement | `1. Define Store interface in store/store.go (2 min)` |

**Rule:** `plan.md` refines `tasks.md` — it does NOT replace it. Both coexist in the same change directory. `writing-plans` reads `tasks.md` as input and outputs `plan.md` with detailed steps, file paths, and test names.

```
openspec/changes/<name>/
├── proposal.md    ← OpenSpec: why + scope boundary
├── specs/         ← OpenSpec: behavior delta specs
├── design.md      ← OpenSpec: technical decisions
├── tasks.md       ← OpenSpec: coarse implementation checklist (WHAT)
└── plan.md        ← Superpowers: refined subtasks (HOW)
```

`docs/superpowers/specs/` and `docs/superpowers/plans/` are legacy brainstorming output paths — they are NOT used by this workflow. All artifacts live under `openspec/changes/<name>/`.

## OpenSpec Command Reference

| Command | Description | When to use |
|---------|-------------|-------------|
| `/opsx:propose` | Generate complete change artifacts in one step | Requirements are clear |
| `/opsx:explore` | Investigate without creating files | Fuzzy requirements, tech evaluation, approach comparison |
| `/opsx:apply` | Implement tasks from tasks.md item by item | Implementation phase |
| `/opsx:archive` | Archive and merge specs | Feature complete |
| `/opsx:new` | Create change skeleton only | Want fine-grained control over artifact creation |
| `/opsx:continue` | Generate the next artifact | Step-by-step review, confirm each artifact |
| `/opsx:ff` | Fast-forward: generate all remaining artifacts | Direction confirmed, accelerate |
| `/opsx:verify` | 3-dimension validation of implementation | Pre-archive quality check |
| `/opsx:sync` | Sync specs without archiving | Parallel changes need reference |
| `/opsx:bulk-archive` | Batch archive multiple changes | Multi-feature unified wrap-up |

## Request Classification

When the user brings a development request, classify FIRST. Then route.

### Boundedness Check — BEFORE routing to Step 2

A task is **NOT "clearly bounded"** (and therefore MUST route through Step 0: exploration or brainstorming) if ANY of these are true:

| Signal | Example | Route to |
|--------|---------|----------|
| Introduces concepts NOT in the current data model | "add users", "add sharing", "add permissions" — and the codebase has no User/Share/Permission struct | `superpowers:brainstorming` |
| Has multiple valid interpretations with different architectures | "add collaboration" could mean real-time sync, async assignment, or shared views | `superpowers:brainstorming` |
| Uses hedging or vague language | "somehow", "或者", "something like", "加点协作能力" | `superpowers:brainstorming` |
| Requires comparing 2+ approaches with significant trade-offs | "should we use WebSocket or polling?" | `superpowers:brainstorming` |
| You don't know which files would change without reading code first | Unfamiliar codebase or new feature area | `/opsx:explore` → then re-run Boundedness Check → if still fuzzy → `superpowers:brainstorming` |

A task IS "clearly bounded" (can skip to Step 2) ONLY when ALL of these are true:
- The data model is already defined (structs/tables exist)
- There is exactly one obvious implementation approach
- The request uses specific, concrete language ("add a DELETE endpoint", "add a `due_date` field to Task")
- You can list the files that will change without reading any code

**Signal priority:** Brainstorming signals beat exploration signals. When a task matches BOTH a brainstorming signal AND the exploration signal, the flow is: `/opsx:explore` (read code, build context) → re-run Boundedness Check → `superpowers:brainstorming` (generate and compare approaches). Never skip the brainstorming step when ANY brainstorming signal is triggered.

**Default rule: if you're not sure, it's not clearly bounded. Route to exploration or brainstorming.**

### CRITICAL: After `/opsx:explore` — Do NOT present options

`/opsx:explore` builds context. It does NOT authorize you to decide what to implement. After it completes:

- **Do NOT** present a numbered list of features and ask the user to pick
- **Do NOT** ask "你希望补充哪些功能？" or "Which features do you want?"
- **Do NOT** merge exploration + decision into one step

**You MUST instead:**
1. Re-run the Boundedness Check against the task
2. If ANY "not bounded" signal still applies → invoke `superpowers:brainstorming` to generate and compare approaches
3. Only skip brainstorming if the exploration revealed exactly ONE obvious gap (e.g., "this CRUD API is missing a DELETE handler")

**Wrong:** Explore → "Here are 4 options, pick one" → implement
**Right:** Explore → Boundedness Check → Brainstorming → `/opsx:propose` → review → implement

```dot
digraph sdd_routing {
    "User request received" [shape=doublecircle];
    "Is it a one-line fix?\n(typo, log line, comment)" [shape=diamond];
    "Make change,\nverify directly" [shape=box style=filled fillcolor="#ccffcc"];
    "Is it a bug with\nunclear cause?" [shape=diamond];
    "superpowers:systematic-debugging\n(root cause first)" [shape=box style=filled fillcolor="#ffcccc"];
    "Are there OpenSpec\nartifacts already?" [shape=diamond];
    "Read existing artifacts,\npick up where left off" [shape=box style=filled fillcolor="#ccccff"];
    "Run Boundedness Check.\nIs it clearly bounded?" [shape=diamond];
    "Unfamiliar codebase\nor uncertain approach?" [shape=diamond];
    "/opsx:propose\n(generate 4 artifacts)" [shape=box style=filled fillcolor="#ffffcc"];
    "/opsx:explore\n(build context first)" [shape=box style=filled fillcolor="#ffffcc"];
    "superpowers:brainstorming\n(Socratic design)" [shape=box style=filled fillcolor="#ffffcc"];
    "Route to 10-Step\nPipeline Step 2" [shape=doublecircle];

    "User request received" -> "Is it a one-line fix?\n(typo, log line, comment)";
    "Is it a one-line fix?\n(typo, log line, comment)" -> "Make change,\nverify directly" [label="yes"];
    "Is it a one-line fix?\n(typo, log line, comment)" -> "Is it a bug with\nunclear cause?" [label="no"];
    "Is it a bug with\nunclear cause?" -> "superpowers:systematic-debugging\n(root cause first)" [label="yes"];
    "Is it a bug with\nunclear cause?" -> "Are there OpenSpec\nartifacts already?" [label="no"];
    "Are there OpenSpec\nartifacts already?" -> "Read existing artifacts,\npick up where left off" [label="yes"];
    "Are there OpenSpec\nartifacts already?" -> "Run Boundedness Check.\nIs it clearly bounded?" [label="no"];
    "Run Boundedness Check.\nIs it clearly bounded?" -> "Unfamiliar codebase\nor uncertain approach?" [label="no — see\nBoundedness Check"];
    "Run Boundedness Check.\nIs it clearly bounded?" -> "Route to 10-Step\nPipeline Step 2" [label="yes — meets ALL\nclear-boundary criteria"];
    "Unfamiliar codebase\nor uncertain approach?" -> "/opsx:explore\n(build context first)" [label="need to read\ncode first"];
    "Unfamiliar codebase\nor uncertain approach?" -> "superpowers:brainstorming\n(Socratic design)" [label="greenfield\nor approach\ncomparison"];
    "/opsx:explore\n(build context first)" -> "Run Boundedness Check.\nIs it clearly bounded?";
    "superpowers:brainstorming\n(Socratic design)" -> "Run Boundedness Check.\nIs it clearly bounded?";
}
```

## Phase Detection

Check the file system to determine where you are in the workflow:

| What exists | Phase | Next action |
|------------|-------|-------------|
| No `openspec/` directory | Uninitialized | Run `openspec init` first |
| `openspec/` exists, no change dir | Ready for proposal | Route to Step 2: `/opsx:propose <name>` or Step 0: exploration |
| `openspec/changes/<name>/` with 4 artifacts, unreviewed | Specs need review | Steps 3-4: Manual review → `/opsx:verify` |
| `openspec/changes/<name>/` with reviewed artifacts | Ready for execution | Step 5: `superpowers:writing-plans` |
| `tasks.md` has unchecked items | In progress | Step 6: `/opsx:apply` + `@test-driven-development` |
| All tasks checked, not archived | Ready for delivery | Steps 7-8: review → verify → Step 9: `/opsx:archive` |

## Transition Rules

### Step 0: Explore → Brainstorming Gate

When the Boundedness Check routes to `/opsx:explore` (because you don't know which files would change), `/opsx:explore` is read-only reconnaissance — it builds context, NOT decisions. After it completes:

1. **Re-run the Boundedness Check.** The task is still "not clearly bounded" unless the codebase exploration revealed a single obvious gap.
2. **If ANY "not bounded" signal still applies** → route to `superpowers:brainstorming`. Do NOT present options to the user. Do NOT ask the user what to implement. Brainstorming is where options are generated and compared.
3. **If the task is now clearly bounded** (rare after a fuzzy request) → proceed to Step 2: `/opsx:propose`.

**Why:** `/opsx:explore` answers "what exists." `superpowers:brainstorming` answers "what should we build." Mixing them (explore → offer choices → implement) skips design entirely. The DOT graph shows the loop: `explore → Boundedness Check` — this must execute, not be skipped.

### Step 0 → Step 2: The Critical Handoff

Brainstorming (Step 0) is optional. When requirements are already clear, skip to Step 2.

**The brainstorming skill says its terminal state is `writing-plans`. THIS IS OVERRIDDEN.** When brainstorming is invoked through sdd-workflow, the pipeline is: brainstorming → /opsx:propose → review → /opsx:verify → writing-plans.

When brainstorming completes and the user approves the design:

1. **DO NOT** invoke `writing-plans` — this bypasses OpenSpec Steps 2-4
2. **DO NOT** write code — the spec isn't locked yet
3. **DO** invoke `/opsx:propose "<name>"` — feed the approved brainstorming design as context
4. **DO** verify `openspec/changes/<name>/` contains: proposal.md, specs/, design.md, tasks.md
5. Only then proceed to Step 3.

**Why:** Brainstorming produces an exploratory design (Phase 1 — Superpowers). OpenSpec locks it into auditable, mergeable artifacts (Phase 2 — OpenSpec). `docs/superpowers/specs/` is transient; `openspec/changes/<name>/` is permanent and traceable. Skipping Step 2 means specs can't be verified, archived, or traced.

### Steps 2-10: Linear Execution

```
Step 2: /opsx:propose <name>     → If Step 0 was done, feed its output as context.
                                    Confirm artifacts before Step 3.

Step 3: Manual review + iterate   → Review proposal/specs/design/tasks item by item.
                                    Optional: /opsx:continue (step) | /opsx:ff (fast-forward).
                                    Standard: every in-scope item has a task checkbox.

Step 4: /opsx:verify              → 3-dimension validation (complete/correct/consistent).
                                    Pass before entering execution phase.

Step 5: superpowers:writing-plans → MUST save to openspec/changes/<name>/plan.md
                                    (NOT docs/superpowers/plans/).
                                    Output: 2-5 minute granular subtasks.

Step 6: /opsx:apply +             → apply = scheduler, TDD = executor.
  @test-driven-development          RED → GREEN → REFACTOR per task.
                                    On error: @systematic-debugging → return to apply.
                                    All tasks complete → Step 7.

Step 7: @requesting-code-review   → Dispatch code-reviewer. Fix Critical/Important issues.

Step 8: @verification-before-     → Fresh go test ./... / pytest / etc.
  completion                        Evidence required BEFORE claiming completion.

Step 9: /opsx:archive <name>      → Delta merge into openspec/specs/.
                                    Change moved to openspec/changes/archive/.
                                    project.md updated.

Step 10: Done                     → Ship it.
```

## Tool Selection Matrix

When both OpenSpec and Superpowers offer a tool for the same phase:

| Scenario | Use This | Not That | Why |
|----------|----------|----------|-----|
| Reading existing code | `/opsx:explore` | `@brainstorming` | Explore reads code; brainstorming generates ideas |
| Defining new feature | `@brainstorming` | `/opsx:explore` | Brainstorming compares approaches |
| Generating spec artifacts | `/opsx:propose` | `@writing-plans` | Propose creates 4 artifacts; writing-plans refines |
| Refining task granularity | `@writing-plans` | Manual only | Writing-plans converts to 2-5min units |
| Executing tasks | `/opsx:apply` + `@test-driven-development` | Either alone | Apply schedules; TDD executes |
| Debugging failures | `@systematic-debugging` | Direct fixes | Root cause investigation first |
| Code review | `@requesting-code-review` + `@receiving-code-review` | "Looks good to me" | Structured independent review |
| Claiming completion | `@verification-before-completion` | "Should work now" | Fresh evidence required |
| Archiving work | `/opsx:archive` | Manual file moves | Archive does delta merge + timestamp |

## Red Flags

These thoughts mean STOP — you're rationalizing skipping the SDD process:

| Thought | Reality |
|---------|---------|
| "This is simple, I don't need a spec" | Simple changes cause complex bugs. A 5-line proposal.md saves hours. |
| "I'll write the spec after the code" | Specs-after describe what you built, not what's needed. |
| "The spec is in the conversation history" | Conversation history evaporates. Files persist. Write it down. |
| "I already know what to build" | Knowing ≠ having it reviewed. Specs are the agreement. |
| "Specs slow me down" | Rework from misaligned expectations is slower. |
| "This is just a prototype" | Prototypes become production. Spec now saves pain later. |
| "I'll just explore the codebase first" | Use `/opsx:explore` — structured, not aimless browsing. |
| "I remember how this codebase works" | Code evolves. Your memory is stale. Read the specs. |
| "This task is clearly bounded, skip brainstorming" | ⛔ STOP. Run the Boundedness Check. Does the task introduce concepts NOT in the current data model? Does it have multiple valid interpretations? If yes → brainstorming. "Add collaboration" on a single-user app is NOT clearly bounded. |
| "I already explored the codebase, I can just list options" | ⛔ STOP. `/opsx:explore` answers "what exists," not "what to build." After explore, re-run Boundedness Check. If the task is still fuzzy → `superpowers:brainstorming`. Presenting a menu of options is NOT a substitute for Socratic design. |
| "The user said '完善' or 'improve' — that's clear enough" | Vague verbs imply the user trusts you to figure out WHAT to improve. That's exactly what brainstorming is for. Explore the codebase → brainstorm what should change → THEN propose. |
| "Brainstorming says go to writing-plans" | ⛔ OVERRIDDEN. sdd-workflow pipeline: brainstorming → `/opsx:propose` → review → verify → THEN writing-plans. |
| "I'll write the design doc — that's the spec" | `docs/superpowers/specs/` is transient. `/opsx:propose` creates permanent `openspec/changes/<name>/` artifacts. |
| "The brainstorming design IS the OpenSpec design" | No. Brainstorming output is INPUT to `/opsx:propose`. It must be translated into the 4 OpenSpec artifacts. |

**All of these mean: follow the SDD process. No shortcuts.**

## Skill Priority

When multiple tools could apply to a development task:

1. **Classification first** — Use the decision tree. One-line fix? Bug? Behavior change?
2. **Exploration before specification** — `/opsx:explore` to read code. `/opsx:propose` to generate artifacts. Never invert.
3. **Review before execution** — Steps 3-4 gate. Specs must be reviewed before any code.
4. **Plan before implementing** — Step 5: `writing-plans` refines tasks.md. Save to `openspec/changes/<name>/plan.md`.
5. **TDD during execution** — Step 6: `/opsx:apply` + `@test-driven-development`. RED → GREEN → REFACTOR.
6. **Verify before claiming** — Step 8: `@verification-before-completion` with fresh evidence. Then Step 9: `/opsx:archive`.

## Skill Types

**Rigid** — Follow exactly. Don't adapt away the sequence:
- `/opsx:propose`, `/opsx:apply`, `/opsx:archive` — CLI tools with defined behavior
- `@test-driven-development` — RED → GREEN → REFACTOR, no shortcuts
- `@systematic-debugging` — Root cause before fixes
- `@verification-before-completion` — Fresh evidence required
- **`sdd-workflow`** (this skill) — Follow the routing exactly

**Flexible** — Adapt principles to context:
- `@brainstorming` — Socratic design, adapt depth to complexity (but terminal routing is OVERRIDDEN by sdd-workflow)
- `@writing-plans` — Task granularity scales with feature complexity
- `/opsx:explore` — Depth of exploration matches uncertainty level

## Related Skills

- **sdd-review-specs** (included: `reference/sdd-review-specs.md`) — Structured review of OpenSpec 4 artifacts before implementation
- **superpowers:brainstorming** — Socratic design for greenfield features
- **superpowers:writing-plans** — Convert coarse tasks to 2-5min bite-sized units
- **superpowers:test-driven-development** — RED-GREEN-REFACTOR cycle
- **superpowers:systematic-debugging** — Root cause investigation before fixes
- **superpowers:verification-before-completion** — Evidence before completion claims
- **superpowers:requesting-code-review** — Structured code review
- **superpowers:finishing-a-development-branch** — Merge/PR/keep/discard decisions

## User Instructions

Instructions say WHAT, not HOW. "Add X" or "Fix Y" doesn't mean skip workflows. The SDD process is the HOW — it exists to ensure alignment before code, not to slow you down.

If you want to bypass a step (skip review, write code directly, skip the spec), say so explicitly. The user is in control. The skill routes, the user decides.
