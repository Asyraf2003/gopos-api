# Handoff: Web AI Owner Terminal Output Test

## Date

2026-06-07

## Active Scope

Fix Web AI template behavior after a real prompt-output test showed residual Codex-default output.

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
docs/templates/0122_web_ai_session_prompts.md
docs/workflow/0071_handoff_protocol.md
docs/workflow/0072_transition_progress_ledger_protocol.md
docs/evidence/0003_laravel_to_go_transition_progress_ledger.md
docs/handoffs/2026-06-07-ai-execution-channel-boundaries.md
```

## Files Changed

```text
docs/README.md
docs/AGENTS.md
docs/0001_index.md
docs/templates/0120_prompt_authoring_rules.md
docs/templates/0122_web_ai_session_prompts.md
docs/workflow/0071_handoff_protocol.md
docs/evidence/0003_laravel_to_go_transition_progress_ledger.md
docs/handoffs/2026-06-07-web-ai-owner-terminal-output-test.md
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
GitHub branches, commits, pull requests, issues, labels, reviewers, merges, refs, or CI
```

## Blueprint Referenced

No implementation blueprint changed. This is a workflow/docs quality hardening step.

## Rules Referenced

```text
docs/README.md
docs/AGENTS.md
docs/0001_index.md
docs/0002_decision_policy.md
docs/0003_session_start_protocol.md
docs/templates/0120_prompt_authoring_rules.md
docs/templates/0122_web_ai_session_prompts.md
docs/workflow/0071_handoff_protocol.md
docs/workflow/0072_transition_progress_ledger_protocol.md
docs/handoffs/2026-06-07-ai-execution-channel-boundaries.md
```

## Test Result That Triggered This Fix

A real Web AI route-to-capability audit analysis stayed read-only and did not claim local execution, but still emitted `HANDOFF TEXT FOR CODEX`.

That output is now treated as non-compliant for normal Web AI analysis unless the owner explicitly requested Codex.

## Decisions Made

- Web AI Codex-default output is treated as a template-compliance failure.
- Normal Web AI analysis must include `COMMAND PLAN FOR OWNER / LOCAL TERMINAL`.
- Normal Web AI analysis must put local execution proof commands under `PROOF THE OWNER OR TERMINAL AGENT MUST RUN`.
- Normal Web AI analysis must omit `OPTIONAL HANDOFF TEXT FOR CODEX`.
- `OPTIONAL HANDOFF TEXT FOR CODEX` is allowed only when the owner explicitly says the next target is Codex, asks to prepare a Codex handoff, or requests a Codex-specific closeout.
- If `OPTIONAL HANDOFF TEXT FOR CODEX` appears, the answer must quote or name the exact owner instruction that requested Codex.
- If there is no exact owner Codex request, Codex handoff is omitted and the owner/local terminal command plan remains.
- `NEXT` must name exactly one next execution channel.
- Cascade updates were required because the rule changes future session starts, outputs, and handoffs.

## Proof Collected

```text
rg -n "COMMAND PLAN FOR OWNER / LOCAL TERMINAL|OPTIONAL HANDOFF TEXT FOR CODEX|exact owner.*Codex|Codex handoff.*omitted|next execution channel|Web AI Codex-default|owner/local terminal" docs/templates docs/workflow docs/handoffs docs/evidence docs/README.md docs/AGENTS.md docs/0001_index.md
```

Result:

```text
docs/templates/0122_web_ai_session_prompts.md includes owner/local terminal command-plan output, exact owner Codex request checks, Codex handoff omission rules, and next execution channel checks.
docs/templates/0120_prompt_authoring_rules.md includes the Web AI analysis output rule, owner/local terminal default channel, and next execution channel rule.
docs/workflow/0071_handoff_protocol.md forbids default Web AI Codex handoff sections and requires one next execution channel.
docs/README.md, docs/AGENTS.md, and docs/0001_index.md include cascade rules for normal Web AI analysis and owner/local terminal command plans.
docs/evidence/0003_laravel_to_go_transition_progress_ledger.md records this workflow/docs quality improvement.
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
rg -n "COMMAND PLAN FOR OWNER / LOCAL TERMINAL|OPTIONAL HANDOFF TEXT FOR CODEX|exact owner.*Codex|Codex handoff.*omitted|next execution channel|Web AI Codex-default|owner/local terminal" docs/templates docs/workflow docs/handoffs docs/evidence docs/README.md docs/AGENTS.md docs/0001_index.md
make verify
```

## Gaps Still Open

- Route-to-capability audit script remains the next capability-control foundation gap.
- POS CRUD remains blocked.

## Next Valid Active Step

Return to `docs/blueprints/0010_capability_control_foundation.md` and add the route-to-capability audit script after this docs-only workflow fix is accepted.

Recommended next execution channel:

- owner/local terminal or Terminal Codex for implementation, depending on the owner's next prompt;
- Web AI for read-only review after owner pushes;
- explicit collaboration packet only when the owner requests that mode.

## Estimated Scope Progress Percentage

Web AI owner/local terminal output hardening scope: 100%.

Laravel-to-Go transition: unchanged at 20%.

## Estimated Context-Window Status

Enough context remains for this docs-only hardening proof.

## Assessment For Web AI Review

Scores use 0 as best and 5 as worst.

- Web AI Codex-default output risk: 1
- Optional Codex handoff leakage risk: 1
- owner/local terminal command-plan omission risk: 1
- next execution channel ambiguity: 1
- template compliance risk: 1
- docs cascade risk: 1

Web AI should re-check these files and sections after owner pushes:

- `docs/templates/0120_prompt_authoring_rules.md`: `AI Execution Channel Rule`
- `docs/templates/0122_web_ai_session_prompts.md`: permission model, normal expected outputs, optional Codex handoff prompt, cleanup prompt, self-check sections
- `docs/workflow/0071_handoff_protocol.md`: `Handoff Target Types` and `Forbidden Handoff Behavior`
- `docs/README.md`: `Cross-AI Work Pattern`
- `docs/AGENTS.md`: `Mandatory working behavior`
- `docs/0001_index.md`: `Non-Negotiable Behavior`
- `docs/evidence/0003_laravel_to_go_transition_progress_ledger.md`: workflow/docs quality note
