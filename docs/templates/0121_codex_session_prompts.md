# Codex Session Prompts

Use these prompts for terminal Codex work in this repository.

## Start A Codex Session

```text
Read the repository instructions before doing any work.

CONTEXT
Repository: /home/asyraf/Code/go/pos-go
Active scope: REPLACE_WITH_SCOPE

FILES TO READ FIRST
docs/README.md
docs/AGENTS.md
docs/0001_index.md
docs/0002_decision_policy.md
docs/0003_session_start_protocol.md
docs/blueprints/REPLACE_WITH_ACTIVE_BLUEPRINT.md

TASK
REPLACE_WITH_TASK

RULES
- Do not implement before reading the active blueprint.
- Use fd for file discovery and rg for text search.
- Use one active step.
- Execute the largest safe slice that still fits that one active step.
- Prefer short progress updates and a compact final report over many tiny question/answer turns.
- Do not stop for discussion unless missing source data, missing proof, or an ADR-level decision blocks implementation.
- If Laravel source data is missing, request the smallest specific folder, file, route, migration, seeder, test, or command output.
- If an ADR or owner decision is needed, ask one concise question with 2-3 viable options, tradeoffs, and the recommended option first when clear.
- Show proof before claiming progress.
- Mark GAP instead of guessing missing repo state.
- Do not touch files outside the active scope.
- If docs workflow rules change, update the impacted README/index/AGENTS/template/audit chain in the same step when feasible.
- Create or update a handoff before context runs low or when a long-running scope ends with durable changes.

EXPECTED OUTPUT
- Fact summary.
- Gap summary.
- One active step.
- Files changed.
- Proof command output.
- Estimated progress percentage for the active scope.
- Estimated context-window status.
- Next valid step.

STYLE
- Keep analysis factual and proof-linked.
- Keep the final answer concise.
- Mention failed commands only when they matter, with the reason and recovery.
- Do not repeat large command output; summarize the relevant result.
```

## Ask Codex To Implement One Blueprint Step

```text
Implement only the next active step from this blueprint:

docs/blueprints/REPLACE_WITH_BLUEPRINT.md

ACTIVE STEP
REPLACE_WITH_ONE_STEP

ALLOWED FILES
REPLACE_WITH_ALLOWED_FILES

FORBIDDEN FILES
REPLACE_WITH_FORBIDDEN_FILES

PROOF REQUIRED
REPLACE_WITH_FOCUSED_COMMAND
make verify

Stop if the blueprint does not contain enough information. Mark the missing item as GAP.

If blocked by missing data, ask for the smallest specific source batch.
If blocked by owner decision, ask with 2-3 options and tradeoffs.
```

## Ask Codex To Review Local Changes

```text
Review the current local changes as a code reviewer.

SCOPE
REPLACE_WITH_SCOPE

FOCUS
- correctness bugs;
- domain invariant regressions;
- API contract regressions;
- database constraint issues;
- missing tests;
- capability/authz bypass risk.

RULES
- Findings first, ordered by severity.
- Use file and line references.
- If there are no findings, say that clearly.
- Do not rewrite files unless I explicitly ask.
```

## Close A Codex Session

```text
Create or update a handoff for this session.

HANDOFF TARGET
docs/handoffs/REPLACE_WITH_DATE_SCOPE.md

INCLUDE
- date;
- active scope;
- files changed;
- proof collected;
- tests or commands run;
- decisions made;
- gaps;
- next valid active step;
- commands the next session should run first;
- estimated scope progress percentage;
- estimated context-window status.

Do not claim unrun tests.
```
