# Handoff: AI Workstyle Template Update

## Date

2026-06-07

## Active Scope

Update Codex and Web AI templates so another AI CLI or browser AI can reproduce the current work style.

## Current Branch Or Source Snapshot

Local workspace: `/home/asyraf/Code/go/pos-go`

## Files Included

```text
docs/templates/0120_prompt_authoring_rules.md
docs/templates/0121_codex_session_prompts.md
docs/templates/0122_web_ai_session_prompts.md
docs/templates/0110_domain_scope_packet.md
docs/templates/README.md
scripts/audit_ai_rules.sh
```

## Files Changed

```text
docs/templates/0120_prompt_authoring_rules.md
docs/templates/0121_codex_session_prompts.md
docs/templates/0122_web_ai_session_prompts.md
docs/templates/0110_domain_scope_packet.md
docs/templates/README.md
scripts/audit_ai_rules.sh
docs/handoffs/2026-06-07-ai-workstyle-template-update.md
```

## Files Forbidden To Touch

No Go source, migrations, runtime config, or production secrets were in scope.

## Blueprint Referenced

```text
docs/blueprints/0012_laravel_to_go_api_transition_master_plan.md
docs/blueprints/0010_capability_control_foundation.md
```

## ADR And Rules Referenced

```text
docs/0002_decision_policy.md
docs/workflow/0071_handoff_protocol.md
docs/workflow/0072_transition_progress_ledger_protocol.md
```

## Decisions Made

- Terminal coding agents should execute the largest safe slice inside one active step.
- Terminal coding agents should use short progress updates and compact final reports.
- Web AI remains read-only by default and should draft patch plans, exact CLI commands, docs text, evidence text, and handoff text.
- Terminal Codex or the owner runs CLI commands and applies repository changes unless exact GitHub mutation permission is provided.
- Missing Laravel data should be requested as the smallest specific source batch.
- ADR-level decisions should be asked with 2-3 options, tradeoffs, and a recommendation when clear.

## Proof Collected

AI rules audit passed:

```text
bash scripts/audit_ai_rules.sh
```

Visible result:

```text
[PASS] AI rules audit passed
```

## Tests Or Commands Run

```text
bash scripts/audit_ai_rules.sh
```

Result:

```text
[PASS] AI rules audit passed
```

## Gaps Still Open

- Runtime capability middleware remains the next implementation step.

## Next Valid Active Step

Continue `docs/blueprints/0010_capability_control_foundation.md` with runtime capability check middleware/policy.

## Estimated Scope Progress Percentage

AI workstyle template update: 100%.

Overall Laravel-to-Go transition: unchanged at 18%.

## Estimated Context-Window Status

Enough context remains for one focused follow-up step if needed.
