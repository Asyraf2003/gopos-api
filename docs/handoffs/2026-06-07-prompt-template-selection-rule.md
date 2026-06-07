# Handoff: Prompt Template Selection Rule

## Date

2026-06-07

## Active Scope

Repair AI prompt and handoff templates to prevent mixed Web AI and Codex next-session prompts.

## Current Branch Or Source Snapshot

Local workspace: `/home/asyraf/Code/go/pos-go`

## Files Included

```text
docs/README.md
docs/AGENTS.md
docs/0001_index.md
docs/0002_decision_policy.md
docs/0003_session_start_protocol.md
docs/templates/0120_prompt_authoring_rules.md
docs/templates/0121_codex_session_prompts.md
docs/templates/0122_web_ai_session_prompts.md
docs/workflow/0071_handoff_protocol.md
docs/workflow/0072_transition_progress_ledger_protocol.md
docs/core/0012_step_by_step_execution.md
docs/core/0013_proof_and_progress.md
docs/evidence/0003_laravel_to_go_transition_progress_ledger.md
```

## Files Changed

```text
docs/README.md
docs/AGENTS.md
docs/0001_index.md
docs/templates/0120_prompt_authoring_rules.md
docs/templates/0121_codex_session_prompts.md
docs/templates/0122_web_ai_session_prompts.md
docs/workflow/0071_handoff_protocol.md
docs/evidence/0003_laravel_to_go_transition_progress_ledger.md
docs/handoffs/2026-06-07-prompt-template-selection-rule.md
```

## Files Forbidden To Touch

```text
internal/**
cmd/**
migrations/**
scripts/**
POS domain CRUD
servicecatalog implementation
productcatalog implementation
capability implementation code
production secrets
```

## Blueprint Referenced

No implementation blueprint changed. This is a workflow/docs quality hardening step.

## ADR And Rules Referenced

```text
docs/README.md
docs/AGENTS.md
docs/0001_index.md
docs/0002_decision_policy.md
docs/0003_session_start_protocol.md
docs/templates/0120_prompt_authoring_rules.md
docs/templates/0121_codex_session_prompts.md
docs/templates/0122_web_ai_session_prompts.md
docs/workflow/0071_handoff_protocol.md
docs/workflow/0072_transition_progress_ledger_protocol.md
docs/core/0012_step_by_step_execution.md
docs/core/0013_proof_and_progress.md
```

## Decisions Made

- Every next-session prompt must select exactly one target agent before drafting.
- Terminal Codex prompts must use `docs/templates/0121_codex_session_prompts.md`.
- Web AI prompts must use `docs/templates/0122_web_ai_session_prompts.md`.
- Hybrid Web AI/Codex prompts are forbidden unless the owner explicitly requests a collaboration packet.
- If the owner does not specify a target agent, the assistant must ask one concise clarification question.
- Handoffs that recommend a next AI session must declare the target agent and template source.
- Cascade updates were required because the rule changes how future sessions start, continue, and hand off.

## Proof Collected

```text
rg -n "Prompt Target Selection Rule|TARGET AGENT|TEMPLATE SOURCE|collaboration packet|Web AI or Codex" docs/templates docs/workflow docs/handoffs docs/evidence
```

Result:

```text
docs/templates/0120_prompt_authoring_rules.md includes Prompt Target Selection Rule.
docs/templates/0121_codex_session_prompts.md includes TARGET AGENT and TEMPLATE SOURCE markers for Terminal Codex prompts.
docs/templates/0122_web_ai_session_prompts.md includes TARGET AGENT and TEMPLATE SOURCE markers for Web AI prompts.
docs/workflow/0071_handoff_protocol.md forbids "Web AI or Codex next session" as a single prompt.
docs/handoffs/2026-06-07-prompt-template-selection-rule.md records the collaboration packet decision.
docs/evidence/0003_laravel_to_go_transition_progress_ledger.md records the workflow/docs quality improvement.
```

```text
make verify
```

Result:

```text
[PASS] go vet audit passed
[PASS] format audit passed
[PASS] AI rules audit passed
[PASS] file size audit passed
[PASS] hexagonal import audit passed
[PASS] gosec audit passed
[PASS] aggregate audit passed
```

## Tests Or Commands Run

```text
rg -n "Prompt Target Selection Rule|TARGET AGENT|TEMPLATE SOURCE|collaboration packet|Web AI or Codex" docs/templates docs/workflow docs/handoffs docs/evidence
make verify
```

## Gaps Still Open

- Route-to-capability audit script remains the next capability-control foundation gap.
- POS CRUD remains blocked.

## Next Valid Active Step

Return to `docs/blueprints/0010_capability_control_foundation.md` and add the route-to-capability audit script after this docs-only workflow fix is accepted.

For a next Terminal Codex session:

- target agent: Terminal Codex;
- template source: `docs/templates/0121_codex_session_prompts.md`;
- next active step: route-to-capability audit script;
- files to read: `docs/README.md`, `docs/AGENTS.md`, `docs/0001_index.md`, `docs/0002_decision_policy.md`, `docs/0003_session_start_protocol.md`, `docs/blueprints/0010_capability_control_foundation.md`, and this handoff;
- files may edit: route-audit docs/script files only when explicitly scoped by the owner;
- files must not edit: POS CRUD, catalog implementation, unrelated auth behavior, production secrets;
- proof commands: owner-provided for the next scoped implementation prompt;
- progress percentage: capability-control foundation remains 75%;
- context-window status: enough context for one focused next step.

## Estimated Scope Progress Percentage

Prompt template hardening scope: 100%.

Laravel-to-Go transition: unchanged at 20%.

## Estimated Context-Window Status

Enough context remains for one focused follow-up step.
