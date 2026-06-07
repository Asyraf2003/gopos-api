# Handoff: AI Execution Channel Boundaries

## Date

2026-06-07

## Active Scope

Clarify Web AI, Terminal Codex, owner/local terminal, and explicit collaboration workflow boundaries in AI prompt and handoff docs.

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
docs/evidence/0003_laravel_to_go_transition_progress_ledger.md
docs/handoffs/2026-06-07-prompt-template-selection-rule.md
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
docs/handoffs/2026-06-07-ai-execution-channel-boundaries.md
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
docs/handoffs/2026-06-07-prompt-template-selection-rule.md
```

## Decisions Made

- Web AI is read-only analysis and planning by default.
- Web AI normally returns analysis, patch plans, command plans, docs text, evidence text, and validation notes.
- Web AI may provide commands for owner/local terminal execution.
- Web AI must not assume Terminal Codex is the executor.
- Terminal Codex is the local CLI implementation agent.
- Terminal Codex works through local CLI execution and owner feedback.
- Terminal Codex must not assume Web AI collaboration unless the owner explicitly provided a collaboration packet.
- Owner/local terminal may execute command plans, collect proof, and push changes.
- Collaboration packets are special-case only and require explicit owner instruction for a specific problem.
- Web AI templates now prefer `COMMAND PLAN FOR OWNER / LOCAL TERMINAL` and `PROOF THE OWNER OR TERMINAL AGENT MUST RUN`.
- `OPTIONAL HANDOFF TEXT FOR CODEX` is only for owner-requested Codex work or explicit Codex handoff tasks.
- Handoff target types are owner/local terminal, Terminal Codex, Web AI, and explicit collaboration packet.
- Cascade updates were required because the rule changes how future sessions start, continue, and hand off.

## Proof Collected

Proof is pending until the required `rg` and `make verify` commands pass after this handoff is written.

## Tests Or Commands Run

Pending:

```text
rg -n "AI Execution Channel Rule|owner/local terminal|COMMAND PLAN FOR OWNER|OPTIONAL HANDOFF TEXT FOR CODEX|collaboration packet|Web AI did not assume Codex|Terminal Codex works through local CLI" docs/templates docs/workflow docs/handoffs docs/evidence docs/README.md docs/AGENTS.md docs/0001_index.md
make verify
```

## Gaps Still Open

- Route-to-capability audit script remains the next capability-control foundation gap.
- POS CRUD remains blocked.

## Next Valid Active Step

Return to `docs/blueprints/0010_capability_control_foundation.md` and add the route-to-capability audit script after this docs-only workflow fix is accepted.

Recommended next execution channel:

- target channel: owner/local terminal or Terminal Codex, depending on the owner's next prompt;
- if owner/local terminal: use Web AI for read-only review after owner pushes;
- if Terminal Codex: use `docs/templates/0121_codex_session_prompts.md`;
- if Web AI: use `docs/templates/0122_web_ai_session_prompts.md`;
- if collaboration: owner must explicitly request a collaboration packet.

## Estimated Scope Progress Percentage

AI execution-channel hardening scope: 90% until proof passes.

Laravel-to-Go transition: unchanged at 20%.

## Estimated Context-Window Status

Enough context remains to run proof and update this handoff with final results.

## Assessment For Web AI Review

Scores use 0 as best and 5 as worst.

- Web AI/Codex confusion risk: 1
- Codex-default bias risk: 1
- owner/local terminal workflow ambiguity: 1
- collaboration packet ambiguity: 1
- template compliance risk: 1
- docs cascade risk: 1

Web AI should re-check these files and sections after owner pushes:

- `docs/templates/0120_prompt_authoring_rules.md`: `AI Execution Channel Rule`
- `docs/templates/0122_web_ai_session_prompts.md`: permission model, expected output, self-check sections
- `docs/templates/0121_codex_session_prompts.md`: top-level Codex execution rules and self-check sections
- `docs/workflow/0071_handoff_protocol.md`: `Handoff Target Types` and forbidden handoff behavior
- `docs/README.md`: cross-AI work pattern
- `docs/AGENTS.md`: mandatory working behavior
- `docs/0001_index.md`: non-negotiable behavior
- `docs/evidence/0003_laravel_to_go_transition_progress_ledger.md`: workflow/docs quality note
