# Codex Session Prompts

Use these prompts for terminal Codex work in this repository.

Terminal Codex prompts must not include Web AI read-only connector instructions unless the owner explicitly requests a collaboration packet.

Before sending any Codex prompt, check:

- prompt target is Terminal Codex only;
- template source is `docs/templates/0121_codex_session_prompts.md`;
- no Web AI mutation or read-only connector language is included;
- proof commands are grouped in one final proof section;
- no git commands are included unless the owner explicitly asks.

## Start A Codex Session

```text
TARGET AGENT: Terminal Codex
TEMPLATE SOURCE: docs/templates/0121_codex_session_prompts.md

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
- Do not include Web AI read-only connector instructions unless the owner explicitly requests a collaboration packet.
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

SELF-CHECK
- Prompt target is Terminal Codex only.
- Template source is docs/templates/0121_codex_session_prompts.md.
- No Web AI mutation/read-only connector language is included.
- Proof commands are grouped in one final proof section.
- No git commands are included unless owner explicitly asks.

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
TARGET AGENT: Terminal Codex
TEMPLATE SOURCE: docs/templates/0121_codex_session_prompts.md

CONTEXT
Repository: /home/asyraf/Code/go/pos-go
Active scope: REPLACE_WITH_SCOPE
Blueprint: docs/blueprints/REPLACE_WITH_BLUEPRINT.md

FILES TO READ FIRST
docs/README.md
docs/AGENTS.md
docs/0001_index.md
docs/0002_decision_policy.md
docs/0003_session_start_protocol.md
docs/blueprints/REPLACE_WITH_BLUEPRINT.md

TASK
Implement only the next active step from the blueprint.

ACTIVE STEP
REPLACE_WITH_ONE_STEP

ALLOWED FILES
REPLACE_WITH_ALLOWED_FILES

FORBIDDEN FILES
REPLACE_WITH_FORBIDDEN_FILES

RULES
- Do not include Web AI read-only connector instructions unless the owner explicitly requests a collaboration packet.
- Use one active step.
- Mark GAP instead of guessing missing repo state.
- Stop if the blueprint does not contain enough information.
- If blocked by missing data, ask for the smallest specific source batch.
- If blocked by owner decision, ask with 2-3 options and tradeoffs.
- Do not use git commands unless the owner explicitly asks.
- Update handoff before finishing durable work.
- Report estimated progress percentage.
- Report estimated context-window status.

PROOF REQUIRED
REPLACE_WITH_FOCUSED_COMMAND
make verify

EXPECTED OUTPUT
- FACT
- CHANGED FILES
- DECISIONS
- PROOF RESULT
- GAP
- NEXT VALID STEP
- PROGRESS
- CONTEXT STATUS

SELF-CHECK
- Prompt target is Terminal Codex only.
- Template source is docs/templates/0121_codex_session_prompts.md.
- No Web AI mutation/read-only connector language is included.
- Proof commands are grouped in one final proof section.
- No git commands are included unless owner explicitly asks.
```

## Ask Codex To Review Local Changes

```text
TARGET AGENT: Terminal Codex
TEMPLATE SOURCE: docs/templates/0121_codex_session_prompts.md

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
- Do not include Web AI read-only connector instructions unless the owner explicitly requests a collaboration packet.
- Findings first, ordered by severity.
- Use file and line references.
- If there are no findings, say that clearly.
- Do not rewrite files unless I explicitly ask.
```

## Close A Codex Session

```text
TARGET AGENT: Terminal Codex
TEMPLATE SOURCE: docs/templates/0121_codex_session_prompts.md

Create or update a handoff for this session.

HANDOFF TARGET
docs/handoffs/REPLACE_WITH_DATE_SCOPE.md

INCLUDE
- date;
- active scope;
- target agent for any recommended next session;
- template source for any recommended next-session prompt;
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
