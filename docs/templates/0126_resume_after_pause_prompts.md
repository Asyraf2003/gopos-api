# Resume After Pause Prompts

Use these prompts after a long break, a context reset, or a model switch.

## Restart After A Long Pause

```text
I am returning to this repository after a pause.

REPOSITORY
/home/asyraf/Code/go/pos-go

TASK
Re-orient safely before doing work.

READ FIRST
docs/README.md
docs/AGENTS.md
docs/0001_index.md
docs/0002_decision_policy.md
docs/0003_session_start_protocol.md
docs/blueprints/README.md
docs/handoffs/README.md

THEN
- List active blueprints.
- List recent handoffs.
- Identify the likely active scope.
- Identify open GAP items.
- Identify last reported progress percentage and context-window status if available.
- Recommend exactly one next active step.

RULES
- Do not edit files.
- Do not assume old chat memory is correct.
- Use file contents as proof.
```

## Restart A Specific Scope

```text
Resume this specific scope.

SCOPE
REPLACE_WITH_SCOPE

READ
docs/README.md
docs/AGENTS.md
docs/0001_index.md
docs/0002_decision_policy.md
docs/0003_session_start_protocol.md
docs/blueprints/REPLACE_WITH_BLUEPRINT.md
docs/evidence/REPLACE_WITH_EVIDENCE.md
docs/handoffs/REPLACE_WITH_HANDOFF.md

TASK
- Summarize current facts.
- Summarize gaps.
- State accepted decisions.
- State next active step.
- State proof needed.

Do not implement until I confirm the next step.
```

## Model Switch Prompt

```text
You are taking over from another AI model.

IMPORTANT
Do not rely on unstated memory. Use only the files and text provided here.

ACTIVE SCOPE
REPLACE_WITH_SCOPE

FILES TO READ
REPLACE_WITH_FILES

PREVIOUS SUMMARY
REPLACE_WITH_SUMMARY

TASK
Validate the summary against the files, then continue with one active step.

OUTPUT
- confirmed facts;
- corrected facts;
- gaps;
- next active step;
- proof required;
- estimated progress percentage;
- estimated context-window status.
```

## No Handoff Exists Prompt

```text
No reliable handoff exists.

TASK
Rebuild context from repository docs only.

READ
docs/README.md
docs/AGENTS.md
docs/0001_index.md
docs/blueprints/
docs/evidence/
docs/handoffs/

OUTPUT
- candidate active scopes;
- newest evidence files;
- newest blueprints;
- safest next step;
- files that should not be touched yet.

Do not edit files in this pass.
```
